import fs from "node:fs";
import puppeteer from "puppeteer-core";
import { S3Client } from "@aws-sdk/client-s3";
import { Upload } from "@aws-sdk/lib-storage";

import { ApiClient } from "./generated";
import { FetchHttpRequest } from "./generated/core/FetchHttpRequest";
import { processReport } from "./report-client/report.mjs";

import type { UUID } from "crypto";
import type { GetSecretValueResponse } from "@aws-sdk/client-secrets-manager";
import type { ReportDownloadJob } from "./generated";

interface EventMessageBody {
  report_config_id: UUID;
  job_id: UUID;
}

// TODO: it would be better to get expiry directly from the
// s3 uploader response if possible, in case it ever changes
const FILE_EXPIRY_DURATION_HOURS = 24;
const MOCK_APP_KEY = "appkey";

const s3ClientConfig = { endpoint: process.env.AWS_S3_ENDPOINT, region: process.env.AWS_S3_REGION };

const apiBaseUrl = process.env.API_BASE_URL;
const s3WriteToBucket = process.env.AWS_S3_WRITE_TO_BUCKET;
const sessionToken = process.env.AWS_SESSION_TOKEN;
const smBaseUrl = process.env.AWS_SM_BASE_URL;
const smApiKeyArn = process.env.AWS_SM_API_KEY_ARN;
const puppeteerExecutablePath = process.env.PUPPETEER_EXECUTABLE_PATH;
const smMockRequest = process.env.AWS_SM_MOCK_REQUEST;
const s3SkipUpload = process.env.AWS_S3_SKIP_UPLOAD;

export async function handler(event: EventMessageBody): Promise<void> {
  const s3Client = new S3Client(s3ClientConfig);

  let apiKey = MOCK_APP_KEY;
  if (!smMockRequest) {
    const req = new FetchHttpRequest({
      BASE: smBaseUrl!,
      VERSION: "",
      WITH_CREDENTIALS: false,
      CREDENTIALS: "omit",
    });
    const res = await req.request<GetSecretValueResponse>({
      method: "GET",
      url: `/secretsmanager/get?secretId=${smApiKeyArn}`,
      headers: { [`X-Aws-Parameters-Secrets-Token`]: sessionToken },
    });
    apiKey = res.SecretString!;
  }

  const apiClient = new ApiClient({ BASE: apiBaseUrl });

  const { report_config_id: rcId, job_id: jobId } = event;

  const browser = await puppeteer.launch({
    executablePath: puppeteerExecutablePath,
    args: [
      "--headless",
      "--disable-dev-shm-usage",
    ]
  });
  const page = await browser.newPage();
  const htmlContent = fs.readFileSync("/usr/src/app/index.html");

  await page.setContent(htmlContent.toString());
  await page.addScriptTag({ content: `${processReport}` })

  await page.evaluateOnNewDocument(async (id, client, key) => {
    const { newPlot, addTraces } = await import("plotly.js-dist-min");
    await page.addScriptTag({ content: `${newPlot} ${addTraces}` })
    processReport(id, client, key, newPlot, addTraces);
  }, rcId, apiClient, apiKey);

  const buf = await page.pdf({ format: "A4" });
  let statusCode: number | undefined = 201;
  if (!s3SkipUpload) {
    statusCode = await upload(s3Client, buf, rcId, jobId);
  }
  await updateJob(apiClient, apiKey, jobId, rcId, statusCode);

  await browser.close();
}

async function upload(
  s3Client: S3Client,
  buf: Buffer,
  rcId: UUID,
  jobId: UUID,
): Promise<number | undefined> {
  const key = `/${rcId}/${jobId}__iso8601_${new Date().toISOString().split("T")[0]}__report.pdf`;

  const uploader = new Upload({
    client: s3Client,
    params: { Bucket: s3WriteToBucket, Key: key, Body: buf },
  });
  const res = await uploader.done();

  return res.$metadata.httpStatusCode;
}

async function updateJob(
  apiClient: ApiClient,
  apiKey: string,
  jobId: string,
  rcId: string,
  statusCode: number | undefined,
): Promise<void> {
  if (statusCode !== 201) {
    const failedJob: ReportDownloadJob = {
      status: "FAIL",
      progress: 0,
    };
    console.error(
      `error: pdf upload failed; status code: ${statusCode}; job_id: ${jobId}; report_config_id: ${rcId};`,
    );
    apiClient.reportConfig
      .putProjectsReportConfigsJobs(jobId, failedJob, apiKey)
      .then(console.log, console.error);
  }

  const j: ReportDownloadJob = {
    status: "SUCCESS",
    progress: 100,
    file_expiry: new Date(
      new Date().getHours() + FILE_EXPIRY_DURATION_HOURS,
    ).toISOString(),
  };

  // NOTE: if this fails, the pdf should be automatically deleted anyway by lifetime policy
  await apiClient.reportConfig
    .putProjectsReportConfigsJobs(jobId, j, apiKey)
    .then(console.log, console.error);
}
