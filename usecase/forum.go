package usecase

import (
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
