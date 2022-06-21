package models

import "time"

type CommentForum struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentForumRequest struct {
	Comment  string 	`json:"comment"`
}

type CommentForumResponse struct {
	CommentForum []CommentForum  `json:"comment_forum"`
	Message     string        `json:"message"`
}