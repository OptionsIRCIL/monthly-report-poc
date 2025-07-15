package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    // Connect to MySQL/MariaDB Database
    db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/ILCWorkGroups")
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Successful connection")
    }
    defer db.Close()

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
