package usecase

import (
	"errors"
	"kel15/models"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) GetAllForum(c *gin.Context) ([]models.Forum, error) {
	forums, err := usecase.repository.GetAllForum(c)
	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (usecase *Usecase) CreateForum(c *gin.Context) (*models.Forum, error) {
	user, _ := c.Get("user")
	userData := user.(*models.User)

	if user == nil {
		return nil, errors.New("user not found")
	}

	var forum models.ForumRequest

	err := c.BindJSON(&forum)

	if err != nil {
		return nil, err
	}

	forumResponse, err := usecase.repository.CreateForum(c, forum.Title, forum.Content, userData.ID)

	if err != nil {
		return nil, err
	}

	for _, tag := range forum.Tags {
		Tag, _ := usecase.repository.CreateTag(c, tag)
		if Tag != nil {
			usecase.repository.CreateForumTag(c, int64(forumResponse.ID), int64(Tag.ID))
		}
	}

	return forumResponse, nil

}
