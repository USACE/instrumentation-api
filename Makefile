packagename = instrumentation-api.zip

.PHONY: build clean

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/api api/main.go

clean:
	rm -rf ./bin ./vendor $(packagename) Gopkg.lock

package: clean build
	zip -r $(packagename) bin

deploy-dev: package
	aws s3 cp $(packagename) s3://corpsmap-lambda-zips/$(packagename)

deploy-test: package
	aws s3 cp $(packagename) s3://rsgis-lambda-zips/$(packagename)
