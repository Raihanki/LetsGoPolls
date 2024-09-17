package repositories

import (
	"database/sql"

	"github.com/Raihanki/LetsGoPolls/internal/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) CreateUser(request entities.CreateUserRequest) (entities.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *"
	row := repo.DB.QueryRow(query, request.Name, request.Email, request.Password)
	user := entities.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (entities.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1"
	row := repo.DB.QueryRow(query, email)

	user := entities.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserById(id int) (entities.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1"
	row := repo.DB.QueryRow(query, id)

	user := entities.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
