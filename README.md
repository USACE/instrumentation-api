# Instrumentation API  ![Build API and Lambda Package](https://github.com/rsgis-dev/instrumentation-api/workflows/Build%20API%20and%20Lambda%20Package/badge.svg)

An Application Programming Interface (API) to manage instrumentation data, built with Golang and Deployed on AWS Lambda.

# How to Develop

## Running a Database for Local Development

1. Install Docker and Docker Compose

2. Change to the /database directory in this repository and type `docker-compose up`. This brings up two services on `localhost`

   1. A postgres database with postgis schema installed

   2. pgadmin4 (a user interface to interact with the database)

   To modify the database using pgadmin4, open a web browser and go to `localhost:8080`.

   Login with `Email:postgres@postgres.com` and `Password:postgres` respectively.

   Create a database connection to the postgres database by right-clicking `servers --> create --> server` in the left menu tree. Enter the following information and click `save`.

   **General Tab**

   | Field | Value     |
   | ----- | --------- |
   | Name  | localhost |

   **Connection Tab**

   | Field             | Value          |
   | ----------------- | -------------- |
   | Host name/address | postgres       |
   | Port              | 5432 (default) |
   | Username          | postgres       |
   | Password          | postgres       |

3. Initialize the database and seed it with some data

   Use the Query Tool in pgadmin4 and the .sql files in the database/ directory in this repository. You can find the query tool by expanding the left menu tree to `Servers --> Databases --> postgres`. Right click `postgres --> Query Tool`. From here you can copy .sql into the Query Tool and run it by pressing `f5`. Note: to only run a portion of the SQL you've copied, you can highlight the section you want to run before hitting `f5`.

   The .sql in the files should be run in this order to initialize the database:

   1. tables.sql

   1. seed_data.sql

## Running the GO API for Local Development

Either of these options starts the API at `localhost:3030`. The API uses JSON Web tokens (JWT) for authorization by default.  To disable JWT for testing or development, you can set the environment variable `JWT_DISABLED=TRUE`.

**With Visual Studio Code Debugger**

You can use the launch.json file in this repository in lieu of `go run root/main.go` to run the API in the VSCode debugger.  This takes care of the required environment variables to connect to the database.

**Without Visual Studio Code Debugger**

Set the following environment variables and type `go run root/main.go` from the top level of this repository.

    * DB_USER=postgres
    * DB_PASS=postgres
    * DB_NAME=postgres
    * DB_HOST=localhost
    * DB_SSLMODE=disable
    * LAMBDA=FALSE
    * JWT_DISABLED=FALSE

Note: make sure environment variable `LAMBDA` is either **not set** or is set to `LAMBDA=FALSE`.

## Running API Docs Locally

From the top level of this repository, type `make docs`. This starts a container that serves content based on "apidoc.yml" in this repository.
Open a browser and navigate to `https://localhost:4000` to view the content.

# How To Deploy

## Postgres Database on AWS Relational Database Service (RDS)

Database should be initialized with the following SQL files in the order listed:

1. rds_init.sql (install PostGIS extensions)

1. tables.sql (create tables for application)

1. roles.sql (database roles, grants, etc.)

   Note: Change 'password' in roles.sql to a real password for the `instrumentation_user` account.

## AWS Lambda API Function

## AWS API Gateway

