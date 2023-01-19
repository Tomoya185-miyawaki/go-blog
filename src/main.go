package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/top", handler2)
	http.ListenAndServe(":8000", nil)
}

func handler(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "test")
}

func handler2(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "a")
}
