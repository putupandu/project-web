package repositories

import (
	"database/sql"
	"e-library/backend/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email, created_at, updated_at FROM users ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id)
	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

func (r *UserRepository) Create(name, email string) (*models.User, error) {
	row := r.DB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at, updated_at",
		name, email,
	)
	var u models.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	return &u, err
}

func (r *UserRepository) Update(id int, name, email string) error {
	_, err := r.DB.Exec(
		"UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3",
		name, email, id,
	)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}