ARG KC_BASE_IMAGE=quay.io/keycloak/keycloak:20.0.3

FROM ${KC_BASE_IMAGE} AS builder

ENV KC_HEALTH_ENABLED=true
ENV KC_METRICS_ENABLED=true

WORKDIR /opt/keycloak

COPY server.keystore /conf/server.keystore
COPY direct-grant.jar /opt/keycloak/providers/direct-grant.jar
COPY login-gov.jar /opt/keycloak/providers/login-gov.jar

RUN /opt/keycloak/bin/kc.sh build

FROM ${KC_BASE_IMAGE}

COPY --from=builder /opt/keycloak/ /opt/keycloak/
COPY server.keystore /conf/server.keystore
COPY ./realm.json /opt/keycloak/data/import/exported-realm.json

ENTRYPOINT [ "/opt/keycloak/bin/kc.sh" ]
