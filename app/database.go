package app

import (
	"database/sql"
	"time"
)

// "github.com/muhsufyan/dasar-golang2/helper"
func NewDB() *sql.DB {
	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/categoryapi")
	// helper.PanicIfError(err)

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// db.SetConnMaxIdleTime(10 * time.Minute)

	// return db
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/categoryapi")
	if err != nil {
		panic(err)
	}
	// close koneksi dilakukan saat test, karena ini menggunakan db pooling
	// set db pool
	db.SetMaxIdleConns(10)                  //min koneksi
	db.SetMaxOpenConns(100)                 //max koneksi
	db.SetConnMaxIdleTime(5 * time.Minute)  //dlm 5 menit jika koneksi tdk digunakan maka koneksi akan diputus (closed)
	db.SetConnMaxLifetime(60 * time.Minute) //stlh 60 menit koneksi akan renew
	return db
}
