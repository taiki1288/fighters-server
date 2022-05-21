package db

import "database/sql"

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/fighters-server?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	err = db.Close()
	return db, err
}