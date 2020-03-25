.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/root root/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

local: build
	sls offline --useDocker --printOutput start

package: clean build
	serverless package

docs:
	docker run -it --rm -p 80:80 \
	-v $(pwd)/swagger.yaml:/usr/share/nginx/html/swagger.yaml \
	-e SPEC_URL=swagger.yaml redocly/redoc:latest
