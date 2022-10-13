packagename = instrumentation-api.zip

.PHONY: build clean docs

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/api main.go

clean:
	rm -rf ./bin ./vendor $(packagename) Gopkg.lock

package: clean build
	zip -r $(packagename) bin

deploy-dev: package
	aws s3 cp $(packagename) s3://corpsmap-lambda-zips/$(packagename)

deploy-test: package
	aws s3 cp $(packagename) s3://rsgis-lambda-zips/$(packagename)

docs:
	npx @redocly/cli serve -p 4000 apidoc.yaml
