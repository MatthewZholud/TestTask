package DbService

import "github.com/MatthewZholud/TestTask/CompanyManager/Entities"

type CompanyDB interface {
	PostCompany(company *Entities.Company) error
	PutCompany(company *Entities.Company) error
	GetCompany(id int64) (Entities.Company, error)
	PostCompanyById(company *Entities.Company) error
	DeleteCompany(id int64) error
	GetEmployeesByCompanyId(id int64) ([]Entities.Employee, error)
}

func (conn *DbStruct) PostCompany(company *Entities.Company) error {
	_, err := conn.db.Exec("INSERT INTO company(name, legal_form) VALUES ($1, $2)", company.Name, company.LegalForm)
	return err
}

func (conn *DbStruct) PutCompany(company *Entities.Company) error {
	_, err := conn.db.Exec("UPDATE company set name = $1, legal_form = $2 where company_id = $3;", company.Name, company.LegalForm, company.ID)
	return err
}

func (conn *DbStruct) GetCompany(id int64) (Entities.Company, error) {
	company := Entities.Company{}
	rows, err := conn.db.Query("SELECT * from company WHERE company_id = $1", id)
	if err != nil {
		return company, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&company); err != nil {
			return company, err
		}
	}
	return company, nil
}

func (conn *DbStruct) PostCompanyById(company *Entities.Company) error {
	_, err := conn.db.Exec("UPDATE company set name = $1, legal_form = $2 where company_id = $3;", company.Name, company.LegalForm, company.ID)
	return err
}

func (conn *DbStruct) DeleteCompany(id int64) error {
	_, err := conn.db.Exec("DELETE FROM company WHERE company_id = $1", id)
	return err
}

func (conn *DbStruct) GetEmployeesByCompanyId(id int64) ([]Entities.Employee, error) {
	rows, err := conn.db.Query("SELECT * from employee WHERE company_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []Entities.Employee{}

	for rows.Next() {
		employee := Entities.Employee{}

		if err := rows.Scan(&employee); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, err
}
