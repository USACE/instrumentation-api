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

DEFAULT_BASE_URL = "http://telemetry"
DEFAULT_INTERVAL_SECONDS = 10
MOCK_API_KEY = "8pszF58y7Hpwr8DgR9UYhovcjJYdBhRSMt9dGX1RBmdj6WtH4NUNFao"
MODEL = "CR6"
SN = "12345"


def post_data(url: str, data: str) -> Any | None:
    json_data = json.dumps(data).encode("utf-8")

    request = Request(url, headers={"Content-Type": "application/json", "X-Api-Key": MOCK_API_KEY}, data=json_data)

    try:
        with urlopen(request, timeout=10) as response:
            print("response status:", response.status)
            # resObj = json.loads(response.read())
            # print("body: ", json.dumps(resObj, indent=2), sep="\n")
            return response
    except HTTPError as error:
        print("HTTPError", error.status, error.reason, error.read())
    except URLError as error:
        print("URLError", error.reason)
    except TimeoutError:
        print("Request timed out")

    return None  # exception was thrown


def create_test_data(interval: int, idx: int) -> dict:
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
                "model": MODEL,
                "serial_no": SN,
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
                "no": idx,
                "vals": [
                    round(random.uniform(11.50, 12.50), 2),
                    round(random.uniform(20.00, 25.00), 2),
                ],
            },
            {
                "time": measurement_time_2,
                "no": idx,
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
    url = f"{args.base_url}/telemetry/datalogger/{MODEL}/{SN}"

    i = 0
    while True:
        data = create_test_data(interval, i)
        i += 1
        print(f"POST: {url}")
        # print("request payload:", json.dumps(data, indent=2), sep="\n")
        post_data(url, data)
        print(f"waiting {interval} seconds...\n")
        time.sleep(interval)


if __name__ == "__main__":
    main()