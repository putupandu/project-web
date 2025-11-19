package repositories

import (
    "database/sql"
    "e-library/backend/internal/models"
)

type BookRepository struct {
    DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
    return &BookRepository{DB: db}
}

func (r *BookRepository) FindAll(filters map[string]interface{}) ([]models.Book, error) {
    rows, err := r.DB.Query("SELECT id, title, author, description, category_id, file_url, cover, views, downloads FROM books")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    books := []models.Book{}
    for rows.Next() {
        var b models.Book
        var categoryID sql.NullInt64
        var filePath sql.NullString
        var coverPath sql.NullString

        if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Description, &categoryID, &filePath, &coverPath, &b.ViewCount, &b.DownloadCount); err != nil {
            return nil, err
        }

        if categoryID.Valid {
            b.CategoryID = new(int)
            *b.CategoryID = int(categoryID.Int64)
        }
        if filePath.Valid {
            b.FilePath = new(string)
            *b.FilePath = filePath.String
        }
        if coverPath.Valid {
            b.CoverPath = new(string)
            *b.CoverPath = coverPath.String
        }

        books = append(books, b)
    }
    return books, nil
}

func (r *BookRepository) FindByID(id int) (*models.Book, error) {
    var b models.Book
    var categoryID sql.NullInt64
    var filePath sql.NullString
    var coverPath sql.NullString

    err := r.DB.QueryRow(
        "SELECT id, title, author, description, category_id, file_url, cover, views, downloads FROM books WHERE id=$1", id,
    ).Scan(&b.ID, &b.Title, &b.Author, &b.Description, &categoryID, &filePath, &coverPath, &b.ViewCount, &b.DownloadCount)

    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }

    if categoryID.Valid {
        b.CategoryID = new(int)
        *b.CategoryID = int(categoryID.Int64)
    }
    if filePath.Valid {
        b.FilePath = new(string)
        *b.FilePath = filePath.String
    }
    if coverPath.Valid {
        b.CoverPath = new(string)
        *b.CoverPath = coverPath.String
    }

    return &b, nil
}

func (r *BookRepository) Create(b *models.Book) error {
    _, err := r.DB.Exec(
        `INSERT INTO books(title, author, description, category_id, file_url, cover, views, downloads) 
        VALUES($1,$2,$3,$4,$5,$6,$7,$8)`,
        b.Title, b.Author, b.Description, b.CategoryID, b.FilePath, b.CoverPath, b.ViewCount, b.DownloadCount,
    )
    return err
}

func (r *BookRepository) Update(id int, b *models.Book) error {
    _, err := r.DB.Exec(
        `UPDATE books SET title=$1, author=$2, description=$3, category_id=$4, file_url=$5, cover=$6, views=$7, downloads=$8, updated_at=NOW() 
        WHERE id=$9`,
        b.Title, b.Author, b.Description, b.CategoryID, b.FilePath, b.CoverPath, b.ViewCount, b.DownloadCount, id,
    )
    return err
}

func (r *BookRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM books WHERE id=$1", id)
    return err
}

func (r *BookRepository) IncrementDownloads(id int) error {
    _, err := r.DB.Exec("UPDATE books SET downloads = downloads + 1 WHERE id=$1", id)
    return err
}

func (r *BookRepository) IncrementViews(id int) error {
    _, err := r.DB.Exec("UPDATE books SET views = views + 1 WHERE id=$1", id)
    return err
}

func (r *BookRepository) Count(filters map[string]interface{}) (int, error) {
    var count int
    err := r.DB.QueryRow("SELECT COUNT(*) FROM books").Scan(&count)
    return count, err
}
