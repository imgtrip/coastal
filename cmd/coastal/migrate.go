package main

import (
    "coastal/internal/env"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    golangMigrate "github.com/golang-migrate/migrate"
    "github.com/golang-migrate/migrate/database/mysql"
    _ "github.com/golang-migrate/migrate/source/file"
)

func migrate() {
    conf := env.GetDBConfig()

    connect := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?multiStatements=true",
        conf.DBUser,
        conf.DBPassword,
        conf.DBHost,
        conf.DBPort,
        conf.DBName,
    )

    db, err := sql.Open("mysql", connect)

    if err != nil {
        panic(err)
    }

    driver, _ := mysql.WithInstance(db, &mysql.Config{})

    m, err := golangMigrate.NewWithDatabaseInstance(
        "file://migrations",
        "mysql",
        driver,
    )
    if err != nil {
        panic(err)
    }

    defer m.Close()

    err = m.Up()
    if err != nil && err != golangMigrate.ErrNoChange {
        version, _, _ := m.Version()
        fmt.Printf("migrate error: %v", err.Error())
        panic(fmt.Sprintf("migrate errored version : %v", version))
    }
}
