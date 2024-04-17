declare global {
    namespace NodeJS {
        interface ProcessEnv {
            AWS_SECRET_ARN_API_KEY: string;
            HOST: string;
        }
    }
}

export { }
