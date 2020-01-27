package asec

import (
	"time"
)

// BeginningDate of the era
var BeginningDate time.Time = time.Date(2020, 1, 5, 17, 55, 0, 0, time.FixedZone("UTC+3", 3*60*60))

// DaysSince returns days duration from beginning date
func DaysSince() int {
	duration := time.Now().Sub(BeginningDate).Hours()
	return int(duration / 24)
}
