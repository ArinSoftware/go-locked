package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to my super site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprintf(w, "To get in touch email to <a href=\"maitto:test@test.com\">test@test.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "<h1>We can't find yhe page you are looking for</h1>")
	}
	//fmt.Fprintf(w, "<h1>Welcome to my super site</h1>")
	//fmt.Fprintf(w, "To get in touch email to <a heref=\"maitto:test@test.com\">test@test.com</a>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":4000", nil)
}
