package main

import (
    "net/http"
    
    "github.com/julienschmidt/httprouter"
)

func (app *application) routes() *http.ServeMux {
    /*
    mux := http.NewServeMux()
    mux.HandleFunc("/api/timestamp", app.getTimestamp)

    fileServer := http.FileServer(http.Dir("./ui/static"))
    mux.Handle("/", fileServer)
    */
    mux := httprouter.New()
    mux.GET("/api/timestamp", app.getTimestamp)
    mux.GET("/api/timestamp/:timeformat", app.getTimestampFromQueryString)

    return mux
}
