package repositories

import (
	"database/sql"
	"e-library/backend/internal/models"
	"fmt"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

//  FIND ALL 
func (r *BookRepository) FindAll(filters map[string]interface{}) ([]models.Book, error) {
	query := `
		SELECT 
			b.id, b.title, b.author, b.year, b.language, b.description,
			b.category_id, b.file_url, b.cover, b.publisher,
			b.views, b.downloads, b.created_at, b.updated_at,
			c.id, c.name
		FROM books b
		LEFT JOIN categories c ON b.category_id = c.id
		WHERE 1=1
	`

	var args []interface{}

	if search, ok := filters["search"].(string); ok && search != "" {
		query += fmt.Sprintf(" AND LOWER(b.title) LIKE LOWER($%d)", len(args)+1)
		args = append(args, "%"+search+"%")
	}

	if categoryID, ok := filters["category_id"].(int); ok && categoryID > 0 {
		query += fmt.Sprintf(" AND b.category_id = $%d", len(args)+1)
		args = append(args, categoryID)
	}

	page := 1
	perPage := 12
	if p, ok := filters["page"].(int); ok && p > 0 {
		page = p
	}
	if pp, ok := filters["per_page"].(int); ok && pp > 0 {
		perPage = pp
	}

	offset := (page - 1) * perPage
	query += fmt.Sprintf(" ORDER BY b.id DESC LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, perPage, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var b models.Book

		var year sql.NullInt64
		var language sql.NullString
		var categoryID sql.NullInt64
		var categoryName sql.NullString
		var fileURL sql.NullString
		var cover sql.NullString
		var publisher sql.NullString
		var createdAt sql.NullTime
		var updatedAt sql.NullTime

		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&year,
			&language,
			&b.Description,
			&categoryID,
			&fileURL,
			&cover,
			&publisher,
			&b.Views,
			&b.Downloads,
			&createdAt,
			&updatedAt,
			&categoryID,
			&categoryName,
		)
		if err != nil {
			return nil, err
		}

		if year.Valid {
			b.Year = int(year.Int64)
		}
		if language.Valid {
			b.Language = language.String
		}
		if fileURL.Valid {
			s := fileURL.String
			b.FileURL = &s
		}
		if cover.Valid {
			s := cover.String
			b.Cover = &s
		}
		if publisher.Valid {
			b.Publisher = publisher.String
		}
		if createdAt.Valid {
			b.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			b.UpdatedAt = updatedAt.Time
		}

		if categoryID.Valid && categoryName.Valid {
			id := int(categoryID.Int64)
			b.CategoryID = &id
			b.Category = &models.Category{
				ID:   id,
				Name: categoryName.String,
			}
		}

		books = append(books, b)
	}

	return books, nil
}

//  COUNT 
func (r *BookRepository) Count(filters map[string]interface{}) (int, error) {
	query := `SELECT COUNT(*) FROM books WHERE 1=1`
	var args []interface{}

	if search, ok := filters["search"].(string); ok && search != "" {
		query += fmt.Sprintf(" AND LOWER(title) LIKE LOWER($%d)", len(args)+1)
		args = append(args, "%"+search+"%")
	}

	if categoryID, ok := filters["category_id"].(int); ok && categoryID > 0 {
		query += fmt.Sprintf(" AND category_id = $%d", len(args)+1)
		args = append(args, categoryID)
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

//  FIND BY ID 
func (r *BookRepository) FindByID(id int) (*models.Book, error) {
	var b models.Book

	var year sql.NullInt64
	var language sql.NullString
	var categoryID sql.NullInt64
	var categoryName sql.NullString
	var fileURL sql.NullString
	var cover sql.NullString
	var publisher sql.NullString
	var createdAt sql.NullTime
	var updatedAt sql.NullTime

	err := r.DB.QueryRow(`
		SELECT 
			b.id, b.title, b.author, b.year, b.language, b.description,
			b.category_id, b.file_url, b.cover, b.publisher,
			b.views, b.downloads, b.created_at, b.updated_at,
			c.id, c.name
		FROM books b
		LEFT JOIN categories c ON b.category_id = c.id
		WHERE b.id = $1
	`, id).Scan(
		&b.ID,
		&b.Title,
		&b.Author,
		&year,
		&language,
		&b.Description,
		&categoryID,
		&fileURL,
		&cover,
		&publisher,
		&b.Views,
		&b.Downloads,
		&createdAt,
		&updatedAt,
		&categoryID,
		&categoryName,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if year.Valid {
		b.Year = int(year.Int64)
	}
	if language.Valid {
		b.Language = language.String
	}
	if fileURL.Valid {
		s := fileURL.String
		b.FileURL = &s
	}
	if cover.Valid {
		s := cover.String
		b.Cover = &s
	}
	if publisher.Valid {
		b.Publisher = publisher.String
	}
	if createdAt.Valid {
		b.CreatedAt = createdAt.Time
	}
	if updatedAt.Valid {
		b.UpdatedAt = updatedAt.Time
	}

	if categoryID.Valid && categoryName.Valid {
		id := int(categoryID.Int64)
		b.CategoryID = &id
		b.Category = &models.Category{
			ID:   id,
			Name: categoryName.String,
		}
	}

	return &b, nil
}

//  CREATE 
func (r *BookRepository) Create(b *models.Book) error {
	_, err := r.DB.Exec(`
		INSERT INTO books
			(title, author, year, language, description, category_id, file_url, cover, publisher,
			 views, downloads, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,0,0,NOW(),NOW())
	`,
		b.Title,
		b.Author,
		b.Year,
		b.Language,
		b.Description,
		nullableInt(b.CategoryID),
		nullableString(b.FileURL),
		nullableString(b.Cover),
		b.Publisher,
	)
	return err
}

//  UPDATE 
func (r *BookRepository) Update(id int, b *models.Book) error {
	_, err := r.DB.Exec(`
		UPDATE books SET
			title=$1, author=$2, year=$3, language=$4, description=$5,
			category_id=$6, file_url=$7, cover=$8, publisher=$9,
			updated_at=NOW()
		WHERE id=$10
	`,
		b.Title,
		b.Author,
		b.Year,
		b.Language,
		b.Description,
		nullableInt(b.CategoryID),
		nullableString(b.FileURL),
		nullableString(b.Cover),
		b.Publisher,
		id,
	)
	return err
}

//  DELETE 
func (r *BookRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM books WHERE id=$1`, id)
	return err
}

//  INCREMENT 
func (r *BookRepository) IncrementViews(id int) error {
	_, err := r.DB.Exec(`UPDATE books SET views = views + 1 WHERE id=$1`, id)
	return err
}

func (r *BookRepository) IncrementDownloads(id int) error {
	_, err := r.DB.Exec(`UPDATE books SET downloads = downloads + 1 WHERE id=$1`, id)
	return err
}

//  HELPERS 
func nullableString(v *string) interface{} {
	if v == nil {
		return nil
	}
	return *v
}

func nullableInt(v *int) interface{} {
	if v == nil {
		return nil
	}
	return *v
}
