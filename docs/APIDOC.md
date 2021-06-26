# Instrumentation API

Table of Contents
- [Instrumentation API](#instrumentation-api)
    - [Projects](#projects)
    - [Project Instruments](#project-instruments)
    - [Instruments](#instruments)
    - [Instrument Groups](#instrument-groups)
    - [Timeseries](#timeseries)
    - [Timeseries Measurements](#timeseries-measurements)

---
### Projects
- List Projects \
  [https://api.rsgis.dev/development/instrumentation/projects](https://api.rsgis.dev/development/instrumentation/projects)
- Get Project (Blue Water Dam Example Project) \
  [https://api.rsgis.dev/development/instrumentation/projects/5b6f4f37-7755-4cf9-bd02-94f1e9bc5984](https://api.rsgis.dev/development/instrumentation/projects/5b6f4f37-7755-4cf9-bd02-94f1e9bc5984)
- Update Project \
  `https://api.rsgis.dev/development/instrumentation/projects/5b6f4f37-7755-4cf9-bd02-94f1e9bc5984`
    - Example `PUT` body
        ```
        {
            "id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
            "federal_id": null,`{{base_url}}/instrumentation/projects/{id}/instruments`
            "name": "Blue Water Reservoir"
        }
        ```
---
### Project Instruments
- List Project Instruments \
  [https://api.rsgis.dev/development/instrumentation/projects/5b6f4f37-7755-4cf9-bd02-94f1e9bc5984/instruments](https://api.rsgis.dev/development/instrumentation/projects/5b6f4f37-7755-4cf9-bd02-94f1e9bc5984/instruments)

---
### Instruments
- List Instruments \
  [https://api.rsgis.dev/development/instrumentation/instruments](https://api.rsgis.dev/development/instrumentation/instruments)
- Update Instrument (for Demo Pz #1) \
  `https://api.rsgis.dev/development/instrumentation/instruments/a7540f69-c41e-43b3-b655-6e44097edb7e`
    - Example `PUT` body
        ```
        {
            "id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "status_id": "94578354-ffdf-4119-9663-6bd4323e58f5",
            "status": "destroyed",
            "status_time": "2001-01-01T00:00:00Z",
            "slug": "demo-piezometer-1",
            "name": "Demo Piezometer 1 Updated Name",
            "type_id": "1bb4bf7c-f5f8-44eb-9805-43b07ffadbef",
            "type": "Piezometer",
            "geometry": {
                "type": "Point",
                "coordinates": [
                    -80.8,
                    26.7
                ]
            },
            "station": null,
            "offset": null,
            "project_id": "5b6f4f37-7755-4cf9-bd02-94f1e9bc5984",
            "zreference": 44.5,
            "zreference_datum_id": "72113f9a-982d-44e5-8fc1-8e595dafd344",
            "zreference_datum": "North American Vertical Datum of 1988 (NAVD 88)",
            "zreference_time": "2006-06-01T00:00:00Z"
        }
        ```
---
### Instrument Groups
- List Instrumentation Groups \
  [https://api.rsgis.dev/development/instrumentation/instrument_groups](https://api.rsgis.dev/development/instrumentation/instrument_groups)

---
### Timeseries
- List Timeseries \
  [https://api.rsgis.dev/development/instrumentation/timeseries](https://api.rsgis.dev/development/instrumentation/timeseries)
- Get Timeseries Metadata (for Demo Pz #1) \
  [https://api.rsgis.dev/development//instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9](https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9)

- Create Single Timeseries \
  `https://api.rsgis.dev/development/instrumentation/timeseries`
    - Example `POST` body
        ```
        {
            "name": "Test Timeseries 4",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
            "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
        }
        ```
- Create Mutiple Timeseries \
  `https://api.rsgis.dev/development/instrumentation/timeseries`
    - Example `POST` body
        ```
        [{
            "name": "Test Timeseries 5",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
            "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
        },
        {
            "name": "Test Timeseries 6",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
            "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
        },
        {
            "name": "Test Timeseries 7",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
            "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce"
        }]
        ```
- Update Timeseries \
  `https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9`
    - Example `PUT` body
        ```
        {
            "id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
            "slug": "test-timeseries-1",
            "name": "New Name for Test Timeseries 1",
            "instrument_id": "a7540f69-c41e-43b3-b655-6e44097edb7e",
            "instrument": "Demo Piezometer 1",
            "parameter_id": "068b59b0-aafb-4c98-ae4b-ed0365a6fbac",
            "parameter": "stage",
            "unit_id": "f777f2e2-5e32-424e-a1ca-19d16cd8abce",
            "unit": "feet"
        }
        ```
---
### Timeseries Measurements
- Get Timeseries Measurements (for Demo Pz #1) \
  [https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/measurements?after=1900-01-01T00:00:00.00Z&before=2021-01-01T00:00:00.00Z](https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/measurements?after=1900-01-01T00:00:00.00Z&before=2021-01-01T00:00:00.00Z)
- Delete Timeseries Measurements \
  [https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/measurements?time=1900-01-01T00:00:00.00Z](https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/measurements?time=1900-01-01T00:00:00.00Z)
- Create Timeseries Measurement(s) for **Single Timeseries** \
  [https://api.rsgis.dev/development/instrumentation/timeseries/measurements](https://api.rsgis.dev/development/instrumentation/timeseries/measurements)
    - Example `POST` body
        ```
        {
            "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
            "items": [
                    {"time": "2020-06-01T00:00:00Z", "value": 10.00},
                    {"time": "2020-06-02T01:00:00Z", "value": 11.10},
                    {"time": "2020-06-03T02:00:00Z", "value": 10.20},
                    {"time": "2020-06-04T03:00:00Z", "value": 10.30},
                    {"time": "2020-06-05T04:00:00Z", "value": 10.40}
                ]
        }
        ```
- Create Timeseries Measurement(s) for **Multiple Timeseries** \
  [https://api.rsgis.dev/development/instrumentation/timeseries/measurements](https://api.rsgis.dev/development/instrumentation/timeseries/measurements)
    - Example `POST` body
        ```
        [
            {
                "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
                "items": [
                {"time": "2020-06-01T00:00:00Z", "value": 10.00},
                    {"time": "2020-06-02T01:00:00Z", "value": 11.10},
                    {"time": "2020-06-03T02:00:00Z", "value": 10.20},
                    {"time": "2020-06-04T03:00:00Z", "value": 10.30},
                    {"time": "2020-06-05T04:00:00Z", "value": 10.40}
            ]
            },
            {
                "timeseries_id": "9a3864a8-8766-4bfa-bad1-0328b166f6a8",
                "items": [
                {"time": "2020-06-01T00:00:00Z", "value": 10.00},
                    {"time": "2020-06-02T01:00:00Z", "value": 11.10},
                    {"time": "2020-06-03T02:00:00Z", "value": 10.20},
                    {"time": "2020-06-04T03:00:00Z", "value": 10.30},
                    {"time": "2020-06-05T04:00:00Z", "value": 10.40}
            ]
            },
            {
                "timeseries_id": "7ee902a3-56d0-4acf-8956-67ac82c03a96",
                "items": [
                {"time": "2020-06-01T00:00:00Z", "value": 10.00},
                    {"time": "2020-06-02T01:00:00Z", "value": 11.10},
                    {"time": "2020-06-03T02:00:00Z", "value": 10.20},
                    {"time": "2020-06-04T03:00:00Z", "value": 10.30},
                    {"time": "2020-06-05T04:00:00Z", "value": 10.40}
            ]
            }
        ]
        ```
---
### Inclinometer Measurements
- Get Inclinometer Measurements (for Demo Pz #1) \
  [https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/inclinometer_measurements?after=1900-01-01T00:00:00.00Z&before=2021-01-01T00:00:00.00Z](https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/inclinometer_measurements?after=1900-01-01T00:00:00.00Z&before=2021-01-01T00:00:00.00Z)
- Delete Inclinometer Measurements \
  [https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/inclinometer_measurements?time=1900-01-01T00:00:00.00Z](https://api.rsgis.dev/development/instrumentation/timeseries/869465fc-dc1e-445e-81f4-9979b5fadda9/inclinometer_measurements?time=1900-01-01T00:00:00.00Z)
  - Create Inclinometer Measurement(s) for **Multiple Inclinometers** \
  [https://api.rsgis.dev/development/instrumentation/timeseries/measurements](https://api.rsgis.dev/development/instrumentation/timeseries/measurements)
    - Example `POST` body
        ```
        [
            {
                "timeseries_id": "869465fc-dc1e-445e-81f4-9979b5fadda9",
                "inclinometers": [
                    {
                      "time": "2021-06-17T00:00:00Z", 
                      "values": [
                            {
                              "depth": 106, 
                              "a0": 590,
                              "a180": -562,
                              "b0": -142,
                              "b180": 176
                            },
                            {
                              "depth": 108, 
                              "a0": 614,
                              "a180": -586,
                              "b0": 107,
                              "b180": -149
                            },
                            {
                              "depth": 110, 
                              "a0": 622,
                              "a180": -592,
                              "b0": -67,
                              "b180": 107
                            },
                            {
                              "depth": 112, 
                              "a0": 623,
                              "a180": -598,
                              "b0": 8,
                              "b180": -48
                            },
                            {
                              "depth": 114, 
                              "a0": 606,
                              "a180": -577,
                              "b0": 124,
                              "b180": -72
                            },
                            {
                              "depth": 116, 
                              "a0": 0,
                              "a180": 0,
                              "b0": 0,
                              "b180": 0
                            }
                      ]
                    }
                ]
            }
        ]
        ```