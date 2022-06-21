package repository

import (
	"github.com/gin-gonic/gin"
)

func (repository *Repository) CreateForumTag(c *gin.Context, forumID int64, tagID int64) (bool, error) {
	query := "INSERT INTO ForumTags (forum_id, tag_id) VALUES (?, ?)"
	_, err := repository.db.Exec(query, forumID, tagID)

	if err != nil {
		return false, err
	}

	return true, nil
}
