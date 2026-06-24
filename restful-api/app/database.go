package app

import (
	"database/sql"
	"restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open(
		"mysql",
		"root:root123@tcp(localhost:3306)/golang_rest_pzn?parseTime=true",
	)
	helper.PanicIfError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
