package repository

import "kel15/services"

type Repository struct {
	db *services.SQLLite3
}

func NewRepository(db *services.SQLLite3) Repository {
	return Repository{db: db}
}
