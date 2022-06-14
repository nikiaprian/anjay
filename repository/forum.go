package repository

import (
	"kel15/models"


	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllForum(c *gin.Context) ([]models.Forum, error) {
	var forums []models.Forum

	query := "SELECT * FROM Forums"

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var forum models.Forum
		err := rows.Scan(&forum.ID, &forum.Title, &forum.Tag, &forum.Content, &forum.CreatedAt, &forum.UpdatedAt)
		if err != nil {
			return nil, err
		}
		forums = append(forums, forum)
	}

	return forums, nil
}