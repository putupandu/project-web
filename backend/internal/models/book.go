package models

type Book struct {
    ID            int     `json:"id"`
    Title         string  `json:"title"`
    Author        string  `json:"author"`
    Description   string  `json:"description,omitempty"`
    CategoryID    *int    `json:"category_id,omitempty"`
    FilePath      *string `json:"file_path,omitempty"`
    CoverPath     *string `json:"cover_path,omitempty"`
    ViewCount     int     `json:"view_count"`
    DownloadCount int     `json:"download_count"`
}
