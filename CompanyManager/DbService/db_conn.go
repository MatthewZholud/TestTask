package DbService

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/lib/pq"
)

type ServeDb interface {
	EmployeeDB
	CompanyDB
}

type DbStruct struct {
	db *sql.DB
}


func Db_Conn() {

	//PsqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	"postgresdb", "5432", "postgres", "mypassword", "time_tracker")
	_ = pq.Efatal

	config := DBConfig{}

	PsqlInfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
		config.GetUser(), config.GetPassword(), config.GetHost(), config.GetDBName(), config.GetPort())
	db, err := sql.Open("postgres", PsqlInfo)
	if err != nil {
		log.Panic(err)
	}


	err = db.Ping()
	if err != nil {
		panic(err)
	}

	InitDbService(&DbStruct{db: db})
}

var Conn ServeDb

func InitDbService(s ServeDb) {
	Conn = s
}
