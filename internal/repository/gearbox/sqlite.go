package driver_type

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

func NewSqliteGearboxRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.Gearbox, error) {
	const op = "storage.gearbox.GetAll"

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
	const op = "storage.gearbox.InsertOrUpdate"

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
