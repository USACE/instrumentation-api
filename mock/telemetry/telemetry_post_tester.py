#!/usr/bin/env python3

import argparse
import json
import random
import time
from datetime import datetime, timedelta
from getpass import getpass
from typing import Any
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen

# the follow parameters can be set here or overwritten with command line flags
# this is an api key used local testing. You will be promter for the api key
# on program start if the --use-mock-api-key flag is not present
DEFAULT_API_KEY = "8pszF58y7Hpwr8DgR9UYhovcjJYdBhRSMt9dGX1RBmdj6WtH4NUNFao"
DEFAULT_USE_MOCK_API_KEY = False
DEFAULT_BASE_URL = "http://telemetry"
DEFAULT_INTERVAL_SECONDS = 60
DEFAULT_MODEL = "CR6"
DEFAULT_SN = "12345"
DEFAULT_TABLE_NAME = "Demo Datalogger Table"
DEFAULT_VERBOSE = False


def post_data(url: str, data: dict, api_key: str) -> Any | None:
    json_data = json.dumps(data).encode("utf-8")

    request = Request(
        url,
        headers={"Content-Type": "application/json", "X-Api-Key": api_key},
        data=json_data,
    )

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


def create_test_data(interval: int, idx: int, model: str, sn: str, table_name: str) -> dict:
    measurement_time_1 = datetime.now().isoformat(timespec="seconds")
    measurement_time_2 = (datetime.now() - timedelta(seconds=(interval / 2))).isoformat(
        timespec="seconds"
    )

    return {
        "head": {
            "transaction": 0,
            "signature": 20883,
            "environment": {
                "station_name": "6239",
                "table_name": table_name,
                "model": model,
                "serial_no": sn,
                "os_version": f"{model}.Std.12.01",
                "prog_name": f"CPU:Updated_{model}_Sample_Template.{model}",
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
                    "NAN",
                    "INF",
                ],
            },
        ],
    }


def main() -> None:
    parser = argparse.ArgumentParser(
        prog="Mock Data Logger",
        description="""
        CLI tool for mocking POST requests sent from Campbell Scientific data loggers.
        Default flags are for use with local docker compose stack""",
        usage="./telemetry_post_tester.py"
        + " [-h | --help] [--base-url url] [--interval n] [--model model]"
        + " [--sn serial_number] [--use-mock-api-key] [--verbose]",
        formatter_class=argparse.ArgumentDefaultsHelpFormatter,
    )
    parser.add_argument(
        "--base-url",
        help="base url to send requests to, e.g. https://midas-telemetry.sec.usace.army.mil",
        default=DEFAULT_BASE_URL,
        type=str,
    )
    parser.add_argument(
        "--interval",
        help="interval in seconds to wait between requests",
        default=DEFAULT_INTERVAL_SECONDS,
        type=int,
    )
    parser.add_argument(
        "--model",
        help="valid model of the registered data logger",
        default=DEFAULT_MODEL,
        type=str,
    )
    parser.add_argument(
        "--sn",
        help="valid serial number of the registered data logger",
        default=DEFAULT_SN,
        type=str,
    )
    parser.add_argument(
        "--table",
        help="table name for payload",
        default=DEFAULT_TABLE_NAME,
        type=str,
    )
    parser.add_argument(
        "--use-mock-api-key",
        help="[optional] use default mock api key for local testing",
        default=DEFAULT_USE_MOCK_API_KEY,
        type=bool,
        action=argparse.BooleanOptionalAction,
    )
    parser.add_argument(
        "--multi-table",
        help="[optional] mock an additional payload sent to another specified table",
        default=DEFAULT_USE_MOCK_API_KEY,
        type=bool,
        action=argparse.BooleanOptionalAction,
    )
    parser.add_argument(
        "--verbose",
        help="[optional] show outgoing the request's mocked payloads",
        default=DEFAULT_VERBOSE,
        type=bool,
        action=argparse.BooleanOptionalAction,
    )
    args = parser.parse_args()

    if not args.use_mock_api_key:
        api_key = getpass(prompt="Enter API Key: ")
    else:
        api_key = DEFAULT_API_KEY

    url = f"{args.base_url}/telemetry/datalogger/{args.model}/{args.sn}"

    i = 0
    while True:
        data = create_test_data(args.interval, i, args.model, args.sn, args.table)
        print(f"POST: {url}")

        if args.verbose:
            print("request payload:", json.dumps(data, indent=2), sep="\n")

        post_data(url, data, api_key)

        print(f"waiting {args.interval} seconds...\n")
        time.sleep(args.interval)

        if args.multi_table:
            data = create_test_data(args.interval, i, args.model, args.sn, "Demo Multi-Table")
            print(f"POST: {url}")

            if args.verbose:
                print("request payload:", json.dumps(data, indent=2), sep="\n")

            post_data(url, data, api_key)
            print(f"waiting {args.interval} seconds...\n")
            time.sleep(args.interval)

        i += 1

if __name__ == "__main__":
    main()
