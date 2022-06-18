package repository

import (
	"kel15/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) CreateCommentBlog(c *gin.Context, comment string, blog_id, user_id int) (*models.CommentBlog, error) {
	
	query := `INSERT INTO CommentBlog (comment, blog_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?);`

	result, err := repository.db.Exec(query, comment, blog_id, user_id, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserById(c, int64(user_id))
	if err != nil {
		return nil, err
	}

	return &models.CommentBlog{
		ID:        int(id),
		Comment:   comment,
		User: 	*user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}