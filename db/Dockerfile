ARG POSTGRES_VERSION
FROM postgres:${POSTGRES_VERSION:-14}

RUN apt-get update && apt-get install postgis postgresql-14-postgis-3 -y
