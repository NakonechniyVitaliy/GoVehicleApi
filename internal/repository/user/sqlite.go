package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/_errors"
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteUserRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetByLogin(ctx context.Context, login string) (*models.User, error) {
	const op = "storage.user.get_by_login"

	const query = `SELECT id, username, login, password FROM users WHERE login = ?`

	var user models.User
	err := s.db.QueryRowContext(ctx, query, login).Scan(&user.ID, &user.Username, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _errors.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: Error to return user %w", op, err)
	}

	return &user, nil
}

func (s *SqliteRepository) Create(ctx context.Context, user models.User) error {
	const op = "storage.user.create"

	const query = `INSERT INTO users (username, login, password) VALUES (?, ?, ?)`

	res, err := s.db.ExecContext(ctx, query, user.Username, user.Login, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return _errors.ErrUserExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return _errors.ErrUserExists
	}

	return nil
}
