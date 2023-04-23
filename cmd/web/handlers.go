package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)


func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // use the notFound() helper
		return
	}

	files := []string {
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serveError(w,err) // use the serverError() helper.
		// app.errorLog.Println(err.Error())
		// http.Error(w,"Internal Server Error ", http.StatusInternalServerError)
		return 
	}

	err = ts.Execute(w,nil)
	if err != nil {
		app.serveError(w,err) // use the serverError() helper
		// app.errorLog.Println(err.Error())
		// http.Error(w,"Internal Server Error", 500)
		return
	}

}


 
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		// http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying  a specific snippet with ID %d ...", id)
}



func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w,http.StatusMethodNotAllowed)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Crete a new snippet"))
}