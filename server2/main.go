package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("started another server")

	route := mux.NewRouter()
	route.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello ")
		_, err := fmt.Fprintf(writer, "Hello from server 2")
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":1008", route)
	if err != nil {
		return
	}
}
