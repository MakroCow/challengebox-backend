package service

import (
	"ChallengeBox/api"
	"ChallengeBox/databaseService"
	"ChallengeBox/models"
	"database/sql"
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
	challenges := databaseService.Challenge()
	api.SendJson(w, challenges)
}
