package models

import "time"

type CommentForum struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	ForumId   *int      `json:"forum_id,omitempty"`
	Comment   string    `json:"comment"`
	IsAnswer  bool      `json:"is_answer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentForumRequest struct {
	Comment string `json:"comment"`
}

type CommentForumResponse struct {
	CommentForum []CommentForum `json:"comment_forum"`
	Message      string         `json:"message"`
}

type SelectedCommentAnswerRequest struct {
	IsAnswer bool `json:"is_answer"`
}
