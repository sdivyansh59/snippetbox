package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serveError helper writes an error message and stack trace to the errorLog then sends a
// generic 500 Internal Server response to the user
func (app *application) serveError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description to the user 
// We'll use this later in the book to send response like 400 "BadREquest"
// when there is a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implements a notFound helper. This is simply a conveince wrapper arond 
// clientError which a 404 Not Found response to the user
func (app *application) notFound( w http.ResponseWriter) {
	app.clientError(w,http.StatusNotFound)
}