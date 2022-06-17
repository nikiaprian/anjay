package repository

import (
	"database/sql"
	"errors"
	"kel15/models"
	"kel15/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetUserById(c *gin.Context, id int64) (*models.User, error) {
	query := `SELECT id, email, username, created_at, updated_at FROM Users WHERE id = ? limit 1`
	row := repository.db.QueryRow(query, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByGoogleId(c *gin.Context, google_id string) (*models.User, error) {
	query := `SELECT Users.id, email, username, Users.created_at, Users.updated_at FROM Users JOIN GoogleAccounts ON Users.id = GoogleAccounts.user_id WHERE GoogleAccounts.google_id = ? limit 1`
	row := repository.db.QueryRow(query, google_id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByEmail(c *gin.Context, email string) (*models.User, error) {
	query := `SELECT id, email, username, password, created_at, updated_at FROM Users WHERE email = ? limit 1`
	row := repository.db.QueryRow(query, email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, errors.New("User not found")
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

func (repository *Repository) CreateUserGoogle(c *gin.Context, req models.UserGoogleResponse) (*models.User, error) {
	email := req.Email
	googleId := req.GoogleId
	username := strings.Split(email, "@")[0]

	user, err := repository.GetUserByGoogleId(c, googleId)

	if user != nil {
		return user, nil
	}

	if err != nil && user != nil {
		return nil, err
	}

	query := `INSERT INTO Users (email, username, provider, created_at, updated_at) VALUES (?, ?, 'google', ?, ?)`

	data, err := repository.db.Exec(query, email, username, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	id, err := data.LastInsertId()

	query = `INSERT INTO GoogleAccounts (user_id, google_id, created_at, updated_at) VALUES (?, ?, ?, ?)`

	data, err = repository.db.Exec(query, id, googleId, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	user, err = repository.GetUserById(c, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
