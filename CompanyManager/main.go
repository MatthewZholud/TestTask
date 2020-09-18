package main

import (
	"fmt"
	"github.com/MatthewZholud/TestTask/CompanyManager/DbService"
	"log"
	"net/http"

	"github.com/MatthewZholud/TestTask/CompanyManager/Handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/employee", Handlers.PostEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee", Handlers.PutEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", Handlers.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", Handlers.PostEmployeeByID).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", Handlers.DeleteEmployeeByID).Methods(http.MethodDelete)

	r.HandleFunc("/company/", Handlers.PostCompany).Methods(http.MethodPost)
	r.HandleFunc("/company/", Handlers.PutCompany).Methods(http.MethodPut)
	r.HandleFunc("/company/{companyId}", Handlers.GetCompany).Methods(http.MethodGet)
	r.HandleFunc("/company/{companyId}", Handlers.PostCompanyByID).Methods(http.MethodPost)
	r.HandleFunc("/company/{companyId}", Handlers.DeleteCompanyByID).Methods(http.MethodDelete)
	r.HandleFunc("/company/{companyId}/employees", Handlers.GetEmployeeByCompanyID).Methods(http.MethodGet)

	return r
}

func main() {
	fmt.Println("Starting server...")
	DbService.Db_Conn()

	r := newRouter()

	PORT := ":8000"
	fmt.Println("Serving on port: ", PORT)

	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}