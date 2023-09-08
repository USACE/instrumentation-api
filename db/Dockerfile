ARG POSTGRES_VERSION
FROM postgres:${POSTGRES_VERSION:-14}

RUN apt update
RUN apt install postgis postgresql-14-postgis-3 -y
