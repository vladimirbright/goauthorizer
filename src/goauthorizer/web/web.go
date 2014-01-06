package web


import (
	"fmt"
	"net/http"
	"text/template"
	"log"
)


var templates = template.Must(template.ParseGlob("src/goauthorizer/templates/*"))


func Render(response http.ResponseWriter, template_name string, context interface{}) {
	log.Printf("Шаблон для рендера: %s ", template_name)
	err := templates.ExecuteTemplate(response, template_name, context)
	if err != nil {
		log.Printf("Ошибка рендера %s", template_name)
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}


func MainPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Шаблоны: %@", templates)
	var p struct {}
	Render(w, "index.html", &p)
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
