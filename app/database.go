package app

import (
	"database/sql"
	"time"

	"github.com/itsahyarr/learn-go-restful-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:thinkpad@tcp(172.20.0.3)/belajar_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
