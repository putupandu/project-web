package repositories


import (
"database/sql"
//

"e-library/backend/internal/models"
)


type SavedBookRepository struct {
DB *sql.DB
}


func NewSavedBookRepository(db *sql.DB) *SavedBookRepository {
return &SavedBookRepository{DB: db}
}


func (r *SavedBookRepository) Save(userID, bookID int) error {
_, err := r.DB.Exec(
`INSERT INTO saved_books (user_id, book_id)
VALUES ($1,$2) ON CONFLICT DO NOTHING`,
userID, bookID,
)
return err
}


func (r *SavedBookRepository) FindByUser(userID int) ([]models.Book, error) {
rows, err := r.DB.Query(`
SELECT b.id, b.title, b.author
FROM books b
JOIN saved_books sb ON sb.book_id = b.id
WHERE sb.user_id = $1`, userID)
if err != nil {
return nil, err
}
defer rows.Close()


var books []models.Book
for rows.Next() {
var b models.Book
if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
return nil, err
}
books = append(books, b)
}
return books, nil
}


func (r *SavedBookRepository) Delete(userID, bookID int) error {
_, err := r.DB.Exec(
`DELETE FROM saved_books WHERE user_id=$1 AND book_id=$2`,
userID, bookID,
)
return err
}