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
    var c models.Category
    err := r.DB.QueryRow("SELECT id, name, slug FROM categories WHERE id=$1", id).Scan(&c.ID, &c.Name, &c.Slug)
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return &c, nil
}

func (r *CategoryRepository) FindAll() ([]models.Category, error) {
    rows, err := r.DB.Query("SELECT id, name, slug FROM categories")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    categories := []models.Category{}
    for rows.Next() {
        var c models.Category
        if err := rows.Scan(&c.ID, &c.Name, &c.Slug); err != nil {
            return nil, err
        }
        categories = append(categories, c)
    }
    return categories, nil
}

func (r *CategoryRepository) FindBySlug(slug string) (*models.Category, error) {
    var c models.Category
    err := r.DB.QueryRow("SELECT id, name, slug FROM categories WHERE slug=$1", slug).Scan(&c.ID, &c.Name, &c.Slug)
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }
    return &c, nil
}
