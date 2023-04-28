package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sdivyansh59/snippetbox/pkg/models/mysql"
)

// Define an application struct to hold application wide
//dependencies for web app.
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetModel
}



func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Define a new command-line flag for the MySQL DNS string.
	dns := flag.String("dns","web:pass@/snippetbox?parseTime=true", "MySQL database") 
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// TO keep the main() tidy I've put the code for creating a connection pool into the seperate openDB()
	// func below. We pass openDB() the DNS 
	// from the command-line flag.
	db, err := openDB(*dns)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection  pool is closed 
	// before the main() function exits.
	defer db.Close()


	// Initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}


	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(), // call the new app.routes() method
	}
	// writing message using 2 new logger
	infoLog.Printf("Starting server on %s", *addr)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}


// The openDB() func wraps sql.Open() and return a sql.DB connection pool for
// a given DNS.
func openDB(dns string ) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
