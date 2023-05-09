package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", loginHandler)

	http.ListenAndServe(":8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	} else if r.Method == "POST" {
		// handle login form submission here
	}
}
