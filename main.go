package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    // Connect to MySQL/MariaDB Database
    db, err := sql.Open("mysql", "root:example@tcp(localhost:3306)/test")
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Successful connection?")
    }
    defer db.Close()

}
