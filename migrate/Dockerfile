ARG BASE_IMAGE=flyway/flyway:8.5.9
FROM $BASE_IMAGE

COPY . /flyway/sql/

# paths / user / pass / jdbc url set in task definition
ENV FLYWAY_EDITION=community
ENV FLYWAY_CONNECT_RETRIES=60
ENV FLYWAY_SCHEMAS=midas,public

ENTRYPOINT [ "flyway" ]

CMD [ "migrate" ]
