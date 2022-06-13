package usecase

import (
	"kel15/models"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
	blogs, err := usecase.repository.GetAllBlog(c)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
