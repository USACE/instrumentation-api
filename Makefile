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
	redoc-cli serve -p 4000 apidoc.yaml
