package Handlers

import (
	_ "encoding/json"
	"net/http"
)

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//groups, err := DbService.Conn.GetGroupsDb()
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Error: %v", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(groups)
}

func PutEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//group := Entities.Groups{}
	//
	//err := r.ParseForm()
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Error: %v", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//group.Title = r.Form.Get("group_title")
	//
	//group.ID, err = DbService.Conn.PostGroup(&group)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Error: %v", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(group)
}


func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//id := mux.Vars(r)["id"]
	//
	//err := DbService.Conn.DeleteGroup(id)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Error: %v", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusNoContent)
}

func PostEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//group := Entities.Groups{}
	//group.ID = (mux.Vars(r))["id"]
	//
	//group.Title = r.Form.Get("group_title")
	////group.Title = "Work"
	//
	//err := DbService.Conn.PutGroup(&group)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Error: %v", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(group)
}

func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//group := Entities.Groups{}
	//group.ID = (mux.Vars(r))["id"]
	//
	//group.Title = r.Form.Get("group_title")
	////group.Title = "Work"
	//
	//err := DbService.Conn.PutGroup(&group)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Error: %v", err))
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(group)
}