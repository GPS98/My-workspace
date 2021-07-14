package main

import (
	"net/http"

	"github.com/russross/blackfriday"
	//go get "github.com/russross/blackfriday"
) //go get

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
