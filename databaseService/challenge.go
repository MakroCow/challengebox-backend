package databaseService

import (
	"ChallengeBox/models"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func SaveChallenge(challenge models.Challenge) int {
	sqlStatementChallenge := `INSERT INTO challenge (title, description, sport_value, intelligence_value) VALUES ($1, $2, $3, $4) RETURNING id`

	id := 0
	err := DB.QueryRow(sqlStatementChallenge, "Title", "jon@calhoun.io", 12, 15).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func Challenge() *[]models.Challenge {
	sqlStatement := `SELECT id, title, description, sport_value, intelligence_value, culinary_value, selfcare_value FROM challenge;`
fmt.Println(DB)
	rows, err := DB.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
		// todo errorhandling
	}

	var challenges = make([]models.Challenge, 0)
	for rows.Next() {
		var challenge models.Challenge

		err := rows.Scan(&challenge.Id, &challenge.Title, &challenge.Description, &challenge.SportValue,
			&challenge.IntelligenceValue, &challenge.CulinaryValue, &challenge.SelfcareValue)
		if err != nil {
			fmt.Println(err)
			return nil //todo better error handling
		}

		challenges = append(challenges, challenge)
	}

	return &challenges
}
