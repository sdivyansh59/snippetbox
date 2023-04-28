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
		return 
	}

	err = ts.Execute(w,nil)
	if err != nil {
		app.serveError(w,err) // use the serverError() helper
		return
	}

}

 
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying  a specific snippet with ID %d ...", id)
}



func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}

	// crete some variable holding dummy data
	title := "0 small"
	content := "0 small/nclimb mount fuji,\nBut slowely\n\n- Kobayanshi"
	expires := "7"

	// pass data to SnippetModel.Insert()
	id, err := app.snippets.Insert(title,content,expires)
	if err != nil {
		app.serveError(w, err)
		return 
	}

	// REdirect the user to the relevent page for the snippet.
	http.Redirect(w,r,fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}