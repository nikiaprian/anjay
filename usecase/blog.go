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
func (usecase *Usecase) CreateBlog(c *gin.Context) (*models.BlogResponse, error) {
	var payload models.BlogRequest
	err := c.BindJSON(&payload)
	if err != nil {
		return nil, err
	}

	blog, err := usecase.repository.CreateBlog(c, payload)
	if err != nil {
		return nil, err
	}

	return blog, nil
}
