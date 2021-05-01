package main

import (
	"encoding/json"
	"net/http"
    "strconv"
)

type timestampObj struct {
	Utc  string `json:"utc"`
	Unix int64  `json:"unix"`
}

type errorObj struct {
    Error   string  `json:"error"`
}

func (app *application) getTimestamp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// call timestamp package to get current time
	timeNow, unixTimeNow := app.timestamp.GetTimeNow()

	tsObj := &timestampObj{
		Utc:  timeNow,
		Unix: unixTimeNow,
	}

	b, err := json.Marshal(tsObj)

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Write(b)
}

func (app *application) getTimestampFromQueryString(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    // timeformat
    queryString := ps.ByName("timeformat")

    // check if is 10 digit int i.e. unix/epoch time value
    if unixval, err := strconv.ParseInt(queryString, 10, 64); len(queryString) == 10 && err == nil {
        // GetTimeFromUnixTime
        tsObj := &timestampObj{}
        tsObj.Utc = app.timestamp.GetTimeFromUnixTime(unixval)
        tsObj.Unix = unixval
        b, err := json.Marshal(tsObj)
        if err != nil {
            http.Error(w, "Internal Server Error", 500)
            return
        }
        w.Write(b)
    }

    // check if time can be generated from datestring
    utcval, unixval, err := GetTimeFromDateString(queryString)
    if err == nil {
        tsObj := &timestampObj{}
        tsObj.Utc = utcval
        tsObj.Unix = unixval
        b, err := json.Marshal(tsObj)
        if err != nil {
            http.Error(w, "Internal Server Error", 500)
            return
        }
        w.Write(b)
    }

    // error situation
    errObj := &errorObj{
        Error: "Date-time in query string is not correct",
    }
    b, _ := json.Marshal(errObj)
    w.Write(b)
}
