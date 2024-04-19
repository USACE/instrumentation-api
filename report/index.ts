import fs from 'fs';
import fetch from 'node-fetch';
import { SQSEvent, SQSRecord } from 'aws-lambda';
import { S3Client } from "@aws-sdk/client-s3"; // ES Modules import
import { Upload } from '@aws-sdk/lib-storage';
import { UUID } from 'crypto';
import { ApiClient, Measurement, PlotConfig, Timeseries } from './generated';
import PDFDocument from 'pdfkit';
import Plotly from 'plotly.js-dist-min'
import { GetSecretValueResponse } from '@aws-sdk/client-secrets-manager';

interface EventMessageBody {
    report_config_id: UUID
}

interface TimeWindow {
    after: string | undefined;
    before: string | undefined;
}

type XyPosTemplate = {
    [key: string]: number[]
}

const XY_POS_TEMPLATE: XyPosTemplate = {
    [`0`]: [0, 15],
    [`1`]: [0, 15],
    [`2`]: [0, 15],
}

const MAX_PLOT_PAGE_POS_IDX = 3;

const s3ClientConfig = { endpoint: process.env.AWS_S3_ENDPOINT };

const apiBaseUrl = process.env.API_BASE_URL;
const smBaseUrl = process.env.AWS_SM_BASE_URL;
const smApiKeyArn = process.env.AWS_SM_API_KEY_ARN;
const s3WriteToBucket = process.env.AWS_S3_WRITE_TO_BUCKET;
const sessionToken = process.env.AWS_SESSION_TOKEN!;

const __smMockRequest = String(process.env.AWS_SM_MOCK_REQUEST).toLowerCase();
const smMockRequest = __smMockRequest === "true" || __smMockRequest === "1" ? true : false;

export async function handler(event: SQSEvent) {
    const s3Client = new S3Client(s3ClientConfig);

    let apiKey = "";
    if (!smMockRequest) {
        const res = await fetch(`${smBaseUrl}/secretsmanager/get?secretId=${smApiKeyArn}`, {
            headers: { [`X-Aws-Parameters-Secrets-Token`]: sessionToken },
            method: "GET"
        });
        const body: GetSecretValueResponse = await res.json();
        apiKey = body.SecretString!;
    }

    const apiClient = new ApiClient({ BASE: apiBaseUrl });

    for (let r of event.Records) {
        const key = `/${r.messageId}/report.pdf`
        const path = `/tmp${key}`

        processRecord(r, apiClient, apiKey, path);

        const buf = fs.createReadStream(path);
        const uploader = new Upload({
            client: s3Client,
            params: { Bucket: s3WriteToBucket, Key: key, Body: buf },
        });
        const res = await uploader.done();

        if (res.$metadata.httpStatusCode !== 200) {
            throw new Error(`pdf upload failed; response: ${res}`);
        }

        // TODO: POST sucessful upload entry to API
        // NOTE: if this fails, the pdf should be automatically deleted anyway by lifetime policy
    }
}

async function processRecord(r: SQSRecord, apiClient: ApiClient, apiKey: string, pdfPath: string) {
    const { report_config_id: rcID }: EventMessageBody = JSON.parse(r.body);

    const doc = new PDFDocument({
        info: {
            Title: "Midas Report",
            Author: "MIDAS",
            Subject: "Batch Plot Measurement Reports",
            Keywords: "MIDAS",
        }
    });
    doc.pipe(fs.createWriteStream(pdfPath));

    const rp = await apiClient.reportConfig.getReportConfigsPlotConfigs(rcID, apiKey);

    // start at index 1 to leave room for title and description
    let plotPagePosIdx = 1;

    for (const pc of rp.plot_configs ?? []) {
        if (plotPagePosIdx > MAX_PLOT_PAGE_POS_IDX) {
            doc.addPage();
            plotPagePosIdx = 0;
        }

        const tss = await apiClient.timeseries.getProjectsPlotConfigurationsTimeseries(pc.id!);
        const { after, before } = parseDateRange(pc.date_range);

        const layout = { width: 800, height: 600 };

        let gd = await Plotly.newPlot("gd", [], layout);

        for (const ts of tss) {
            const mm = await apiClient.timeseries.getTimeseriesMeasurements(ts.id!, after, before, 3000);
            const trace = createTraceData(ts, mm.items!, pc);

            await Plotly.addTraces(gd, trace);
        }

        const dataUrl = await Plotly.toImage(gd, { format: 'png', ...layout });

        const xyPos = XY_POS_TEMPLATE[String(plotPagePosIdx)];
        if (xyPos === undefined) {
            throw new Error("invalid template xy position index");
        }

        doc.image(dataUrl, ...xyPos);
    }
}

function parseDateRange(dateStr: string | undefined): TimeWindow {
    if (dateStr === undefined) {
        return { after: undefined, before: undefined }
    }

    let _after;
    let _before = Date.now()
    let d = new Date();

    switch (String(dateStr).toLowerCase()) {
        case "lifetime":
            break;
        case "5 years":
            _after = _before - d.setUTCFullYear(d.getUTCFullYear() - 5);
        case "1 year":
            _after = _before - d.setUTCFullYear(d.getUTCFullYear() - 1);
        default:
            const dateParts = String(dateStr).split("-", 1);

            if (dateParts.length !== 2 || dateParts[0].length !== 10 || dateParts[1].length !== 10) {
                throw new Error("invalid date cannot be parsed");
            }

            const [afterMonth, afterDay, afterYear] = dateParts[0].split("/");
            const [beforeMonth, beforeDay, beforeYear] = dateParts[1].split("/");
            _after = Date.parse(`${afterYear}-${afterMonth}-${afterDay}`);
            _before = Date.parse(`${beforeYear}-${beforeMonth}-${beforeDay}`);
    }

    let after;
    let before;

    if (_after !== undefined) {
        after = new Date(_after).toUTCString();
    }
    if (_before !== undefined) {
        before = new Date(_before).toUTCString();
    }

    return { after, before }
}


function createTraceData(ts: Timeseries, mm: Measurement[], pc: PlotConfig): Plotly.Data {
    const filteredItems = mm.filter(m => {
        if (pc.show_masked && pc.show_nonvalidated) return true;
        if (pc.show_masked && !m.validated) return false;
        else if (pc.show_masked && m.validated) return true;

        if (pc.show_nonvalidated && m.masked) return false;
        else if (pc.show_nonvalidated && !m.masked) return true;

        if (m.masked || !m.validated) return false;
        return true;
    });

    const x: Plotly.Datum[] = new Array(filteredItems.length);
    const y: Plotly.Datum[] = new Array(filteredItems.length);

    for (let i = 0; i < filteredItems.length; i++) {
        x[i] = mm[i].time as Plotly.Datum;
        y[i] = mm[i].value as Plotly.Datum;
    }

    return ts.parameter === 'precipitation' ? {
        x: x,
        y: y,
        type: 'bar',
        yaxis: 'y2',
        name: `${ts.instrument} - ${ts.name} (${ts.unit})` || '',
        showlegend: true,
    } : {
        x: x,
        y: y,
        type: 'scattergl',
        mode: 'lines',
        name: `${ts.instrument} - ${ts.name} (${ts.unit})` || '',
        showlegend: true,
    };
}