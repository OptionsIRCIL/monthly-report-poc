//go:build my || mysql || mariadb

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

    // MySQL/MariaDB
    connection := fmt.Sprintf("%s:%s@%s/%s", sql_user, sql_pass, sql_location, sql_db)
    db, err := sql.Open("mysql", connection)
    return db, err
}
