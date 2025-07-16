//go:build !mssql && !mysql && !mariadb && !ms && !microsoft

package main

import (
    "database/sql"
    "errors"
    "log"
)

func SQLOpen() (*sql.DB, error) {
    log.Fatal("No database chosen. Recompile this program with `-tags {mssql,mysql,mariadb}`")
    return nil, errors.New("No database chosen.")
}

