package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path))
	})

	http.HandleFunc("/hi", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hi!!!")
	})

	http.ListenAndServe(":8000", nil)
}
