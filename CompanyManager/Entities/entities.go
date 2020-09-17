package Entities

import "time"

type Company struct {
		ID int64
		Name string
		LegalForm string
}

type Employee struct {
	Required   []string `json:"required"`
	ID int64 `json:"id""`
	Name string `json:"name"`
	SecondName string `json:"second_name"`
	Surname string `json:"surname"`
	HireDate time.Time `json:"hire_date"`
	Position string `json:"position"`
	//Position struct {
	//		Enum        []string `yaml:"enum"`hire_date
	//	} `yaml:"position"`
	CompanyID int64 `json:"company_id"`
}

type ArrayOfEmployees struct {
	Employee []Employee
}

type APIResponse struct {
		Code string `json:"code"`
		Type string `json:"type"`
		Message string `json:"message"`
}
