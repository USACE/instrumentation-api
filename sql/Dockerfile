FROM flyway/flyway

COPY . /flyway/sql/

ENV FLYWAY_EDITION=community
ENV FLYWAY_CONNECT_RETRIES=60
ENV FLYWAY_SCHEMAS=midas

ENTRYPOINT [ "flyway" ]

CMD [ "migrate" ]
