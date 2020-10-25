package service

import (
	"ChallengeBox/api"
	"ChallengeBox/databaseService"
	"net/http"
	"strconv"
	"strings"
)

func TagById(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	splittedPath := strings.Split(path, "/")
	challengeId := splittedPath[len(splittedPath)-1]
	challengeIdInt, err := strconv.Atoi(challengeId)
	if err != nil {
		return //todo better error handling
	}
	tags := databaseService.RetrieveTagsByChallengeId(challengeIdInt)

	api.SendJson(w, tags)
}

func Tag(w http.ResponseWriter, r *http.Request) {
	tags := databaseService.RetrieveTags()
	api.SendJson(w, tags)
}
