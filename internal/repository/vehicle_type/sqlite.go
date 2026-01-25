package vehicleType

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
		table: "vehicle_types",
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, vehicleTypeID int) (*models.VehicleType, error) {
	const op = "storage.vehicle_types.GetVehicleType"

	const query = `SELECT ablative, category_id, name, plural, rewrite, singular FROM vehicle_types WHERE id = ?`

	var vehicleType models.VehicleType
	err := s.db.QueryRowContext(ctx, query, vehicleTypeID).Scan(
		&vehicleType.Ablative, &vehicleType.CategoryID, &vehicleType.Name, &vehicleType.Plural, &vehicleType.Rewrite, &vehicleType.Singular,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: Error to return vehicleType %w", op, err)
	}

	return &vehicleType, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.VehicleType, error) {
	const op = "storage.vehicle_types.GetVehicleTypes"

	const query = `SELECT ablative, category_id, name, plural, rewrite, singular FROM vehicle_types`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var vehicle_types []models.VehicleType
	for rows.Next() {
		var vt models.VehicleType
		if err := rows.Scan(
			&vt.Ablative, &vt.CategoryID, &vt.Name, &vt.Plural, &vt.Rewrite, &vt.Singular); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		vehicle_types = append(vehicle_types, vt)
	}

	return vehicle_types, nil
}

func (s *SqliteRepository) Update(ctx context.Context, vehicleType models.VehicleType) error {
	const op = "storage.vehicleType.UpdateVehicleType"

	const query = `
		UPDATE vehicle_types
		SET
			ablative = ?,
			category_id = ?,
			name = ?,
			plural = ?,
			rewrite = ?,
			singular = ?
		WHERE id = ?
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		vehicleType.Ablative,
		vehicleType.CategoryID,
		vehicleType.Name,
		vehicleType.Plural,
		vehicleType.Rewrite,
		vehicleType.Singular,
		vehicleType.ID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrVehicleTypeNotFound
	}

	return nil
}

func (s *SqliteRepository) Delete(ctx context.Context, vehicleTypeID int) error {
	const op = "storage.vehicleType.DeleteVehicleType"

	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM vehicle_types WHERE id = ?",
		vehicleTypeID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrVehicleTypeNotFound
	}

	return nil
}

func (s *SqliteRepository) Create(ctx context.Context, vehicleType models.VehicleType) error {
	const op = "storage.vehicleType.NewVehicleType"

	const query = `
		INSERT INTO vehicle_types (
		ablative, 
	    category_id, 
	    name, 
	    plural, 
	    rewrite, 
	    singular
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		vehicleType.Ablative,
		vehicleType.CategoryID,
		vehicleType.Name,
		vehicleType.Plural,
		vehicleType.Rewrite,
		vehicleType.Name,
		vehicleType.Singular,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return storage.ErrVehicleTypeExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrVehicleTypeExists
	}

	return nil
}
