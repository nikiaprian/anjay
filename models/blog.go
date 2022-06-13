package models

import "time"

type Blog struct {
	ID        uint      `json:"id"`
	Photo     string    `json:"photo"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BlogRequest struct {
	Photo   string `json:"photo"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BlogResponse struct {
	Blog    *Blog  `json:"blog"`
	Message string `json:"message"`
}
