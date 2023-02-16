package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

const Port_Number = "8090"
const Host_Address = "127.0.0.1"

// HomePageHandler is the home page handler
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "home.html")
}

// AboutPageHandler is the about page handler
func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "about.html")
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + templateName)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutPageHandler)

	baseUrl := fmt.Sprintf("%v:%v", Host_Address, Port_Number)
	fmt.Printf("Trying to run the server on: http://%v\n", baseUrl)
	err := http.ListenAndServe(baseUrl, nil)
	if err != nil {
		log.Fatal("While trying to up the serve an error ocurred", err)
		os.Exit(1)
	}
}
