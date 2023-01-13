package timeseries

import "time"

// TimeWindow is a bounding box for time
type TimeWindow struct {
	After  time.Time `json:"after"`
	Before time.Time `json:"before"`
}

// TimeWindow.SetWindow sets the before and after of a time window
// If after or before are not provided return last 7 days of data from current time
func (tw *TimeWindow) SetWindow(after, before string) error {
	if after == "" || before == "" {
		tw.Before = time.Now()
		tw.After = tw.Before.AddDate(0, 0, -7)
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
