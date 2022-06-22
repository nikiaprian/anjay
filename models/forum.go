package models

import "time"

type Forum struct {
	ID        int        `json:"id"`
	User      User       `json:"user"`
	Title     string     `json:"title"`
	Tags      []Tag 	`json:"tag"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type ForumRequest struct {
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

type ForumResponse struct {
	Forum   []Forum `json:"forum"`
	Message string  `json:"message"`
}
