# inspired by https://github.com/Zenika/alpine-chrome/blob/ab0d418bcbc81e1254bdb7875ec06753e3fd68c5/Dockerfile

ARG NODE_VERSION=22.1
ARG ALPINE_VERSION=3.19

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} as builder
WORKDIR /usr/app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION}
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

# Needed for Lambda CMake Backtrace
RUN apk add --no-cache --update --repository=https://dl-cdn.alpinelinux.org/alpine/v3.16/main/ libexecinfo-dev

# packages needed for Lambda ric / rie and Nodejs
RUN apk add --no-cache \
	make gcc g++ python3 git npm yarn \
	wget cmake unzip curl-dev autoconf automake libtool

RUN ln -s /usr/share/aclocal /usr/local/share/aclocal

COPY chrome/local.conf /etc/fonts/local.conf

# Add Chrome as a user
RUN mkdir -p /usr/src/app \
	&& adduser -D chrome \
	&& chown -R chrome:chrome /usr/src/app
# Run Chrome as non-privileged
USER chrome
WORKDIR /usr/src/app

ENV CHROME_BIN=/usr/bin/chromium-browser \
	CHROME_PATH=/usr/lib/chromium/

# Autorun chrome headless
ENV CHROMIUM_FLAGS="--disable-software-rasterizer --disable-dev-shm-usage"

RUN npm install aws-lambda-ric

RUN mkdir -p ~/.aws-lambda-rie \
	&& wget -O ~/.aws-lambda-rie/aws-lambda-rie \
	https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
	&& chmod +x ~/.aws-lambda-rie/aws-lambda-rie

COPY --from=builder /usr/app/dist/ ./
COPY --from=builder /usr/app/index.html ./
COPY --from=builder /usr/app/entry_script.sh ./

ENV PUPPETEER_EXECUTABLE_PATH /usr/bin/chromium-browser

ENTRYPOINT ["./entry_script.sh"]
CMD ["main.handler"]
