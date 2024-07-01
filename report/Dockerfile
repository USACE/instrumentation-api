# inspired by https://github.com/Zenika/alpine-chrome/blob/ab0d418bcbc81e1254bdb7875ec06753e3fd68c5/Dockerfile

ARG BUILDER_IMAGE=node:22.1-alpine3.19
ARG BASE_IMAGE=alpine:3.19

FROM ${BUILDER_IMAGE} as builder
WORKDIR /usr/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM ${BASE_IMAGE}

RUN addgroup -S pptruser && adduser -S -G pptruser pptruser \
	&& mkdir -p /home/pptruser/Downloads /usr/src/app \
	&& chown -R pptruser:pptruser /home/pptruser \
	&& chown -R pptruser:pptruser /usr/src/app

COPY --from=builder --chown=pptruser /usr/lib /usr/lib
COPY --from=builder --chown=pptruser /usr/local/lib /usr/local/lib
COPY --from=builder --chown=pptruser /usr/local/include /usr/local/include
COPY --from=builder --chown=pptruser /usr/local/bin /usr/local/bin

# Needed for Lambda CMake Backtrace
RUN apk add --no-cache --update --repository=https://dl-cdn.alpinelinux.org/alpine/v3.16/main/ libexecinfo-dev
# packages needed for Lambda ric / rie and Nodejs
RUN apk add --no-cache --repository=https://dl-cdn.alpinelinux.org/alpine/v3.19/main/ \
	make gcc g++ python3 git npm yarn \
	wget cmake unzip curl-dev autoconf automake libtool

WORKDIR /usr/src/app
RUN npm install aws-lambda-ric

# Installs latest Chromium dependencies from whatever branches they are available
RUN apk upgrade --no-cache --available \
	&& apk add --no-cache \
	ttf-freefont \
	font-noto-emoji \
	# Chromium version is dependent on branch
	# https://pkgs.alpinelinux.org/packages?name=chromium-swiftshader
	# make sure Chromium version is supported by Puppeteer
	# https://pptr.dev/supported-browsers
	&& apk add --no-cache \
	--repository=https://dl-cdn.alpinelinux.org/alpine/v3.19/community \
	chromium-swiftshader \
	&& apk add --no-cache \
	--repository=https://dl-cdn.alpinelinux.org/alpine/edge/community \
	font-wqy-zenhei

RUN ln -s /usr/share/aclocal /usr/local/share/aclocal

COPY chrome/local.conf /etc/fonts/local.conf

USER pptruser

ENV CHROME_BIN=/usr/bin/chromium-browser \
	CHROME_PATH=/usr/lib/chromium/

# these flags are acceptable because we are only using chrome as a renderer for Plotly
# no external site data is loaded (sans the internal MIDAS API) via cdn and all packages are bundled
# this is needed because lambda cannot load custom security profiles (seccomp) and uses seccomp BPF be default
ENV CHROMIUM_FLAGS="--disable-software-rasterizer --disable-dev-shm-usage --single-process --no-sandbox"

RUN mkdir -p ~/.aws-lambda-rie \
	&& wget -O ~/.aws-lambda-rie/aws-lambda-rie \
	https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
	&& chmod +x ~/.aws-lambda-rie/aws-lambda-rie

COPY --from=builder --chown=pptruser /usr/app/dist/ ./
COPY --from=builder --chown=pptruser /usr/app/entry_script.sh ./

ENV PUPPETEER_EXECUTABLE_PATH /usr/bin/chromium-browser

ENTRYPOINT ["./entry_script.sh"]
CMD ["main.handler"]
