package main

import (
	"fmt"
	"log"
	"mod/CompanyManager/Handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	//staticFileDirectory := http.Dir("./assets/")
	//staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	//r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/employee", Handlers.PostEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee", Handlers.PutEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", Handlers.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", Handlers.PostEmployeeByID).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", Handlers.DeleteEmployeeByID).Methods(http.MethodDelete)

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