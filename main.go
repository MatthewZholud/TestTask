package main

import (
	"fmt"
	"github.com/MatthewZholud/TestTask/Handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

func main() {
	fmt.Println("Starting server...")

	r := newRouter()

	PORT := ":8000"
	fmt.Println("Serving on port: ", PORT)

	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}