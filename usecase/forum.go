package usecase

import (
	"errors"
	"kel15/models"
	"strconv"

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

func (usecase *Usecase) UpdateForum(c *gin.Context) (*models.Forum, error) {
	user := c.MustGet("user").(*models.User)

	if user == nil {
		return nil, errors.New("user not found")
	}

	id := c.Param("id")
	Convid, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	var forum models.ForumRequest

	err = c.BindJSON(&forum)

	if err != nil {
		return nil, err
	}

	forumResponse, err := usecase.repository.UpdateForum(c, Convid, forum.Title, forum.Content)

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

func (usecase *Usecase) DeleteForum(c *gin.Context) (*models.Forum, error) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		return nil, errors.New("user not found")
	}

	id := c.Param("id")
	Convid, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	err = usecase.repository.DeleteForum(c, Convid)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
