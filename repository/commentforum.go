package repository

import (
	"kel15/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) CreateCommentForum(c *gin.Context, comment string, forum_id, user_id int) (*models.CommentForum, error) {

	query := `INSERT INTO CommentForum (comment, forum_id, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?);`

	result, err := repository.db.Exec(query, comment, forum_id, user_id, time.Now(), time.Now())
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

	return &models.CommentForum{
		ID:        int(id),
		Comment:   comment,
		User: 	*user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *Repository) GetAllCommentByForumID(c *gin.Context, id int) ([]models.CommentForum, error) {

	query := `SELECT Comments.id, Comments.comment, Comments.created_at, Comments.updated_at,
			 Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at
		 	 FROM CommentForum as Comments 
			JOIN Users ON Comments.user_id = Users.id
			WHERE Comments.forum_id = ?;`

	rows, err := repository.db.Query(query, c.Param("id"))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comment models.CommentForum
	var comments []models.CommentForum
	var User models.User

	for rows.Next() {

		err := rows.Scan(&comment.ID, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt,
			&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comment.User = User
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *Repository) DeleteCommentForum(c *gin.Context, id int) error {

	query := `DELETE FROM CommentForum WHERE id = ?;`

	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}