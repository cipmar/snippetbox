package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"snippetbox.cipmar.net/internal/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	//
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	//
	// err = ts.ExecuteTemplate(w, "base", nil)
	//
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	//
	// w.Write([]byte("Hello from snippetbox"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Display a form for creating a snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	title := "Luceafarul"
	content := "A fost odata ca-n povesti\nA fost ca niciodata\nDin rude mari, impaatesti\nO prea frumoasa fata."
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippets/view/%d", id), http.StatusSeeOther)
}
