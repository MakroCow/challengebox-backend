package service

import (
	"ChallengeBox/databaseService"
	"net/http"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	err := databaseService.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	// todo weiterleitung einrichten
}
