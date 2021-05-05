package main

import "fmt"

type TS struct {
	Time  int
	Value int
}

type TSPointer struct {
	Time        int
	Measurement *TS
}

func main() {
	a := []TS{
		{1, 5},
		{4, 8},
		{5, 9},
	}
	var normalized []TSPointer
	// Start Time, End Time of Time Window
	t, tEnd := 0, 5
	// Time Window Interval
	interval := 1

	wkIdx := 0

	for t <= tEnd {
		fmt.Printf("Working Index: %d; Working on time: %d; Comparing to: Time: %d ; Value: %d\n", wkIdx, t, a[wkIdx].Time, a[wkIdx].Value)
		if t < a[0].Time {
			fmt.Printf("time %d is below the minimum comparison time: %d; TS Value Literally Not Computable or Inferrable in Sane Way; Moving On\n", t, a[wkIdx].Time)
			t += interval
		}
		if t >= a[wkIdx].Time {
			fmt.Printf("time %d is equal to or greater than comparison time: %d\n", t, a[wkIdx].Time)
			// If Already at the end of the interpable array
			if wkIdx == len(a)-1 {
				fmt.Println("Already on the last index")
				fmt.Printf("Set Value for time %d; Time: %d; Value: %d\n", t, a[wkIdx].Time, a[wkIdx].Value)
				normalized = append(normalized, TSPointer{t, &a[wkIdx]})
				t += interval
				continue
			}
			fmt.Printf("Bump Working Index From %d --> %d\n", wkIdx, wkIdx+1)
			wkIdx += 1
			continue
		}

		normalized = append(normalized, TSPointer{t, &a[wkIdx-1]})
		t += interval
	}

	for _, v := range normalized {
		fmt.Printf("Computed Value: Time %d, Value %d; From Actual Measurement: Time %d, Value %d\n", v.Time, v.Measurement.Value, v.Measurement.Time, v.Measurement.Value)
	}
}
