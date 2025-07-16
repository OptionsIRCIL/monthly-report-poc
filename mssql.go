//go:build ms || mssql || microsoft

package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/denisenkom/go-mssqldb"
)

func SQLOpen() (*sql.DB, error) {
    err := godotenv.Load()
    if err != nil {
	    log.Fatal("Error loading .env file")
    }

    sql_user     := os.Getenv("SQL_USERNAME")
    sql_pass     := os.Getenv("SQL_PASSWORD")
    sql_location := os.Getenv("SQL_LOCATION")
    sql_db       := os.Getenv("SQL_DATABASE")

    // Microsoft SQL Server
    connection := fmt.Sprintf("odbc:server=%s;user id=%s;password=%s;database=%s;TrustServerCertificate=true;encrypt=disable;ApplicationIntent=readonly", sql_location, sql_user, sql_pass, sql_db)

    db, err := sql.Open("mssql", connection)
    return db, err
}
