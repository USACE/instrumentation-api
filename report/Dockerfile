ARG NODE_VERSION=20

FROM public.ecr.aws/lambda/nodejs:${NODE_VERSION} as builder
WORKDIR /usr/app
COPY package.json ./
RUN npm install
COPY index.ts  ./
RUN npm run build

FROM public.ecr.aws/lambda/nodejs:${NODE_VERSION}
WORKDIR ${LAMBDA_TASK_ROOT}
COPY --from=builder /usr/app/dist/* ./
CMD ["index.handler"]
