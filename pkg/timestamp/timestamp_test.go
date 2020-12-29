package timestamp

import (
	"testing"
)

// Test unix time string with basetime 1500000000
func TestGetTimeFromUnixTime(t *testing.T) {
	var basetime int64
	basetime = 1500000000

	ts := Timestamp{}

	time := ts.GetTimeFromUnixTime(basetime)

	if time != "Fri, 14 Jul 2017 02:40:00 UTC" {
		t.Errorf("ts = \t\"%v\";\twant\t\"Fri, 14 Jul 2017 02:40:00 UTC\"", time)
	}
}
