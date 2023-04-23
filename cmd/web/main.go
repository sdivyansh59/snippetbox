package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold application wide
//dependencies for web app.
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}



func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}


	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}
	// writing message using 2 new logger
	infoLog.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
