package databaseService

import (
	"ChallengeBox/models"
	"fmt"
)


func RetrieveTag(challengeId int) *[]models.Tag {

	sqlStatement := `SELECT tag.id, tag.name FROM tag, tag_challenge where tag.id = tag_challenge.tag_id and challenge_id = $1`

	var tags []models.Tag
	tags = make([]models.Tag, 0)

	rows, err := DB.Query(sqlStatement, challengeId)
	if err != nil {
		fmt.Print(err)
		return nil //todo better error handling
	}
	defer rows.Close()

	for rows.Next() {
		var tagId int
		var tagName string
		err := rows.Scan(&tagId, &tagName)
		if err != nil {
			fmt.Println(err)
			return nil //todo better error handling
		}

		tags = append(tags, *models.NewTag(tagId, tagName))
	}

	return &tags
}