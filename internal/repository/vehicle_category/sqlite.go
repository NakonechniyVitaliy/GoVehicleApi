package vehicle_category

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

func NewSqliteVehicleCategoryRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.VehicleCategory, error) {
	const op = "storage.vehicleCategories.GetAll"

	const query = `SELECT id, name, value FROM vehicle_categories`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var vehicleCategories []models.VehicleCategory
	for rows.Next() {
		var vc models.VehicleCategory
		if err := rows.Scan(
			&vc.ID, &vc.Name, &vc.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		vehicleCategories = append(vehicleCategories, vc)
	}

	return vehicleCategories, nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, vehicleCategory models.VehicleCategory) error {
	const op = "storage.vehicleCategory.InsertOrUpdate"

	const query = `
		INSERT INTO vehicle_categories (
		name, value
		) VALUES (?, ?)
		ON CONFLICT(value) DO UPDATE SET 
			name = excluded.name
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		vehicleCategory.Name,
		vehicleCategory.Value,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
