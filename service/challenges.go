package service

import (
	"ChallengeBox/databaseService"
	"ChallengeBox/models"
	"database/sql"
	"fmt"
	"net/http"
)

var DB *sql.DB

func AddChallenge(c models.Challenge) int {
	return databaseService.SaveChallenge(c)
}

func getTags() *[]models.Tag {
	return nil
}

func Challenges(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, databaseService.Challenge())
}