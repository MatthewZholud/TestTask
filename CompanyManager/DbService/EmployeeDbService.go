package DbService

import (
	"github.com/MatthewZholud/TestTask/CompanyManager/Entities"
)


type EmployeeDB interface {
	PostEmployee(employee *Entities.Employee) error
	PutEmployee(employee *Entities.Employee) error
	GetEmployee(id int64) (Entities.Employee, error)
	PostEmployeeById(employee *Entities.Employee) error
	DeleteEmployee(id int64) error
}

func (conn *DbStruct) GetEmployee(id int64) (Entities.Employee, error) {
	employee := Entities.Employee{}
	rows, err := conn.db.Query("SELECT * from employee WHERE employee_id = $1", id)
	if err != nil {
		return employee , err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&employee); err != nil {
			return employee, err
		}
	}
	return employee, nil
}

func (conn *DbStruct) PostEmployee(employee *Entities.Employee) error {
	_, err := conn.db.Exec("INSERT INTO employee(name, secondName, surname, photoUrl, hireDate, position, company_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		employee.Name, employee.SecondName, employee.Surname, employee.PhotoUrl, employee.HireDate, employee.Position, employee.CompanyID)
	return err
}

func (conn *DbStruct) PutEmployee(employee *Entities.Employee) error {
	_, err := conn.db.Exec("UPDATE employee set name = $1, secondName = $2, surname = $3, photoUrl = $4, hireDate = $5," +
		" position = $6, company_id = $7 where employee_id = $7;", employee.Name, employee.SecondName, employee.Surname,
		employee.PhotoUrl, employee.HireDate, employee.Position, employee.CompanyID, employee.ID)
	return err
}

func (conn *DbStruct) PostEmployeeById(employee *Entities.Employee) error {
	_, err := conn.db.Exec("UPDATE employee set name = $1, secondName = $2, surname = $3, photoUrl = $4, hireDate = $5," +
		" position = $6, company_id = $7 where employee_id = $7;", employee.Name, employee.SecondName, employee.Surname,
		employee.PhotoUrl, employee.HireDate, employee.Position, employee.CompanyID, employee.ID)
	return err
}

func (conn *DbStruct) DeleteEmployee(id int64) error {
	_, err := conn.db.Exec("DELETE FROM employee WHERE employee_id = $1", id)
	return err
}
