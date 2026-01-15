// File: internal/models/book.go
//
package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Year        int       `json:"year"`
	CategoryID  *int      `json:"category_id,omitempty"`
	Category    *Category `json:"category,omitempty"`
	ISBN        string    `json:"isbn,omitempty"`
	Description string    `json:"description,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	Language    string    `json:"language,omitempty"`
	Cover       *string   `json:"cover,omitempty"`
	FileURL     *string   `json:"file_url,omitempty"`
	Views       int       `json:"views"`
	Downloads   int       `json:"downloads"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
