//go:build !ms && !mssql && !microsoft && !my && !mysql && !mariadb 

package main

import (
    "database/sql"
    "errors"
    "log"
)

func SQLOpen() (*sql.DB, error) {
    log.Fatal("No database chosen. Recompile this program with `-tags {ms,mssql,microsoft,my,mysql,mariadb}`")
    return nil, errors.New("No database chosen.")
}

