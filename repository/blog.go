package repository

import (
	"kel15/models"
	"time"

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
func (repository *Repository) CreateBlog(c *gin.Context, req models.BlogRequest) (*models.BlogResponse, error) {
	photo := req.Photo
	title := req.Title
	content := req.Content

	query := "INSERT INTO Blogs (photo, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	row1, err := repository.db.Exec(query, photo, title, content, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := row1.LastInsertId()
	if err != nil {
		return nil, err
	}
	query = "SELECT * FROM Blogs WHERE id = ?"
	row2 := repository.db.QueryRow(query, id)
	if err != nil {
		return nil, err
	}

	var blog models.Blog
	err = row2.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &models.BlogResponse{
		Blog:    &blog,
		Message: "Blog berhasil ditambahkan",
	}, nil
}
func (repository *Repository) UpdateBlog(c *gin.Context, req models.BlogRequest, id int) (*models.BlogResponse, error) {
	photo := req.Photo
	title := req.Title
	content := req.Content

	query := "UPDATE Blogs SET photo = ?, title = ?, content = ?,updated_at = ? WHERE id= ?"
	_, err := repository.db.Exec(query, photo, title, content, time.Now(), id)
	if err != nil {
		return nil, err
	}

	query = "SELECT * FROM Blogs WHERE id = ?"
	row := repository.db.QueryRow(query, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var blog models.Blog

	err = row.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &models.BlogResponse{
		Blog:    &blog,
		Message: "Blog berhasil diubah",
	}, nil
}

func (repository *Repository) DeleteBlog(c *gin.Context, id int) error {
	query := "DELETE from Blogs where id=?"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
