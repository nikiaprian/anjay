package models

import "time"

type Blog struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Photo     string    `json:"photo"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tag 	  []BlogTag `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type BlogTag struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}
type BlogRequest struct {
	// Photo   string 	`json:"photo,omitempty"`
	Title    string 	`json:"title"`
	Tags     []string 	`json:"tag"`
	Content  string 	`json:"content"`
}

type BlogResponse struct {
	Blog    []Blog  `json:"blog"`
	Message string `json:"message"`
}
