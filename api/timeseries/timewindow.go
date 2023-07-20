package timeseries

import "time"

// TimeWindow is a bounding box for time
type TimeWindow struct {
	Start time.Time `json:"after"`
	End   time.Time `json:"before"`
}

// TimeWindow.SetWindow sets the before and after of a time window
// If after or before are not provided return last 7 days of data from current time
func (tw *TimeWindow) SetWindow(start, end string, defaultStart, defaultEnd time.Time) error {
	if start == "" || end == "" {
		tw.Start = defaultStart
		tw.End = defaultEnd
	} else {
		// Attempt to parse query param "after"
		tA, err := time.Parse(time.RFC3339, start)
		if err != nil {
			return err
		}
		tw.Start = tA
		// Attempt to parse query param "before"
		tB, err := time.Parse(time.RFC3339, end)
		if err != nil {
			return err
		}
		tw.End = tB
	}
	return nil
}
