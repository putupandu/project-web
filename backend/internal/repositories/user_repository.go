package repositories

import (
	"database/sql"
	"e-library/backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email=$1",
		email,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Create(u *models.User) error {
	return r.db.QueryRow(
		"INSERT INTO users(name, email, password) VALUES($1,$2,$3) RETURNING id",
		u.Name, u.Email, u.Password,
	).Scan(&u.ID)
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	u := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, name, email FROM users WHERE id=$1",
		id,
	).Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Update(u *models.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name=$1, email=$2 WHERE id=$3",
		u.Name, u.Email, u.ID,
	)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
