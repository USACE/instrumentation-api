declare global {
  namespace NodeJS {
    interface ProcessEnv {
      AWS_S3_ENDPOINT: string;
      API_BASE_URL: string;
      AWS_SM_BASE_URL: string;
      AWS_SM_API_KEY_ARN: string;
      AWS_S3_WRITE_TO_BUCKET: string;
      AWS_S3_SKIP_UPLOAD: boolean;
      AWS_SESSION_TOKEN: string;
      AWS_SM_MOCK_REQUEST: boolean;
      PUPPETEER_EXECUTABLE_PATH: string;
    }
  }
}

export {};
