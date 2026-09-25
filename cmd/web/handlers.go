package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// w.Write([]byte("Hello from SnippetBox")) // require raw bytes, thus the conversion from string to bytes

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	/* Extract the value of the id parameter from the query string and try to
	* convert it to an integer using the strconv.Atoi() function. If it can't
	* be converted to an integer, or the value is less than 1, we return a 404 page
	* not found response.
	 */
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		/*
		* .Set() replaces if there is an existing header
		* .Add() appends to existing header
		* .Del() delets all values of the given header
		* doesn't remove system generated headers, instead they can be suppressed using w.Header()['Date'] = nil
		* .Get() retrieves the first value
		* .Values() gets a slice of all
		 */
		w.Header().Set("Allow", "POST")

		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		/*
		* This above code snippet can be replaced by the line written below
		* In terms of functionality it is exactly the same
		* which is a shortcut that calls w.WriteHeader() and w.Write() underneath
		* The main difference is now http.ResponseWriter is being passed to another function
		* Which sends the response to the user for us
		 */
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
