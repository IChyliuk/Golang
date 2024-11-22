package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homeeeeee!")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("templates/login.html")
	template.Execute(w, nil)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login/", loginPage)
	http.ListenAndServe("localhost:8080", nil)

}

type Login struct {
	name string
	pass string
}
