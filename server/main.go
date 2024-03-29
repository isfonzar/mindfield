package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type HomePage struct {
}

func main() {
	// Serve files under the ./static/css directory under the /css prefix
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./public/assets/css"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Parse the HTML template
	tmpl, err := template.ParseFiles("html/pages/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an instance of `HomePage`
	data := HomePage{}

	// Execute the template, inserting the name into the placeholder
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// Parse the HTML template
	tmpl, err := template.ParseFiles("html/pages/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an instance of `HomePage`
	data := HomePage{}

	// Execute the template, inserting the name into the placeholder
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
