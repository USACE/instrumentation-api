package timeseries

import "time"

// TimeWindow is a bounding box for time
type TimeWindow struct {
	After  time.Time `json:"after"`
	Before time.Time `json:"before"`
}
