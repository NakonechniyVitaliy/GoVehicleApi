package vehicle

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/models"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/storage"
)

type SqliteRepository struct {
	db    *sql.DB
	table string
}

func NewSqliteVehicleRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, vehicleID uint16) (*models.Vehicle, error) {
	const op = "storage.vehicles.getByID"

	const query = `SELECT id, brand, driver_type, gearbox, body_style, category, mileage, model, price FROM vehicles WHERE id = ?`

	var vehicle models.Vehicle
	err := s.db.QueryRowContext(ctx, query, vehicleID).Scan(
		&vehicle.ID, &vehicle.Brand, &vehicle.DriverType, &vehicle.Gearbox, &vehicle.BodyStyle, &vehicle.Category, &vehicle.Mileage, &vehicle.Model, &vehicle.Price,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrVehicleNotFound
		}
		return nil, fmt.Errorf("%s: Error to return vehicle %w", op, err)
	}

	return &vehicle, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.Vehicle, error) {
	const op = "storage.vehicles.getAll"

	const query = `SELECT id, brand, driver_type, gearbox, body_style, category, mileage, model, price FROM vehicles`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var vehicles []models.Vehicle
	for rows.Next() {
		var v models.Vehicle
		if err := rows.Scan(
			&v.ID, &v.Brand, &v.DriverType, &v.Gearbox, &v.BodyStyle, &v.Category, &v.Mileage, &v.Model, &v.Price); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		vehicles = append(vehicles, v)
	}

	return vehicles, nil
}

func (s *SqliteRepository) Update(ctx context.Context, vehicle models.Vehicle, vehicleID uint16) (*models.Vehicle, error) {
	const op = "storage.vehicle.update"

	const query = `
		UPDATE vehicles
		SET
			brand = ?,
			driver_type = ?,
			gearbox = ?,
			body_style = ?,
			category = ?,
			mileage = ?,
			model = ?,
			price = ?
		WHERE id = ?`

	res, err := s.db.ExecContext(
		ctx,
		query,
		vehicle.Brand,
		vehicle.DriverType,
		vehicle.Gearbox,
		vehicle.BodyStyle,
		vehicle.Category,
		vehicle.Mileage,
		vehicle.Model,
		vehicle.Price,
		vehicleID,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return nil, storage.ErrVehicleNotFound
	}

	updatedVehicle, err := s.GetByID(ctx, vehicleID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return updatedVehicle, nil
}

func (s *SqliteRepository) Delete(ctx context.Context, vehicleID uint16) error {
	const op = "storage.vehicle.delete"

	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM vehicles WHERE id = ?",
		vehicleID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrVehicleNotFound
	}

	return nil
}

func (s *SqliteRepository) Create(ctx context.Context, vehicle models.Vehicle) (*models.Vehicle, error) {
	const op = "storage.vehicle.NewVehicle"

	const query = `
		INSERT INTO vehicles (
			  brand, 
			  driver_type, 
			  gearbox, 
			  body_style, 
			  category, 
			  mileage, 
			  model, 
			  price
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		vehicle.Brand,
		vehicle.DriverType,
		vehicle.Gearbox,
		vehicle.BodyStyle,
		vehicle.Category,
		vehicle.Mileage,
		vehicle.Model,
		vehicle.Price,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return nil, storage.ErrVehicleExists
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return nil, storage.ErrVehicleExists
	}

	createdVehicleID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	createdVehicle, err := s.GetByID(ctx, uint16(createdVehicleID))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return createdVehicle, nil
}
