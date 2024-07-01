declare global {
  namespace NodeJS {
    interface ProcessEnv {
      API_BASE_URL: string;
      AWS_S3_ENDPOINT: string;
      AWS_S3_WRITE_TO_BUCKET: string;
      AWS_S3_WRITE_TO_BUCKET_PREFIX: string;
      AWS_SM_API_KEY_SECRET_ID: string;
      AWS_SM_ENDPOINT: string;
      AWS_SM_KEY: string;
      AWS_SM_MOCK_REQUEST: string;
      AWS_SM_REGION: string;
      PUPPETEER_EXECUTABLE_PATH: string;
    }
  }
  interface Window {
    processReport: (
      id: UUID,
      url: string,
      apiKey: string,
      isLandscape: boolean,
    ) => Promise<{ districtName: string, projectName: string }>;
  }
}

export {};
