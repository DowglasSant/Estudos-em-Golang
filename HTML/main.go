package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("*.html"))

	u := usuario{"Dowglas", "dowglas.santana98@gmail.com"}

	templates.ExecuteTemplate(w, "home.html", u)
}

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {

	http.HandleFunc("/home", home)

	fmt.Println("Server listening on port 5000!")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
