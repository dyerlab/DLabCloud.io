package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/navbar.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}

}

/*************************************************

		Snippet Object Routing

***************************************************/

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return

	}

	snip, err := app.snippets.Get(id)
	if err != nil {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with id=%d\n title=%s\n contents=%s\n", id, snip.Title, snip.Contents)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow","POST")
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"

	uid, _ := app.snippets.Insert(title,content)

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", uid), http.StatusSeeOther)
}





/*************************************************

		Manuscript Object Routing

***************************************************/

func (app *application) showManuscript(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Displaying manuscript %d", id)
}

func (app *application) createManuscript(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creaing new manuscript..."))
}

// AllManuscripts lists all manuscripts in the database.
// func allManuscripts(w http.ResponseWriter, r *http.Request) {
// 	db, err := gorm.Open(dbName, dbConnect)
// 	if err != nil {
// 		panic("Could not open database")
// 	}
// 	defer db.Close()

// 	var manuscripts []Manuscript
// 	db.Find(&manuscripts)

// 	data := ManuscriptData{
// 		Title:       "Manuscripts",
// 		Manuscripts: manuscripts,
// 	}

// 	fmt.Printf("Serving up %d manuscripts", len(manuscripts))

// 	tmpl := template.Must(template.ParseFiles("manuscripts.html"))
// 	tmpl.Execute(w, data)
// 	//json.NewEncoder(w).Encode(manuscripts)
// }
