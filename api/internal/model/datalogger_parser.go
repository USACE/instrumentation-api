package model

import (
	"encoding/json"
	"math"
)

type DataloggerPayload struct {
	Head Head    `json:"head"`
	Data []Datum `json:"data"`
}

type Datum struct {
	Time string        `json:"time"`
	No   int64         `json:"no"`
	Vals []FloatNanInf `json:"vals"`
}

type Head struct {
	Transaction int64       `json:"transaction"`
	Signature   int64       `json:"signature"`
	Environment Environment `json:"environment"`
	Fields      []Field     `json:"fields"`
}

type Environment struct {
	StationName string `json:"station_name"`
	TableName   string `json:"table_name"`
	Model       string `json:"model"`
	SerialNo    string `json:"serial_no"`
	OSVersion   string `json:"os_version"`
	ProgName    string `json:"prog_name"`
}

type Field struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Units    string `json:"units"`
	Process  string `json:"process"`
	Settable bool   `json:"settable"`
}

type FloatNanInf float64

func (j *FloatNanInf) UnmarshalJSON(v []byte) error {
	switch string(v) {
	case `"NAN"`, "NAN":
		*j = FloatNanInf(math.NaN())
	case `"INF"`, "INF":
		*j = FloatNanInf(math.Inf(1))
	default:
		var fv float64
		if err := json.Unmarshal(v, &fv); err != nil {
			*j = FloatNanInf(math.NaN())
			return nil
		}
		*j = FloatNanInf(fv)
	}
	return nil
}
