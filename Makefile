packagename = corpsmap-instrumentation-api.zip

.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/root root/main.go

clean:
	rm -rf ./bin ./vendor $(packagename) Gopkg.lock

package: clean build
	zip -r $(packagename) bin

deploy: package
	aws s3 cp $(packagename) s3://corpsmap-lambda-zips/$(packagename)

docs:
	redoc-cli serve -p 4000 apidoc.yaml
