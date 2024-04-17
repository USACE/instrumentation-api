# MIDAS API

A REST API to manage instrumentation data for the MIDAS (Monitoring Instrumentation Data Acquisition System) web application. Built with Go, using the [Echo](https://github.com/labstack/echo) web framework.

MIDAS project management is tracked in https://github.com/USACE/instrumentation

## Documentation

"[swag](https://github.com/swaggo/swag)" code annotations are used in [./api/internal/handler/](./api/internal/handler/) to generate OpenApi 2.0 docs that are available via `$API_HOSTNAME/swagger/index.html` on the API server. Generating documnetation can be done using the compose script: `./compose.sh mkdocs`

## How to Develop

### Quickstart - Running the API Stack Locally with Docker Compose

1. Install [Docker Engine](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/). In place of these two, one can install [Docker Desktop](https://docs.docker.com/desktop/).

2. [Install Go (Golang)](https://go.dev/doc/install) - needed for pulling peer dependencies when building the containers.

3. Copy the `.env.example` file to `.env` (e.g., `cp .env.example .env`). This provides configuration options to Docker Compose.

4. Use the [./compose.sh](./compose.sh) shell script to build and (re)start the Docker Compose services. You can now access the application locally through Docker. Use the `up` argument to start, `down` to spin down, and `clean` to spin down and remove volumes. Optionally add the `mock` argument to start/remove the mock datalogger service.

```sh
./compose.sh up    # or ./compose.sh up mock
./compose.sh down  # or ./compose.sh down mock
```

#### General Tab

| Field | Value                               |
| ----- | ----------------------------------- |
| Name  | localhost (or other preferred name) |

#### Connection Tab

| Field             | Value          |
| ----------------- | -------------- |
| Host name/address | db             |
| Port              | 5432 (default) |
| Username          | postgres       |
| Password          | postgres       |

Initialize the database and seed it with some data (./compose.sh up runs this for you)

### Running the Go API for Local Development

Either of these options starts the API at `localhost:$API_PORT`, where `$API_PORT` is the variable set in your project's `.env` file. The API uses JSON Web tokens (JWT) for authorization by default. To disable JWT for testing or development, you can set the environment variable `JWT_DISABLED=TRUE`.

### With Visual Studio Code Debugger

You can use the launch.json file in this repository in lieu of `go run main.go` to run the API in the VSCode debugger. This takes care of the required environment variables to connect to the database.

### Without Visual Studio Code Debugger

Set the following environment variables and type `go run main.go` from the top level of this repository.

```sh
INSTRUMENTATION_DB_USER=postgres
INSTRUMENTATION_DB_PASS=postgres
INSTRUMENTATION_DB_NAME=postgres
INSTRUMENTATION_DB_HOST=localhost
INSTRUMENTATION_DB_SSLMODE=disable
JWT_DISABLED=FALSE
```

## Running Tests

Regression tests are maintained for the project using Go [httptest](https://pkg.go.dev/net/http/httptest) from the standard library. These are integration tests that run inside of the `api` docker container. Tests are located in the handlers folder, using the virtual package (not compiled to production binary) `handler_test`. Each test is an instance of `HTTPTest` struct with relavent properties. (test name, url path, etc.). A notable property is the `ExpectedSchema`, which accepts a `gojsonschema.Schema`. This is used to validate the response body matches the json schema string, otherwise, the test will fail. All tests are run against a test database seeded with SQL scripts locally; see the [./migrate/local/](./migrate/local/) folder.

To run specific tests, you can pass optional commands to the test script in `./compose.sh` like so, where `*args` are passed to `go test *args`. The `-rm` flag will automatically remove containers after they run:
```bash
./compose.sh test [-rm] *args
```
## How To Deploy

### Deploying Develop Core API & Telemetry API

Development deployments are done though [CI (Continuous Integration) scripts](./.github) using [Github Actions](https://docs.github.com/en/actions). The [core api](./api) and [telemetry api](./telemetry) are tested, built, and pushed to AWS ECR, where they should re-deploy on container push when the CI pipelines successfully finish. If the container does not automattically re-deploy on ECR push, it can be manually deployed from the AWS console.

Test and prodution deployments are currently done manually. The [./build_ib.sh](./build_ib.sh) can be used to build the application with hardened images sourced from Ironbank. Afterwards, the build images should be pushed to test and/or prod via cli.

### Postgres Database on AWS Relational Database Service (RDS)

First, make sure any extensions (such as PostGIS) are installed on the RDS instance.

Flyway Migrations are used for automated database migrations in order to keep the schemas of the local environment in sync with develop and stable. Any differentiations in these databases (such as loading test data in the local environment) must be specified in their respective folders within the [./sql](./sql) directory ("common" applying to all environments). For database versioning, each migration script must be incrementally applied (e.g. `V1.2.3__migration.sql` -> `V1.2.4__migration.sql`). Versioned migrations must not change after they are run against the database. Instead, you must create another "version" to make modifications to a schema. Repeatable migrations (e.g. `R__01_repeatable_migration.sql`), may be modified if the overall schema does not change.

## dcs-loader

SQS-Worker to parse CSV Files of timeseries measurements on AWS S3 and post contents to the core api. 

Works with ElasticMQ for local testing with a SQS-compatible interface. Variables noted "Used for local testing" typically do not need to be provided when deployed, for example to AWS. They can be omitted completely or set to "" if not required.

### Environment Variables

| Variable                       | Example Value                                                            | Notes                  |
| ------------------------------ | ------------------------------------------------------------------------ | ---------------------- |
| AWS_ACCESS_KEY_ID              | AKIAIOSFODNN7EXAMPLE                                                     | Used for local testing |
| AWS_SECRET_ACCESS_KEY          | wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY                                 | Used for local testing |
| AWS_DEFAULT_REGION             | us-east-1                                                                | Used for local testing |
| LOADER_POST_URL                | http://instrumentation-api_api_1/instrumentation/timeseries_measurements |                        |
| LOADER_API_KEY                 | appkey                                                                   |                        |
| LOADER_AWS_S3_ENDPOINT         | http://minio:9000                                                        | Used for local testing |
| LOADER_AWS_S3_REGION           | us-east-1                                                                |                        |
| LOADER_AWS_S3_DISABLE_SSL      | False                                                                    |                        |
| LOADER_AWS_S3_FORCE_PATH_STYLE | True                                                                     |                        |
| LOADER_AWS_SQS_QUEUE_NAME      | instrumentation-dcs-goes                                                 |                        |
| LOADER_AWS_SQS_ENDPOINT        | http://elasticmq:9324                                                    | Used for local testing |
| LOADER_AWS_SQS_QUEUE_URL       | http://elasticmq:9324/queue/instrumentation-dcs-goes                     | Used for local testing |
| LOADER_AWS_SQS_REGION          | elasticmq                                                                | Used for local testing |

### Example Input File

```
869465fc-dc1e-445e-81f4-9979b5fadda9,2021-03-01T15:30:00Z,27.6800
869465fc-dc1e-445e-81f4-9979b5fadda9,2021-03-01T15:00:00Z,27.6200
869465fc-dc1e-445e-81f4-9979b5fadda9,2021-03-01T14:30:00Z,27.5500
869465fc-dc1e-445e-81f4-9979b5fadda9,2021-03-01T14:00:00Z,27.4400
```
