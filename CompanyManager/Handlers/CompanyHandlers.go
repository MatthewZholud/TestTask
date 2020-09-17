package Handlers

import (
	"encoding/json"
	"github.com/MatthewZholud/TestTask/CompanyManager/DbService"
	"github.com/MatthewZholud/TestTask/CompanyManager/Entities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


func PostCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := Entities.Company{}
	err := parseJsonToStruct(w, r, &company)
	if err != nil {
		return
	}
	err = DbService.Conn.PostCompany(&company)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid Company")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(company)
}

func PutCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := Entities.Company{}

	err := parseJsonToStruct(w, r, &company)
	if err != nil {
		return
	}
	err = DbService.Conn.PutCompany(&company)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Company not found")
		return
	}
	json.NewEncoder(w).Encode(company)
}


func GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := Entities.Company{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
		return
	}
	company.ID = int64(id)

	company, err = DbService.Conn.GetCompany(company.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Company not found")
		return
	}
	json.NewEncoder(w).Encode(company)
	w.WriteHeader(http.StatusOK)
}

func PostCompanyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := Entities.Company{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
		return
	}
	company.ID = int64(id)

	company.Name = r.Form.Get("name")
	company.LegalForm = r.Form.Get("status")


	err = DbService.Conn.PostCompanyById(&company)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
		return
	}
	json.NewEncoder(w).Encode(company)
}

func DeleteCompanyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := Entities.Company{}
	if IsNumericAndPositive(mux.Vars(r)["id"]) != true {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
		return
	}
	company.ID = int64(id)
	err = DbService.Conn.DeleteCompany(company.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Company not found")
		return
	}
}

func GetEmployeeByCompanyID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	company := Entities.Company{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
		return
	}
	company.ID = int64(id)
	employees, err := DbService.Conn.GetEmployeesByCompanyId(company.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Company not found")
		return
	}
	json.NewEncoder(w).Encode(employees)
	w.WriteHeader(http.StatusOK)
}