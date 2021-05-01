package main

import (
    "flag"
    "log"
    "net/http"
    "os"

    "github.com/castevet6/timestamp-ms/pkg/timestamp"
)

// dependency wrapper for application
type application struct {
    timestamp   timestamp.Timestamp
    infoLog     *log.Logger
    errorLog    *log.Logger
}

func main() {
    // define args flag for HTTP address, and parse flags
    addr := flag.String("addr", ":4000", "HTTP Network Address")
    flag.Parse()

    // create logger instances
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // crate instance for timestamp package
    ts := timestamp.Timestamp{}
    
    // initialize app
    app := &application{
        errorLog:   errorLog,
        infoLog:    infoLog,
        timestamp:  ts,
    }

    // initialize server struct
    srv := &http.Server{
        Addr:       *addr,
        ErrorLog:   errorLog,
        Handler:    app.routes(),
    }

    infoLog.Printf("Starting server on %s\n", *addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}
