package main

import (
	"net/http"
	"io"
)

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Docker is Go!</h1>")
}

