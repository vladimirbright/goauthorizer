package web


import (
	"fmt"
	"net/http"
	"text/template"
	"log"
)


type Context struct {
	Title string
	Header string
}

var (
	templates map[string]*template.Template
)


func init() {

	templates = make(map[string]*template.Template)

	templates["index"] = template.Must(
		template.ParseFiles("src/goauthorizer/templates/_base.html",
							"src/goauthorizer/templates/index.html"))

	log.Printf("Loaded templates: %@", templates)
}


func Render(response http.ResponseWriter, template_name string, context Context) {
	log.Printf("Template to render: %s ", template_name)
	log.Printf("Context to render: %s ", context)
	err := templates[template_name].Execute(response, context)
	if err != nil {
		log.Printf("Error on render %s", template_name)
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}


func MainPage(w http.ResponseWriter, r *http.Request) {
	log.Println("MainPage")
	var p Context
	p = Context{"Прювет!", "МюдвеД"}
	Render(w, "index", p)
}


func LoginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func RegisterPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func NotFoundPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
