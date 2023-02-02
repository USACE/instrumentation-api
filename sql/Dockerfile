FROM flyway/flyway

COPY . /flyway/sql/

# user / pass / jdbc url set in task definition
ENV FLYWAY_EDITION=community
ENV FLYWAY_CONNECT_RETRIES=60
ENV FLYWAY_SCHEMAS=midas

ENTRYPOINT [ "flyway" ]

CMD [ "migrate" ]
