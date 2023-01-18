#!/usr/bin/env python3

# CLI tool for mocking POST requests sent from Campbell Scientific data loggers

import argparse
from datetime import datetime, timedelta
import time
import json
from urllib.error import HTTPError, URLError
from urllib.request import urlopen, Request
import random

from typing import Any

DEFAULT_BASE_URL = "http://localhost:8080"
DEFAULT_INTERVAL_SECONDS = 10


def post_data(url: str, data: str) -> Any | None:
    json_data = json.dumps(data).encode("utf-8")

    request = Request(url, headers={"Content-Type": "application/json"}, data=json_data)

    try:
        with urlopen(request, timeout=10) as response:
            return response
    except HTTPError as error:
        print("HTTPError", error.status, error.reason)
    except URLError as error:
        print("URLError", error.reason)
    except TimeoutError:
        print("Request timed out")

    return None  # exception was thrown


def create_test_data(interval: int) -> dict:
    measurement_time_1 = datetime.now().isoformat(timespec="seconds")
    measurement_time_2 = (
        datetime.now() - timedelta(seconds=(interval / 2))
    ).isoformat(timespec="seconds")

    return {
        "head": {
            "transaction": 0,
            "signature": 20883,
            "environment": {
                "station_name": "6239",
                "table_name": "Test",
                "model": "CR6",
                "serial_no": "6239",
                "os_version": "CR6.Std.12.01",
                "prog_name": "CPU:Updated_CR6_Sample_Template.CR6",
            },
            "fields": [
                {
                    "name": "batt_volt_Min",
                    "type": "xsd:float",
                    "units": "Volts",
                    "process": "Min",
                    "settable": False,
                },
                {
                    "name": "PanelT",
                    "type": "xsd:float",
                    "units": "Deg_C",
                    "process": "Smp",
                    "settable": False,
                },
            ],
        },
        "data": [
            {
                "time": measurement_time_1,
                "no": 0,
                "vals": [
                    round(random.uniform(11.50, 12.50), 2),
                    round(random.uniform(20.00, 25.00), 2),
                ],
            },
            {
                "time": measurement_time_2,
                "no": 0,
                "vals": [
                    round(random.uniform(11.50, 12.50), 2),
                    round(random.uniform(20.00, 25.00), 2),
                ],
            },
        ],
    }


def main() -> None:
    parser = argparse.ArgumentParser("telemetry timeseries uploader")
    parser.add_argument(
        "--base-url",
        help="base url to send requests",
        default=DEFAULT_BASE_URL,
        type=str,
    )
    parser.add_argument(
        "--interval",
        help="interval in seconds to wait between requests",
        default=DEFAULT_INTERVAL_SECONDS,
        type=int,
    )
    args = parser.parse_args()

    interval = args.interval
    url = f"{args.base_url}/telemetry/measurements"

    while True:
        data = create_test_data(interval)
        print(f"POST: {url}")
        print("request payload:", json.dumps(data, indent=2), sep="\n")
        res = post_data(url, data)
        print(
            f"\nresponse status: {res.status}",
            f"response body: {res.read().decode()}\n",
        )
        print(f"waiting {interval} seconds...\n")
        time.sleep(interval)


if __name__ == "__main__":
    main()
