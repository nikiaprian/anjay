package usecase

import (
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

func (usecase *Usecase) CreateForum(c *gin.Context) (*models.ForumResponse, error) {
	var payload models.ForumRequest
	err := c.BindJSON(&payload)
	if err != nil {
		return nil, err
	}

	forum, err := usecase.repository.CreateForum(c, payload)
	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (usecase *Usecase) UpdateForum(c *gin.Context) (*models.ForumResponse, error) {
	var payload models.ForumRequest
	err := c.BindJSON(&payload)
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	if err != nil {
		return nil, err
	}

	forum, err := usecase.repository.UpdateForum(c, payload, id)
	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (usecase *Usecase) DeleteForum(c *gin.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		return err
	}

	err = usecase.repository.DeleteForum(c, id)
	if err != nil {
		return err
	}

	return nil
}