package body_style

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitalii/GoVehicleApi/internal/repository/_errors"
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteBodyStyleRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, bodyStyleID uint16) (*models.BodyStyle, error) {
	const op = "storage.body_styles.get_by_id"

	const query = `SELECT id, name, value FROM body_styles WHERE id = ?`

	var bodyStyle models.BodyStyle
	err := s.db.QueryRowContext(ctx, query, bodyStyleID).Scan(
		&bodyStyle.ID, &bodyStyle.Name, &bodyStyle.Value,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, _errors.ErrBodyStyleNotFound
		}
		return nil, fmt.Errorf("%s: error to return body style %w", op, err)
	}

	return &bodyStyle, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.BodyStyle, error) {
	const op = "storage.body_styles.get_all"

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

func (s *SqliteRepository) Update(
	ctx context.Context,
	bodyStyle models.BodyStyle,
	bodyStyleID uint16,
) (*models.BodyStyle, error) {
	const op = "storage.body_style.update"

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
		bodyStyleID,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return nil, _errors.ErrBodyStyleNotFound
	}

	createdBodyStyleID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	createdBS, err := s.GetByID(ctx, uint16(createdBodyStyleID))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return createdBS, nil
}

func (s *SqliteRepository) Delete(ctx context.Context, bodyStyleID uint16) error {
	const op = "storage.body_style.delete"

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
		return _errors.ErrBodyStyleNotFound
	}

	return nil
}

func (s *SqliteRepository) Create(ctx context.Context, bodyStyle models.BodyStyle) (*models.BodyStyle, error) {
	const op = "storage.body_style.create"

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
			return nil, _errors.ErrBodyStyleExists
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return nil, _errors.ErrBodyStyleExists
	}

	bodyStyleID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	createdBS, err := s.GetByID(ctx, uint16(bodyStyleID))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return createdBS, nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, bodyStyle models.BodyStyle) error {
	const op = "storage.body_style.insert_or_update"

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
