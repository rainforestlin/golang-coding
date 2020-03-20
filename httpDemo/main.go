package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
type IndexHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func (i *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "url is %s", r.URL.Path[1:])
}
func main() {
	handler := MyHandler{}
	index := IndexHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/index", &handler)
	http.Handle("/go_path", &index)

	server.ListenAndServe()
}
