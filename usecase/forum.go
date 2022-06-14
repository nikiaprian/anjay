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