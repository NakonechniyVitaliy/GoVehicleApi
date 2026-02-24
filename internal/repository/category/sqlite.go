package category

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
)

type SqliteRepository struct {
	db    *sql.DB
	table string
}

func NewSqliteCategoryRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	const op = "storage.categories.GetAll"

	const query = `SELECT id, name, value FROM categories`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var vc models.Category
		if err := rows.Scan(
			&vc.ID, &vc.Name, &vc.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		categories = append(categories, vc)
	}

	return categories, nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, category models.Category) error {
	const op = "storage.category.InsertOrUpdate"

	const query = `
		INSERT INTO categories (
		name, value
		) VALUES (?, ?)
		ON CONFLICT(value) DO UPDATE SET 
			name = excluded.name
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Value,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
