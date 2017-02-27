package main

import (
	"fmt"
	"html/template"
	"iching/bagua"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

type NewPage struct {
	Title   string
	Body    []string
	Chapter int
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("chapters/"+filename, p.Body, 0600)
}

func loadChapter(title string) (*Page, error) {
	filename := "chapters/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func chapHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadChapter(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadChapter(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/chap/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "index.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view|chap)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	hexStr, hexSig := bagua.BaguaGenerator()
	p := &NewPage{
		Title:   "Welcome Page",
		Body:    hexStr,
		Chapter: hexSig,
	}
	_ = templates.ExecuteTemplate(w, "index.html", p)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/chap/", makeHandler(chapHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Println("Listening on 8080...")
	http.ListenAndServe(":8080", nil)

}
