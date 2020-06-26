package timeseries

func printit() string {
	return "hello from printit"
}

// Work in progress

// Shifter adjusts measurements using any number of adjusters
// Assumes []ZReference and []Measurement are already sorted on time.Time descending
// https://stackoverflow.com/questions/34329441/golang-struct-array-values-not-appending-in-loop
func Shifter(measurements []Measurement, shifts ...[]Measurement) ([]Measurement, error) {

	ss := measurements
	for _, a := range shifts {
		add(ss, a)
	}

	return ss, nil
}

func add(mm []Measurement, aa []Measurement) []Measurement {

	mIdx := 0
	for aIdx := range aa {
		for _, m := range mm[mIdx:] {
			if m.Time.Before(aa[aIdx].Time) {
				break
			}
			mm[mIdx].Value = m.Value + aa[aIdx].Value
			mIdx++
		}
	}
	// If there are additional measurements that fall outside the range of the adjuster measurements
	aValue := aa[len(aa)-1].Value
	for _, m := range mm[mIdx:] {
		mm[mIdx].Value = m.Value + aValue
	}

	return mm
}
