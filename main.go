package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/denisenkom/go-mssqldb"
)

func queryDb(db *sql.DB) {

	// List employees
    employees, err := db.Query("SELECT FirstName, LastName FROM DBA_t_Employees")
    if err != nil {
        panic(err.Error())
    }
    defer employees.Close()

    fmt.Println("Employees:")
    for employees.Next() {
        var firstName sql.NullString
        var lastName sql.NullString
        err := employees.Scan(&firstName, &lastName)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(firstName.String, lastName.String)
    }
}

func main() {

	db, err := SQLOpen()
	if err != nil {
		panic(err.Error())
	}
    defer db.Close()

    // queryDb(db)
    createReport()
}
