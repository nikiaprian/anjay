package repository

import (
	"kel15/models"
	"time"

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

func (repository *Repository) CreateForum(c *gin.Context, req models.ForumRequest) (*models.ForumResponse, error) {
	title := req.Title
	tag := req.Tag
	content := req.Content

	query := "INSERT INTO Forums (title, tag, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	row1, err := repository.db.Exec(query, title, tag, content, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	id, err := row1.LastInsertId()
	if err != nil {
		return nil, err
	}

	query = "SELECT * FROM Forums WHERE id = ?"
	row2 := repository.db.QueryRow(query, id)
	if err != nil {
		return nil, err
	}

	var forum models.Forum
	err = row2.Scan(&forum.ID, &forum.Title, &forum.Tag, &forum.Content, &forum.CreatedAt, &forum.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &models.ForumResponse{
		Forum:   &forum,
		Message: "Forum Berhasil Ditambahkan",
	}, nil
}