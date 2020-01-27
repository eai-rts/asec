package asec

import (
	"github.com/uniplaces/carbon"
)

// BEGINNING of the era
const BEGINNING = "2020-01-05"

var beginning := carbon.Parse(carbon.DateFormat, "2000-08-20")

// DaysCounter counts how many days sinсe the beginning of era 
type DaysCounter interface {
	DaysSince()
}