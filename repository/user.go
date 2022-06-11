package repository

import (
	"database/sql"
	"errors"
	"kel15/models"
	"kel15/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetUserByEmail(c *gin.Context, req models.UserLoginRequest) (*models.User, error) {
	email := req.Email
	query := `SELECT id, email, username, password, created_at, updated_at FROM Users WHERE email = ? limit 1`
	row := repository.db.QueryRow(query, email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByEmailOrUsername(c *gin.Context, req models.UserRegisterRequest) (*models.User, error) {
	email := req.Email
	username := req.Username
	query := `SELECT id, email, password, created_at, updated_at FROM Users WHERE email = ? or username = ? limit 1`
	row := repository.db.QueryRow(query, email, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) CreateUser(c *gin.Context, req models.UserRegisterRequest) error {
	email := req.Email
	username := req.Username
	password, err := utils.GeneratePassword(req.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO Users (email, username, password, provider, created_at, updated_at) VALUES (?, ?, ?, 'local', ?, ?)`
	_, err = repository.db.Exec(query, email, username, password, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}
