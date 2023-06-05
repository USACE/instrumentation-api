FROM ubuntu

RUN apt update && apt upgrade -y; \
    apt install -y curl

COPY ./mock_check_alerts.sh /app/mock_check_alerts.sh
WORKDIR /app

ENTRYPOINT ["./mock_check_alerts.sh"]
