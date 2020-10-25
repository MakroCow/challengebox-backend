package main

import (
	"ChallengeBox/databaseService"
	"ChallengeBox/environment"
	"database/sql"
	"fmt"
)

var db *sql.DB

func ConnectDatabase() {
	config := environment.Parse()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName)
	var err error
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")

	distributeDBObject()
}

func distributeDBObject() {
	//service.DB = db
	databaseService.DB = db
}
