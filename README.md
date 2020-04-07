# Instrumentation API

An Application Programming Interface (API) to manage instrumentation data, built with Golang and Deployed on AWS Lambda.

# How to Develop

## After Cloning the Repository

1. Install serverless command line tools with `npm install -g serverless`. Note: This is a "global" `-g` install.

2. Install node packages associated with this project. From the top directory of this repository, type `npm install`.

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

   1. postgis_init.sql

   2. tables.sql

   3. seed_data.sql

## Running the GO API for Local Development

### Without AWS Lambda environment simulated (use this most of the time)

Either of these options starts the API at `localhost:3030`

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

Note: make sure environment variable `LAMBDA` is either **not set** or is set to `LAMBDA=FALSE`.

### With AWS Lambda environment simulation

1. Set environment variable `LAMBDA=TRUE`. Note: `True` is case insensitive, so True, tRue or true all work.

2. From the top level of this repository, type `make local`. This starts the API at  `localhost:3030/development/`.  This runs the API in a docker container that simulates the AWS Lambda environment and request payload from the AWS API Gateway.

## Running API Docs Locally

From the top level of this repository, type `make docs`. This starts a container that serves content based on "apidoc.yml" in this repository.
Open a browser and navigate to `https://localhost:4000` to view the content.


# How To Deploy

## Postgres Database on AWS Relational Database Service (RDS)

Database should be initialized with the following SQL files in the order listed:

1. pre_postgis_init_rds.sql (install PostGIS extensions)

2. postgis_init.sql (install other extensions)

3. tables.sql (create tables for application)

3. roles.sql (database roles, grants, etc.)

   Note: Change 'password' in roles.sql to a real password for the `instrumentation_user` account.

## AWS Lambda API Function

## AWS API Gateway

