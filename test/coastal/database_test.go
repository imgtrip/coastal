package main

import (
	"coastal/internal/env"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func databaseUp() {
	databaseConnect().Up()
}

func databaseDown() {
	databaseConnect().Drop()
}

func databaseConnect() *migrate.Migrate {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		env.Process.DBUser,
		env.Process.DBPassword,
		env.Process.DBHost,
		env.Process.DBPort,
		env.Process.DBName,
	)

	db, err := sql.Open("mysql", connect)
	if err != nil {
		panic(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../cmd/coastal/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	return m
}
