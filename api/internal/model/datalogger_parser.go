package model

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"math"
	"os"
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
	case `"NAN"`:
		*j = FloatNanInf(math.NaN())
	case `"INF"`:
		*j = FloatNanInf(math.Inf(1))
	default:
		var fv float64
		if err := json.Unmarshal(v, &fv); err != nil {
			return err
		}
		*j = FloatNanInf(fv)
	}
	return nil
}

// ParseTOA5 parses a Campbell Scientific TOA5 data file that is simlar to a csv.
// The unique properties of TOA5 are that the meatdata are stored in header of file (first 4 lines of csv)
func ParseTOA5(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// read headers
	// LINE 1: information about the data logger (e.g. serial number and program name)
	header1, err := r.Read()
	if err != nil {
		return nil, err
	}
	// LINE 2: data header (names of the variables stored in the table)
	header2, err := r.Read()
	if err != nil {
		return nil, err
	}
	// LINE 3: units for the variables if they have been defined in the data logger
	header3, err := r.Read()
	if err != nil {
		return nil, err
	}
	// LINE 4: abbreviation for processing data logger performed
	// (e.g. sample, average, standard deviation, maximum, minimum, etc.)
	header4, err := r.Read()
	if err != nil {
		return nil, err
	}
	log.Printf("header1: %v", header1)
	log.Printf("header2: %v", header2)
	log.Printf("header3: %v", header3)
	log.Printf("header4: %v", header4)

	// continue read until EOF
	data, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	log.Printf("data: %v", data)

	return data, nil
}
