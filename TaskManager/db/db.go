package db

import (
	"database/sql"
	"fmt"
	_"github.com/lib/pq"
)

var DB *sql.DB
func ConnectDb() error{
	ConnStr :="host=localhost port=5432 user=postgres password=atharva@2004 dbname=TaskManager sslmode=disable"

	db,error:=sql.Open("postgres",ConnStr)
	if error !=nil {
		return error
	}
	DB=db
	fmt.Println("connected to Postrgess")
	return nil
}