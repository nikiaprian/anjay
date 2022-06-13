package repository

import (
	"kel15/models"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
	var blogs []models.Blog

	query := "SELECT * FROM Blogs"

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}
