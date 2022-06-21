package repository

import (
	"database/sql"
	"kel15/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetTagByID(c *gin.Context, id int64) (*models.ForumTag, error) {
	var tag models.ForumTag

	query := "SELECT * FROM Tags WHERE id = ?"
	row := repository.db.QueryRow(query, id)

	err := row.Scan(&tag.ID, &tag.Tag)
	if err != nil {
		return &tag, err
	}

	return &tag, nil
}

func (repository *Repository) GetTagByName(c *gin.Context, tag string) (*models.ForumTag, error) {
	var Tag models.ForumTag

	query := "SELECT * FROM Tags WHERE tag = ?"
	row := repository.db.QueryRow(query, tag)

	err := row.Scan(&Tag.ID, &Tag.Tag)
	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &Tag, nil
}

func (repository *Repository) CreateTag(c *gin.Context, tag string) (*models.ForumTag, error) {
	Tag, err := repository.GetTagByName(c, tag)
	if Tag != nil {
		return Tag, nil
	}

	query := "INSERT INTO Tags (tag, created_at, updated_at) VALUES (?, ?, ?)"
	result, err := repository.db.Exec(query, tag, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return repository.GetTagByID(c, id)
}
