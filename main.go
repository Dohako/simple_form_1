package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	data := struct {
		First string
		Last  string
	}{
		First: first_name,
		Last:  last_name,
	}

	tpl.ExecuteTemplate(w, "processor.html", data)
}
