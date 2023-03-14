FROM python:3.11-slim-buster

ENV PYTHONUNBUFFERED=1

COPY ./telemetry_post_tester.py /app/telemetry_post_tester.py
WORKDIR /app

ENTRYPOINT [ "./telemetry_post_tester.py" ]
