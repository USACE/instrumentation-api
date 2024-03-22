package model

import "time"

// TimeWindow is a bounding box for time
type TimeWindow struct {
	After  time.Time `json:"after"`
	Before time.Time `json:"before"`
}

// SetWindow sets the before and after of a time window
// If after or before are not provided return last 7 days of data from current time
func (tw *TimeWindow) SetWindow(after, before string, defaultStart, defaultEnd time.Time) error {
	if after == "" || before == "" {
		tw.After = defaultStart
		tw.Before = defaultEnd
	} else {
		// Attempt to parse query param "after"
		tA, err := time.Parse(time.RFC3339, after)
		if err != nil {
			return err
		}
		tw.After = tA
		// Attempt to parse query param "before"
		tB, err := time.Parse(time.RFC3339, before)
		if err != nil {
			return err
		}
		tw.Before = tB
	}
	return nil
}
