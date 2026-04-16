package gearbox

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NakonechniyVitalii/GoVehicleApi/internal/models"
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteGearboxRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, gearboxID uint16) (*models.Gearbox, error) {
	const op = "storage.gearboxes.get_by_id"

	const query = `SELECT id, name, value FROM gearboxes WHERE id = ?`

	var gearbox models.Gearbox
	err := s.db.QueryRowContext(ctx, query, gearboxID).Scan(
		&gearbox.ID, &gearbox.Name, &gearbox.Value,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: Error to return gearbox %w", op, err)
	}

	return &gearbox, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.Gearbox, error) {
	const op = "storage.gearbox.get_all"

	const query = `SELECT id, name, value FROM gearboxes`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var gearboxes []models.Gearbox
	for rows.Next() {
		var g models.Gearbox
		if err := rows.Scan(
			&g.ID, &g.Name, &g.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		gearboxes = append(gearboxes, g)
	}

	return gearboxes, nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, gearbox models.Gearbox) error {
	const op = "storage.gearbox.insert_or_update"

	const query = `
		INSERT INTO gearboxes (
		name, value
		) VALUES (?, ?)
		ON CONFLICT(value) DO UPDATE SET 
			name = excluded.name
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		gearbox.Name,
		gearbox.Value,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
