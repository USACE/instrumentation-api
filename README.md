# Instrumentation API  ![Build API and Lambda Package](https://github.com/usace/instrumentation-api/workflows/Build%20API%20and%20Lambda%20Package/badge.svg)

An Application Programming Interface (API) to manage instrumentation data, built with Golang and Deployed on AWS Lambda.

# Documentation

Documentation for the API is maintained in a Markdown file held at [`docs/APIDOC.md`](./docs/APIDOC.md). A [Postman](https://www.postman.com/api-documentation-tool/) documentation and testing environment is also maintained at [`tests/postman_environment.local`](./tests/postman_environment.local.json).

# How to Develop

## Running a Database for Local Development

1. Install, at a minimum, [Docker Engine](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/). In place of these two, one can install [Docker Desktop](https://docs.docker.com/desktop/).

2. Copy the `.env.example` file to `.env` (e.g., `cp .env.example .env`). This provides configuration options to Docker Compose.

3. Change to the /database directory in this repository and type `docker-compose up`. This brings up two services on `localhost`

   1. A postgres database with postgis schema installed using the Docker image [mdillon/postgis](https://hub.docker.com/r/mdillon/postgis/)

   2. pgadmin4 (a user interface to interact with the database) using the Docker image [dpage/pgadmin4](https://hub.docker.com/r/dpage/pgadmin4/)

   To modify the database using pgadmin4, open a web browser and go to `http://localhost:8080`, or whichever port is set in `.env`.

   Login with `Email:postgres@postgres.com` and `Password:postgres` respectively.

   Create a database connection to the postgres database by right-clicking `servers --> register --> server` in the left menu tree. Enter the following information and click `save`.

   **General Tab**

   | Field | Value                               |
   | ----- | ----------------------------------- |
   | Name  | localhost (or other preferred name) |

   **Connection Tab**

   | Field             | Value          |
   | ----------------- | -------------- |
   | Host name/address | postgres       |
   | Port              | 5432 (default) |
   | Username          | postgres       |
   | Password          | postgres       |

4. Initialize the database and seed it with some data (docker-compose runs this for you)

   Use the Query Tool in pgadmin4 and the .sql files in the database/ directory in this repository. You can find the query tool by expanding the left menu tree to `Servers --> Databases --> postgres`. Right click `postgres --> Query Tool`. From here, copy [tables.sql](database/sql/10-tables.sql) into the Query Tool and run it by pressing `f5`. Note: to only run a portion of the SQL you've copied, you can highlight the section you want to run before hitting `f5`.

## Running the GO API for Local Development

Either of these options starts the API at `localhost:3030`. The API uses JSON Web tokens (JWT) for authorization by default.  To disable JWT for testing or development, you can set the environment variable `JWT_DISABLED=TRUE`.

**With Visual Studio Code Debugger**

You can use the launch.json file in this repository in lieu of `go run main.go` to run the API in the VSCode debugger.  This takes care of the required environment variables to connect to the database.

**Without Visual Studio Code Debugger**

Set the following environment variables and type `go run main.go` from the top level of this repository.

    * INSTRUMENTATION_DB_USER=postgres
    * INSTRUMENTATION_DB_PASS=postgres
    * INSTRUMENTATION_DB_NAME=postgres
    * INSTRUMENTATION_DB_HOST=localhost
    * INSTRUMENTATION_DB_SSLMODE=disable
    * LAMBDA=FALSE
    * JWT_DISABLED=FALSE

Note: When running the API locally, make sure environment variable `LAMBDA` is either **not set** or is set to `LAMBDA=FALSE`.

## Running Tests

Regression tests are maintained for the project in the [aforementioned](#documentation) [Postman](https://www.postman.com/api-documentation-tool/) environments. They are run automatically by GitHub Actions through the script `test.sh`.

In both cases, the Postman environment regression tests are run, then output. If the environment variable `REPORT` is set to `true`, then this output is sent to an HTML file. Otherwise, it is printed to the caller's stdout.

# How To Deploy

## Postgres Database on AWS Relational Database Service (RDS)

Database should be initialized with the following SQL files in the order listed:

1. rds_init.sql (install PostGIS extensions)

1. tables.sql (create tables for application)

1. roles.sql (database roles, grants, etc.)

   Note: Change 'password' in roles.sql to a real password for the `instrumentation_user` account.

# How to Update

Updating an instance of `instrumentation-api` is trivially completed by rebuilding the Docker container used by it, then restarting the service.

If a postgres database has already been created and is in use, updates are less trivial. Before rebuilding and restarting the aforementioned API instance, database migrations must be carried out **manually**. Snippets for doing so are supplied in [`database/snippets`](./database/snippets).
