package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const SRV_PORT = ":8080"
const SRV_HOST = ""
const DATA_DIR = "data/"

type Page struct {
	Title    string
	Body     []byte
	HTMLBody template.HTML
}

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9/\\.\\-]*$")

func (p *Page) save() error {
	filename := DATA_DIR + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := DATA_DIR + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body, HTMLBody: template.HTML(blackfriday.MarkdownCommon(body))}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
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
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func staticHandler(w http.ResponseWriter, r *http.Request, url string) {
	http.ServeFile(w, r, "static/"+url)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parts []string
		var title string

		parts = strings.SplitN(r.URL.Path[1:], "/", 2)

		if len(parts) == 1 {
			title = ""
		} else {
			title = parts[1]
		}

		// route := parts[0]
		if !titleValidator.MatchString(title) {
			fmt.Println("not found: " + title)
			http.NotFound(w, r)
			return
		}
		fn(w, r, title)
	}
}

func index(w http.ResponseWriter, r *http.Request, title string) {
	http.Redirect(w, r, "/view/Main", http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/static/", makeHandler(staticHandler))
	http.HandleFunc("/", makeHandler(index))
	fmt.Println("Setting up server on " + SRV_HOST + SRV_PORT)
	err := http.ListenAndServe(SRV_HOST+SRV_PORT, nil)
	if err != nil {
		fmt.Println("An error ocurred: " + err.Error())
	}
}
