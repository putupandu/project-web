package models

import "time"
//
type SavedBook struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	BookID    int       `json:"book_id"`
	Book      Book      `json:"book"`
	CreatedAt time.Time `json:"created_at"`
}
