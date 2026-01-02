package models

import "time"

// 🔧 UPDATED Book model: field names match repository (Views, Downloads), cover/file nullable pointers
type Book struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Year        int        `json:"year"`
	CategoryID  *int       `json:"category_id,omitempty"`
	Category    *Category  `json:"category,omitempty"`
	ISBN        string     `json:"isbn,omitempty"`
	Description string     `json:"description,omitempty"`
	Cover       *string    `json:"cover,omitempty"`   // stores filename (or full URL after handler formats)
	FileURL     *string    `json:"file_url,omitempty"` // stores filename (or full URL after handler formats)
	Views       int        `json:"views"`
	Downloads   int        `json:"downloads"`
	Publisher     string   `json:"publisher,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
