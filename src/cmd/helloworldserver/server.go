package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		name := request.FormValue("name")
		fmt.Fprintf(writer, "Hello World %s\n", name)
	})
	http.ListenAndServe(":8000", nil)
}
