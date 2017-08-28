package main

import (
	"net/http"
	"io"
	"github.com/34South/envr"
	"os"
)

func init() {
	envr.New("now-test", []string{
		"TEST_VAR",
	}).Auto()
}


func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Docker is Go!</h1><h2>Hi, " + os.Getenv("TEST_VAR") + "</h2>")
}

