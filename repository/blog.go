package repository

import (
	"kel15/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
	
	query := `SELECT Blogs.id, Blogs.photo, Blogs.title, Blogs.content, Blogs.created_at, Blogs.updated_at,
			  Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at
			  FROM Blogs 
			  JOIN Users ON Blogs.user_id = Users.id`

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blog models.Blog
	var blogs []models.Blog
	var User models.User

	for rows.Next() {
		
		err := rows.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt,
				&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt)
		if err != nil {
			return nil, err
		}
		blog.User = User
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repository *Repository) CreateBlog(c *gin.Context, req models.BlogRequest, photo string, user_id int) (*models.Blog, error) {
	title := req.Title
	content := req.Content

	query := "INSERT INTO Blogs (photo, user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	row1, err := repository.db.Exec(query, photo, user_id, title, content, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := row1.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserById(c, int64(user_id))
	if err != nil {
		return nil, err
	}

	return &models.Blog{
		ID:        uint(id),
		User:      *user,
		Photo:     photo,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// func (repository *Repository) UpdateBlog(c *gin.Context, req models.BlogRequest, id int, photo string, user_id int) (*models.Blog, error) {
// 	title := req.Title
// 	content := req.Content

// 	query := `UPDATE Blogs INNER JOIN Users ON Blogs.user_id = Users.id
// 			  SET photo = ?, user_id = ?, title = ?, content = ?, created_at = ?, updated_at = ?
// 			  WHERE Blogs.id = ?`

// 	_, err := repository.db.Exec(query, photo, user_id, title, content, time.Now(), time.Now(), id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := repository.GetUserById(c, int64(id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Blog{
// 		ID:        uint(id),
// 		User:      *user,
// 		Photo:     photo,
// 		Title:     title,
// 		Content:   content,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}, nil
// }


func (repository *Repository) DeleteBlog(c *gin.Context, id int) error {
	query := "DELETE from Blogs where id=?"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) GetBlog(c *gin.Context, id int) ([]models.Blog, error) {
	query := `SELECT Blogs.id, Blogs.photo, Blogs.title, Blogs.content, Blogs.created_at, Blogs.updated_at,
			  Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at
			  FROM Blogs
			  JOIN Users ON Blogs.user_id = Users.id
			  WHERE Blogs.id = ?`

	row, err := repository.db.Query(query, c.Param("id"))
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var blog models.Blog
	var blogs []models.Blog
	var User models.User

	for row.Next() {
		err := row.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt,
				&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt)

		if err != nil {
			return nil, err
		}
		blog.User = User
		blogs = append(blogs, blog)

	}
	return blogs, nil
}
