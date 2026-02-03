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

func NewSqlite(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.DriverType, error) {
	const op = "storage.driverTypes.GetAll"

	const query = `SELECT id, name, value FROM driver_types`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var driverTypes []models.DriverType
	for rows.Next() {
		var vc models.DriverType
		if err := rows.Scan(
			&vc.ID, &vc.Name, &vc.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		driverTypes = append(driverTypes, vc)
	}

	return driverTypes, nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, driverType models.DriverType) error {
	const op = "storage.driverType.InsertOrUpdate"

	const query = `
		INSERT INTO driver_types (
		name, value
		) VALUES (?, ?)
		ON CONFLICT(value) DO UPDATE SET 
			name = excluded.name
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		driverType.Name,
		driverType.Value,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
