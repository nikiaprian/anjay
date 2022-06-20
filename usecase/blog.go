package usecase

import (
	"kel15/models"
	"strconv"

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
	fileName, _ := c.Get("file")

	var payload models.BlogRequest

	blog, err := usecase.repository.CreateBlog(c, payload, fileName.(string))
	if err != nil {
		return nil, err
	}

	return blog, nil
}
func (usecase *Usecase) UpdateBlog(c *gin.Context) (*models.BlogResponse, error) {
	var payload models.BlogRequest
	err := c.BindJSON(&payload)
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	if err != nil {
		return nil, err
	}

	blog, err := usecase.repository.UpdateBlog(c, payload, id, "")
	if err != nil {
		return nil, err
	}

	return blog, nil
}
func (usecase *Usecase) DeleteBlog(c *gin.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		return err
	}

	err = usecase.repository.DeleteBlog(c, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) GetBlog(c *gin.Context) (*models.BlogResponse, error) {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	blog, err := usecase.repository.GetBlog(c, id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}
