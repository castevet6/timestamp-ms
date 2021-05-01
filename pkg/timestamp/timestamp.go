package timestamp

import (
	"errors"
	"fmt"
	_ "strconv"
	"time"
)

const shortForm1 = "2006-Jan-02"
const shortForm2 = "2006-01-02"

type Timestamp struct{}

func (ts *Timestamp) GetTimeNow() (string, int64) {
	utc_t := time.Now().UTC()
	unix_t := utc_t.Unix()

	return utc_t.Format(time.RFC1123), unix_t
}

func (ts *Timestamp) GetTimeFromDateString(s string) (string, int64, error) {
	// Try to parse the date string directly to Time type
	if t, err := tryParseDateString(s); err == nil {
		// valid datestring
		return t.Format(time.RFC1123), t.Unix(), nil
	}

	// not valid datestring
	return "", 0, errors.New("invalid date string layout")
}

func (ts *Timestamp) GetTimeFromUnixTime(i int64) string {
	t := time.Unix(i, 0).UTC()
	fmt.Println(t.Format(time.RFC1123))
	return t.Format(time.RFC1123)
}

func tryParseDateString(s string) (time.Time, error) {
	var (
		t   time.Time
		err error
	)

	/* Parse custom layouts */
	if t, err = time.Parse(shortForm1, s); err == nil {
		return t, nil
	}

	if t, err = time.Parse(shortForm2, s); err == nil {
		return t, nil
	}

	if t, err = time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}

	return time.Time{}, errors.New("Invalid datestring")
}
