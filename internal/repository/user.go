package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(data *models.User) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), `
		INSERT INTO users(email,password) VALUES
		($1,$2)
		RETURNING id,email,password
	`, data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &users, err
}

func (r *UserRepo) GetAllUsers() (*[]models.User, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT id,email,password FROM users
	`)

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &users, err
}

func (r *UserRepo) GetUserByEmail(data *models.User) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT id,email,password FROM users
		WHERE email=$1
	`, data.Email)

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &users, err
}

func (r *UserRepo) GetUserById(data *models.User) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT id,email,password FROM users
		WHERE id=$1
	`, data.Id)

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &users, err
}

func (r *UserRepo) UpdateUser(data *models.User) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), `
		UPDATE users
		SET email=COALESCE(NULLIF($2, ''), email),
		password=COALESCE(NULLIF($3, ''), password)
		WHERE id=$1
		RETURNING id, email, password
	`, data.Id, data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &users, err
}

func (r *UserRepo) DeleteUser(data *models.User) (*models.User, error) {
	rows, err := r.db.Query(context.Background(), `
		DELETE FROM users WHERE id=$1
		RETURNING id,email,password
	`, data.Id)

	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &users, err
}
