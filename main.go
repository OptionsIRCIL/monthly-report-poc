package main

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

func queryDb() [][]string {
	var tableContent [][]string

	// Open Database
	db, err := SQLOpen()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// List employees
	employees, err := db.Query("SELECT FirstName, LastName FROM DBA_t_Employees")
	if err != nil {
		panic(err.Error())
	}
	defer employees.Close()

	// Move data into the required format
	for employees.Next() {
		var firstName sql.NullString
		var lastName sql.NullString

		err := employees.Scan(&firstName, &lastName)
		if err != nil {
			panic(err.Error())
		}

		row := []string{}
		row = append(row, firstName.String)
		row = append(row, lastName.String)

		tableContent = append(tableContent, row)
	}

	return tableContent
}

func main() {
	createReport()
}
