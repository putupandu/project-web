package repositories

import (
    "database/sql"
    "e-library/backend/internal/models"
)

type CategoryRepository struct {
    DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
    return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) FindByID(id int) (*models.Category, error) {

    // ⚠️ DIUBAH — tambahkan created_at dan updated_at agar sesuai struct
    query := `
        SELECT id, name, slug, description, created_at, updated_at
        FROM categories 
        WHERE id = $1
    `

    var c models.Category
    err := r.DB.QueryRow(query, id).Scan(
        &c.ID,
        &c.Name,
        &c.Slug,        // ⚠️ DIUBAH — urutan slug & description dibetulkan
        &c.Description,
        &c.CreatedAt,   // ⚠️ DIUBAH — ditambahkan
        &c.UpdatedAt,   // ⚠️ DIUBAH — ditambahkan
    )
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {

    // ⚠️ DIUBAH — tambahkan created_at & updated_at
    query := `
        SELECT id, name, slug, description, created_at, updated_at
        FROM categories
        ORDER BY id DESC
    `

    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    categories := []models.Category{}
    for rows.Next() {
        var c models.Category
        if err := rows.Scan(
            &c.ID,
            &c.Name,
            &c.Slug,
            &c.Description,
            &c.CreatedAt,   // ⚠️ DIUBAH
            &c.UpdatedAt,   // ⚠️ DIUBAH
        ); err != nil {
            return nil, err
        }
        categories = append(categories, c)
    }
    return categories, nil
}

func (r *CategoryRepository) FindBySlug(slug string) (*models.Category, error) {

    // ⚠️ DIUBAH — tambahkan created_at & updated_at
    query := `
        SELECT id, name, slug, description, created_at, updated_at
        FROM categories
        WHERE slug = $1
    `

    var c models.Category
    err := r.DB.QueryRow(query, slug).Scan(
        &c.ID,
        &c.Name,
        &c.Slug,
        &c.Description,
        &c.CreatedAt,   // ⚠️ DIUBAH
        &c.UpdatedAt,   // ⚠️ DIUBAH
    )
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *CategoryRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM categories WHERE id = $1", id)
    return err
}

func (r *CategoryRepository) Create(cat *models.Category) (*models.Category, error) {

    // ⚠️ DIUBAH — urutan kolom harus sama dengan struct
    query := `
        INSERT INTO categories (name, slug, description)
        VALUES ($1, $2, $3)
        RETURNING id, name, slug, description, created_at, updated_at
    `

    var newCat models.Category
    err := r.DB.QueryRow(query, cat.Name, cat.Slug, cat.Description).Scan(
        &newCat.ID,
        &newCat.Name,
        &newCat.Slug,        // ⚠️ DIUBAH — pastikan posisi benar
        &newCat.Description,
        &newCat.CreatedAt,
        &newCat.UpdatedAt,
    )
    return &newCat, err
}
