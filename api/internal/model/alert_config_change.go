package model

import (
	"encoding/json"
	"fmt"
)

type AlertConfigChangeOpts struct {
	RateOfChange float64 `json:"rate_of_change" db:"rate_of_change"`
}

func (o *AlertConfigChangeOpts) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return fmt.Errorf("type assertion failed")
	}
	return json.Unmarshal([]byte(b), o)
}
