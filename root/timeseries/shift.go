// package timeseries

// Work in progress

// import (
// 	"api/root/models"
// )

// // Shifter adjusts timeseries measurements using any number of adjusters
// // Assumes []ZReference and []TimeseriesMeasurement are already sorted on time.Time descending
// // https://stackoverflow.com/questions/34329441/golang-struct-array-values-not-appending-in-loop
// func Shifter(measurements []models.TimeseriesMeasurement, shifts ...[]models.TimeseriesMeasurement) ([]models.TimeseriesMeasurement, error) {

// 	ss := measurements
// 	for _, a := range shifts {
// 		add(ss, a)
// 	}

// 	return ss, nil
// }

// func add(mm []models.TimeseriesMeasurement, aa []models.TimeseriesMeasurement) []models.TimeseriesMeasurement {

// 	mIdx := 0
// 	for aIdx := range aa {
// 		for _, m := range mm[mIdx:] {
// 			if m.Time.Before(aa[aIdx].Time) {
// 				break
// 			}
// 			mm[mIdx].Value = m.Value + aa[aIdx].Value
// 			mIdx++
// 		}
// 	}
// 	// If there are additional measurements that fall outside the range of the adjuster measurements
// 	aValue := aa[len(aa)-1].Value
// 	for _, m := range mm[mIdx:] {
// 		mm[mIdx].Value = m.Value + aValue
// 	}

// 	return mm
// }
