package body_style

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
)

type SqliteRepository struct {
	db    *sql.DB
	table string
}

func NewSqlite(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db:    db,
		table: "body_styles",
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, bodyStyleID uint16) (*models.BodyStyle, error) {
	const op = "storage.body_styles.getByID"

	const query = `SELECT id, name, value FROM body_styles WHERE id = ?`

	var bodyStyle models.BodyStyle
	err := s.db.QueryRowContext(ctx, query, bodyStyleID).Scan(
		&bodyStyle.ID, &bodyStyle.Name, &bodyStyle.Value,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: Error to return bodyStyle %w", op, err)
	}

	return &bodyStyle, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.BodyStyle, error) {
	const op = "storage.body_styles.GetBodyStyles"

	const query = `SELECT id, name, value FROM body_styles`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var body_styles []models.BodyStyle
	for rows.Next() {
		var bs models.BodyStyle
		if err := rows.Scan(
			&bs.ID, &bs.Name, &bs.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		body_styles = append(body_styles, bs)
	}

	return body_styles, nil
}

func (s *SqliteRepository) Update(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.bodyStyle.UpdateBodyStyle"

	const query = `
		UPDATE body_styles
		SET
			name = ?,
			value = ?
		WHERE id = ?
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		bodyStyle.Name,
		bodyStyle.Value,
		bodyStyle.ID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBodyStyleNotFound
	}

	return nil
}

func (s *SqliteRepository) Delete(ctx context.Context, bodyStyleID uint16) error {
	const op = "storage.bodyStyle.DeleteBodyStyle"

	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM body_styles WHERE id = ?",
		bodyStyleID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBodyStyleNotFound
	}

	return nil
}

func (s *SqliteRepository) Create(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.bodyStyle.NewBodyStyle"

	const query = `
		INSERT INTO body_styles (
		name, 
	    value
		) VALUES (?, ?)
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		bodyStyle.Name,
		bodyStyle.Value,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return storage.ErrBodyStyleExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBodyStyleExists
	}

	return nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.bodyStyle.InsertOrUpdate"

	const query = `
		INSERT INTO body_styles (
		name, value
		) VALUES (?, ?)
		ON CONFLICT(value) DO UPDATE SET 
			name = excluded.name,
			value = excluded.value
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		bodyStyle.Name,
		bodyStyle.Value,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
