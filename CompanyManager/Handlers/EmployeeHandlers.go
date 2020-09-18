package Handlers

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/MatthewZholud/TestTask/CompanyManager/DbService"
	"github.com/MatthewZholud/TestTask/CompanyManager/Entities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Entities.Employee{}

	err := parseJsonToStruct(w, r, &employee)
	if err != nil {
		return
	}
	err = DbService.Conn.PostEmployee(&employee)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func PutEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Entities.Employee{}

	err := parseJsonToStruct(w, r, &employee)
	if err != nil {
		return
	}
	err = DbService.Conn.PutEmployee(&employee)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Entities.Employee{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
		return
	}
	employee.ID = int64(id)
	employee, err = DbService.Conn.GetEmployee(employee.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}
	json.NewEncoder(w).Encode(employee)
	w.WriteHeader(http.StatusOK)
}

func PostEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Entities.Employee{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
		return
	}
	employee.ID = int64(id)

	id, err = strconv.Atoi(r.Form.Get("CompanyId"))
	if err != nil {
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
		return
	}
	employee.CompanyID = int64(id)
	employee.Name = r.Form.Get("name")
	employee.SecondName = r.Form.Get("secondName")
	employee.Surname = r.Form.Get("surname")
	employee.HireDate = r.Form.Get("hireDate")
	employee.Position = r.Form.Get("position")

	err = DbService.Conn.PostEmployeeById(&employee)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Entities.Employee{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
		return
	}
	employee.ID = int64(id)
	err = DbService.Conn.DeleteEmployee(employee.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusNotFound, "Employee not found")
		return
	}
}
