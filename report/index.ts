import fs from "fs";
import fetch from "node-fetch";
import { SQSEvent, SQSRecord } from "aws-lambda";
import { S3Client } from "@aws-sdk/client-s3"; // ES Modules import
import { Upload } from "@aws-sdk/lib-storage";
import { UUID } from "crypto";
import {
  ApiClient,
  Measurement,
  PlotConfig,
  PlotConfigTimeseriesTrace,
  Timeseries,
} from "./generated";
import PDFDocument from "pdfkit";
import Plotly, { Dash, PlotType } from "plotly.js-dist-min";
import { GetSecretValueResponse } from "@aws-sdk/client-secrets-manager";

interface EventMessageBody {
  report_config_id: UUID;
}

interface TimeWindow {
  after: string | undefined;
  before: string | undefined;
}

const MOCK_APP_KEY = "appkey";

const XY_POS_LOOKUP: number[][] = [
  [0, 15],
  [0, 15],
];
const MAX_POS_LOOKUP_IDX = 1;

const s3ClientConfig = { endpoint: process.env.AWS_S3_ENDPOINT };

const apiBaseUrl = process.env.API_BASE_URL;
const smBaseUrl = process.env.AWS_SM_BASE_URL;
const smApiKeyArn = process.env.AWS_SM_API_KEY_ARN;
const s3WriteToBucket = process.env.AWS_S3_WRITE_TO_BUCKET;
const sessionToken = process.env.AWS_SESSION_TOKEN!;

const __smMockRequest = String(process.env.AWS_SM_MOCK_REQUEST).toLowerCase();
const smMockRequest =
  __smMockRequest === "true" || __smMockRequest === "1" ? true : false;

export async function handler(event: SQSEvent) {
  const s3Client = new S3Client(s3ClientConfig);

  let apiKey = MOCK_APP_KEY;
  if (!smMockRequest) {
    const res = await fetch(
      `${smBaseUrl}/secretsmanager/get?secretId=${smApiKeyArn}`,
      {
        headers: { [`X-Aws-Parameters-Secrets-Token`]: sessionToken },
        method: "GET",
      },
    );
    const body: GetSecretValueResponse = await res.json();
    apiKey = body.SecretString!;
  }

  const apiClient = new ApiClient({ BASE: apiBaseUrl });

  for (let r of event.Records) {
    const key = `/${r.messageId}/report.pdf`;
    const path = `/tmp${key}`;

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

async function processRecord(
  r: SQSRecord,
  apiClient: ApiClient,
  apiKey: string,
  pdfPath: string,
) {
  const { report_config_id: rcID }: EventMessageBody = JSON.parse(r.body);

  const doc = new PDFDocument({
    info: {
      Title: "Midas Report",
      Author: "MIDAS",
      Subject: "Batch Plot Measurement Reports",
      Keywords: "MIDAS",
    },
  });
  doc.pipe(fs.createWriteStream(pdfPath));

  const rp = await apiClient.reportConfig.getReportConfigsPlotConfigs(
    rcID,
    apiKey,
  );

  // start at index 1 to leave room for title and description
  let plotPagePosIdx = 1;

  for (const pc of rp.plot_configs ?? []) {
    if (plotPagePosIdx > MAX_POS_LOOKUP_IDX) {
      doc.addPage();
      plotPagePosIdx = 0;
    }

    const { after, before } = parseDateRange(pc.date_range);

    const layout = { width: 800, height: 600 };

    let gd = await Plotly.newPlot("gd", [], layout);

    const traces = pc.display?.traces ?? [];

    // traces are pre-sorted
    const tracePromises = traces.map((tr) => {
      return async function () {
        const mm = await apiClient.timeseries.getTimeseriesMeasurements(
          tr.timeseries_id!,
          after,
          before,
          3000,
        );
        const trace = createTraceData(tr, mm.items!, pc);

        await Plotly.addTraces(gd, trace, tr.trace_order);
      };
    });

    Promise.all(tracePromises);

    const dataUrl = await Plotly.toImage(gd, { format: "png", ...layout });

    const xyPos = XY_POS_LOOKUP[plotPagePosIdx];
    if (xyPos === undefined) {
      throw new Error("invalid template xy position index");
    }

    doc.image(dataUrl, ...xyPos);
  }
}

function parseDateRange(dateStr: string | undefined): TimeWindow {
  if (dateStr === undefined) {
    return { after: undefined, before: undefined };
  }

  let a;
  let delta;
  let b = Date.now();
  let d = new Date(b);

  switch (String(dateStr).toLowerCase()) {
    case "lifetime":
      // arbirarity min date
      a = Date.parse("1800-01-01");
      break;
    case "5 years":
      delta = b - d.setUTCFullYear(d.getUTCFullYear() - 5);
      a = b - delta;
      break;
    case "1 year":
      delta = b - d.setUTCFullYear(d.getUTCFullYear() - 1);
      a = b - delta;
      break;
    default:
      const dateParts = String(dateStr).split(" ", 1);
      if (dateParts.length !== 2) {
        throw new Error("could not parse custom date string");
      }
      a = Date.parse(dateParts[0]);
      b = Date.parse(dateParts[1]);
  }

  let after;
  let before;

  if (a !== undefined) {
    after = new Date(a).toISOString();
  }
  if (b !== undefined) {
    before = new Date(b).toISOString();
  }

  return { after, before };
}

function createTraceData(
  tr: PlotConfigTimeseriesTrace,
  mm: Measurement[],
  pc: PlotConfig,
): Plotly.Data {
  const filteredItems = mm.filter((m) => {
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

  return {
    x: x,
    y: y,
    mode: `lines${tr.show_markers ? "+markers" : ""}`,
    line: {
      dash: tr.line_style as Dash,
      color: tr.color,
      width: tr.width,
    },
    marker: {
      size: Number(tr.width) ? Number(tr.width) * 2 + 3 : 5,
      color: tr.color,
    },
    name: tr.name,
    type: tr.trace_type as PlotType,
    yaxis: tr.y_axis,
    showlegend: true,
  };
}
