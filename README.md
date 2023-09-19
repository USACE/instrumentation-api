# Instrumentation API

A REST API to manage instrumentation data for the MIDAS (Monitoring Instrumentation Data Acquisition System) web application. Built with Go, using the [Echo](https://github.com/labstack/echo) web framework.

## Documentation

A [Postman](https://www.postman.com/api-documentation-tool/) documentation and testing environment is maintained at [`tests/postman/postman_environment.local`](./tests/postman/postman_environment.local.json). An [OpenAPI Doc](./docs/swagger/apidoc.json) is [automatically generated](https://github.com/USACE/instrumentation-api/blob/423e257f2a4fead223ec53e39008324e81345eb3/docker-compose.yml#L148) when running Swagger locally with docker-compose. See the "Running the Swagger UI to access API documentation locally" section below for more information.

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

### Running a Database for Local Development

After starting up Docker Compose, you will find these two services (among others) on `localhost`

1. A Postgres database with postgis schema installed using the Docker image [mdillon/postgis](https://hub.docker.com/r/mdillon/postgis/)

2. [pgAdmin4](https://www.pgadmin.org/) using the Docker image [dpage/pgadmin4](https://hub.docker.com/r/dpage/pgadmin4/)

To modify the database using pgAdmin4, open a web browser and go to `http://localhost:8081`, or whichever port number is the value set to variable `PGADMIN_PORT` in `.env`.

Login with `Email:postgres@postgres.com` and `Password:postgres` respectively.

Create a database connection to the Postgres database by right-clicking `servers --> register --> server` in the left menu tree. Enter the following information and click `save`.

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
INSTRUMENTATION_LAMBDACONTEXT=FALSE
JWT_DISABLED=FALSE
```

Note: When running the API locally, make sure environment variable `INSTRUMENTATION_LAMBDACONTEXT` is either **not set** or is set to `INSTRUMENTATION_LAMBDACONTEXT=FALSE`. `_LAMBDA_SERVER_PORT` and `AWS_LAMBDA_RUNTIME_API` should also be set if running under an AWS Lambda context.

## Running Tests

Regression tests are maintained for the project in the [aforementioned](#documentation) [Postman](https://www.postman.com/api-documentation-tool/) environments. They are run automatically by GitHub Actions through the [./compose.sh](./compose.sh) script using `./compose.sh test`.

In both cases, the Postman environment regression tests are run, then output. If the environment variable `REPORT` is set to `true`, then this output is sent to an HTML file. Otherwise, it is printed to the caller's stdout.

## Running the Swagger UI to access API documentation locally

An API Document conforming to the OpenAPI 3.0.0 specification is generated from the most recent Postman Collection saved to [`tests/postman/instrumentation-regression.postman_collection.json`](./tests/postman/instrumentation-regression.postman_collection.json). When the collection file is modified and overwritten, an updated apidoc.json will be automatically created by a `swagger_init` docker service at [docs/swagger/apidoc.json](./docs/swagger/apidoc.json). To start the Swagger UI server and sync the apidoc.json with the Postman Collection, run `docker compose -f docker-compose.swagger.yml up -d`. This command is also executed in [./startup.sh](./startup.sh).

Note:

- This service will need to be restarted if any changes are made to the Postman Collection file (i.e. when it is manually exported and overwritten).

- Unlike the postman collection, the `.env.json` file supplied to the migration script is **NOT** automatically generated. If you make any changes or additions to the Postman environment used [tests/postman/postman_environment.docker-compose.json](./tests/postman/postman_environment.docker-compose.json), these changes must also be made to the configuration supplied to the apidoc generation script, [docs/swagger/postman-compose.env.json](./docs/swagger/postman-compose.env.json).

- Swagger UI configuration can be adjusted with [docs/swagger/swagger-config.json](./docs/swagger/swagger-config.json). See [swagger-ui/docs/usage/configuration.md](https://github.com/swagger-api/swagger-ui/blob/0b8de2c1796e67602bcbbc6d35c99cb167acf388/docs/usage/configuration.md) for the full list of configuration options.

## How To Deploy

### Deploying Develop and Stable Core API & Telemetry API

Deployments are done though [CI (Continuous Integration) scripts](./.github) using [Github Actions](https://docs.github.com/en/actions). The [core api](./api) and [telemetry api](./telemetry) are tested, built, and pushed to AWS ECR, where they **should** re-deploy on container push when the CI pipelines successfully finish (please check the redployment actually happens and force new deployment if not, it has been the case before that the deployment trigger config gets overwritten...).

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
| ------------------------------ | ------------------------------------------------------------------------ | ---------------------- |
| LOADER_POST_URL                | http://instrumentation-api_api_1/instrumentation/timeseries_measurements |                        |
| LOADER_API_KEY                 | appkey                                                                   |                        |
| ------------------------------ | ------------------------------------------------------------------------ | ---------------------- |
| LOADER_AWS_S3_ENDPOINT         | http://minio:9000                                                        | Used for local testing |
| LOADER_AWS_S3_REGION           | us-east-1                                                                |                        |
| LOADER_AWS_S3_DISABLE_SSL      | False                                                                    |                        |
| LOADER_AWS_S3_FORCE_PATH_STYLE | True                                                                     |                        |
| ------------------------------ | ------------------------------------------------------------------------ | ---------------------- |
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
