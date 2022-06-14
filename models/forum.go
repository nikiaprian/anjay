package models

import "time"

type Forum struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Tag       string    `json:"tag"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ForumRequest struct {
	Title   string `json:"title"`
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type ForumResponse struct {
	Forum   *Forum `json:"forum"`
	Message string `json:"message"`
}
