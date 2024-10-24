import util from "util";
import os from "os";

import puppeteer, { Page, Browser } from "puppeteer-core";
import {
  SecretsManagerClient,
  GetSecretValueCommand,
} from "@aws-sdk/client-secrets-manager";
import { S3Client } from "@aws-sdk/client-s3";
import { Upload } from "@aws-sdk/lib-storage";
import createClient from "openapi-fetch";

import { castleLogoSvg } from "./svgInlineContent";
import { logoBackground } from "./base64InlineContent";
import { getIndexHtml, getHeaderTmpl, getFooterTmpl } from "./htmlContent";

import type { UUID } from "crypto";
import type { SecretsManagerClientConfig } from "@aws-sdk/client-secrets-manager";
import type { S3ClientConfig } from "@aws-sdk/client-s3";
import type { AwsCredentialIdentity } from "@aws-sdk/types";
import type { paths, components } from "../generated";
import type { SQSEvent } from "aws-lambda";

type ReportDownloadJob = components["schemas"]["ReportDownloadJob"];

interface EventMessageBody {
  report_config_id: UUID;
  job_id: UUID;
  is_landscape: boolean;
}

// TODO: it would be better to get expiry directly from the
// s3 uploader response if possible, in case it ever changes
const FILE_EXPIRY_DURATION_HOURS = 24;
const MOCK_APP_KEY = "appkey";
const WAIT_FOR_MS = 60_000 * 5; // 5 minutes

const accessKeyId = process.env.AWS_ACCESS_KEY_ID;
const secretAccessKey = process.env.AWS_SECRET_ACCESS_KEY;

let credentials: AwsCredentialIdentity | undefined;
if (accessKeyId && secretAccessKey) {
  credentials = {
    accessKeyId,
    secretAccessKey,
  };
}

const s3ClientConfig: S3ClientConfig = {
  endpoint: process.env.AWS_S3_ENDPOINT,
  region: process.env.AWS_S3_REGION,
  forcePathStyle: true,
};

if (credentials) {
  s3ClientConfig.credentials = credentials;
}

const smClientConfig: SecretsManagerClientConfig = {
  endpoint: process.env.AWS_SM_ENDPOINT,
  region: process.env.AWS_SM_REGION,
};

const apiBaseUrl = process.env.API_BASE_URL;
const s3WriteToBucket = process.env.AWS_S3_WRITE_TO_BUCKET;
const s3WriteToBucketPrefix = process.env.AWS_S3_WRITE_TO_BUCKET_PREFIX;
const smApiKeySecretId = process.env.AWS_SM_API_KEY_SECRET_ID;
const smKey = process.env.AWS_SM_KEY ?? "";
const puppeteerExecutablePath = process.env.PUPPETEER_EXECUTABLE_PATH;
const smMockRequest =
  String(process.env.AWS_SM_MOCK_REQUEST).toLowerCase() === "true";
const chromeDumpIO = String(process.env.CHROME_DUMPIO).toLowerCase() === "true";

type ProcessEventArgs = {
  event: EventMessageBody;
  s3Client: S3Client;
  apiKey: string;
};

type retryFunc<T> = (args: T) => Promise<boolean>;

async function retry(
  func: retryFunc<ProcessEventArgs>,
  args: ProcessEventArgs,
  retries: number,
): Promise<boolean> {
  for (let i = 0; i < retries; ++i) {
    if (await func(args)) {
      return true;
    }
    console.log("retrying...");
    await sleep(1000);
  }
  return false;
}

function sleep(ms: number) {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

// these flags are acceptable because we are only using chrome as a renderer for Plotly
// no external site data is loaded (sans the internal MIDAS API) via cdn and all packages are bundled
// this is needed because lambda cannot load custom security profiles (seccomp) and uses seccomp BPF be default
// docker also provides a layer of isolation, as this container is run as non-root, least privileged user
const chromiumArgs = [
  "--headless",
  "--disable-dev-shm-usage",
  "--disable-software-rasterizer",
  "--disable-infobars",
  "--no-sandbox",
  "--disable-setuid-sandbox",
  "--disable-gpu=False",
  "--enable-webgl",
  "--single-process",
  "--user-data-dir=/tmp/user-data",
  "--data-path=/tmp/data-path",
  "--homedir=/tmp",
  "--disk-cache-dir=/tmp/cache-dir",
  "--database=/tmp/database",
];

let browserPromise: Promise<Browser>;

async function prepareBrowser() {
  return puppeteer.launch({
    executablePath: puppeteerExecutablePath,
    args: chromiumArgs,
    dumpio: chromeDumpIO,
    headless: true,
    timeout: WAIT_FOR_MS,
    devtools: false,
    protocolTimeout: 20_000,
    ignoreHTTPSErrors: true,
    defaultViewport: null,
    ignoreDefaultArgs: ["--enable-automation"],
  });
}

async function waitForDOMStable(
  page: Page,
  options = { timeout: 30_000, idleTime: 2000 },
): Promise<void> {
  await page.evaluate(
    ({ timeout, idleTime }) =>
      new Promise((resolve, reject) => {
        setTimeout(() => {
          observer.disconnect();
          const msg =
            `timeout of ${timeout} ms ` +
            "exceeded waiting for DOM to stabilize";
          reject(Error(msg));
        }, timeout);
        const observer = new MutationObserver(() => {
          clearTimeout(timeoutId);
          timeoutId = setTimeout(finish, idleTime);
        });
        const config = {
          attributes: true,
          childList: true,
          subtree: true,
        };
        observer.observe(document.body, config);
        const finish = () => {
          observer.disconnect();
          resolve("done");
        };
        let timeoutId = setTimeout(finish, idleTime);
      }),
    options,
  );
}

export async function handler(event: SQSEvent): Promise<void> {
  console.log("Executing user:\n", util.inspect(os.userInfo()));
  const s3Client = new S3Client(s3ClientConfig);

  let apiKey = MOCK_APP_KEY;
  if (!smMockRequest) {
    const client = new SecretsManagerClient(smClientConfig);
    const command = new GetSecretValueCommand({ SecretId: smApiKeySecretId });
    const res = await client.send(command);
    const resJson = res.SecretString ? JSON.parse(res.SecretString) : undefined;
    apiKey = resJson[smKey] ?? "";
  }

  for (const rec of event.Records) {
    await retry(
      processEvent,
      { event: JSON.parse(rec.body), s3Client, apiKey },
      5,
    );
  }
}

async function processEvent(args: ProcessEventArgs): Promise<boolean> {
  const { event, s3Client, apiKey } = args;

  console.log("recieved event from queue", event);

  const {
    report_config_id: rcId,
    job_id: jobId,
    is_landscape: isLandscape,
  } = event;

  let browser: Browser | null = null;
  let page: Page | null = null;

  try {
    if (!browserPromise) {
      browserPromise = prepareBrowser();
    }
    browser = await browserPromise;
    page = await browser.newPage();

    // bubble up events from headless browser
    page.on("console", (message) => console.log(`Console: ${message.text()}`));
    page.on("pageerror", ({ message }) => console.log(`Error: ${message}`));
    page.on("requestfailed", (request) =>
      console.log(
        `Request failed: ${request.failure()?.errorText} ${request.url()}`,
      ),
    );

    await page.setContent(
      getIndexHtml(event.is_landscape ? "landscape" : "portrait"),
    );
    // This is supposed to fix the problem of too many WebGL contexts
    // in the case where a page has many different plots but there are some
    // errors that the WebGLContext elements are not supported.
    // TODO: check that WebGL 1 is enabled on the apk installation of chrome
    //
    // await page.addScriptTag({ url: "https://unpkg.com/virtual-webgl@1.0.6/src/virtual-webgl.js" });
    await page.addScriptTag({ path: "./report.mjs" });

    const { districtName, projectName } = await page.evaluate(
      async (id, url, apikey, isLandscape) => {
        return await window.processReport(id, url, apikey, isLandscape);
      },
      rcId,
      apiBaseUrl,
      apiKey,
      isLandscape,
    );

    // wait for all content to load before exporting to PDF
    await page.waitForNetworkIdle({ timeout: WAIT_FOR_MS });
    await waitForDOMStable(page, { timeout: WAIT_FOR_MS, idleTime: 2000 });

    const buf = await page.pdf({
      format: "letter",
      scale: 1,
      displayHeaderFooter: true,
      headerTemplate: getHeaderTmpl(logoBackground),
      footerTemplate: getFooterTmpl(castleLogoSvg, districtName, projectName),
      preferCSSPageSize: true,
      timeout: WAIT_FOR_MS,
    });
    let statusCode: number | undefined;

    const fileKey = `${s3WriteToBucketPrefix}/${rcId}/${jobId}/${new Date().toISOString().split("T")[0]}_midas_report.pdf`;

    statusCode = await upload(s3Client, buf, fileKey);
    await updateJob(apiKey, jobId, rcId, statusCode, fileKey);

    await browser.close();
    console.log("completed job", jobId);
    return true;
  } catch (err) {
    console.error("error during job...", err);
    if (browser?.connected) {
      console.log("closing browser...");
      await browser.close();
    }
    return false;
  }
}

async function upload(
  s3Client: S3Client,
  buf: Buffer,
  key: string,
): Promise<number | undefined> {
  const uploader = new Upload({
    client: s3Client,
    params: { Bucket: s3WriteToBucket, Key: key, Body: buf },
  });
  const res = await uploader.done();

  return res.$metadata.httpStatusCode;
}

async function updateJob(
  apiKey: string,
  jobId: string,
  rcId: string,
  statusCode?: number,
  fileKey?: string,
): Promise<void> {
  const apiClient = createClient<paths>({
    baseUrl: apiBaseUrl,
    headers: {
      "Content-Type": "application/json",
    },
  });
  const reportJobPayload = {
    params: {
      query: {
        key: apiKey,
      },
      path: {
        job_id: jobId,
      },
    },
  };

  if (statusCode !== 200) {
    const failedJob: ReportDownloadJob = {
      status: "FAIL",
      progress: 0,
    };
    console.error(
      `error: pdf upload failed; status code: ${statusCode}; job_id: ${jobId}; report_config_id: ${rcId};`,
    );

    const { data: failData, error: failErr } = await apiClient.PUT(
      "/report_jobs/{job_id}",
      {
        ...reportJobPayload,
        body: failedJob,
      },
    );
    throw new Error(JSON.stringify(failData ?? failErr));
  }

  const j: ReportDownloadJob = {
    status: "SUCCESS",
    progress: 100,
    file_key: fileKey,
    file_expiry: new Date(
      new Date().getTime() + FILE_EXPIRY_DURATION_HOURS * 60 * 60 * 1000,
    ).toISOString(),
  };

  // NOTE: if this fails, the pdf should be automatically deleted anyway by lifetime policy
  const { data, error } = await apiClient.PUT("/report_jobs/{job_id}", {
    ...reportJobPayload,
    body: j,
  });
  if (error) {
    throw new Error(JSON.stringify(error));
  }
  console.log("SUCCESS", data);
}
