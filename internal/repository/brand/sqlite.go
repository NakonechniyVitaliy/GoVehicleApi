package brand

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
		db: db,
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, brandID int) (*models.Brand, error) {
	const op = "storage.brands.GetBrand"

	const query = `SELECT category_id, cnt, country_id, eng, marka_id, name, slang, value FROM brands WHERE marka_id = ?`

	var brand models.Brand
	err := s.db.QueryRowContext(ctx, query, brandID).Scan(
		&brand.Category, &brand.Count, &brand.Country, &brand.EngName, &brand.MarkaID, &brand.Name, &brand.Slang, &brand.Value,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: Error to return brand %w", op, err)
	}

	return &brand, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	const op = "storage.brands.GetBrands"

	const query = `SELECT category_id, cnt, country_id, eng, marka_id, name, slang, value FROM brands`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var brands []models.Brand
	for rows.Next() {
		var b models.Brand
		if err := rows.Scan(
			&b.Category, &b.Count, &b.Country, &b.EngName, &b.MarkaID, &b.Name, &b.Slang, &b.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		brands = append(brands, b)
	}

	return brands, nil
}

func (s *SqliteRepository) Update(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.UpdateBrand"

	const query = `
		UPDATE brands
		SET
			category_id = ?,
			cnt = ?,
			country_id = ?,
			eng = ?,
			name = ?,
			slang = ?
		WHERE marka_id = ?
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		brand.Category,
		brand.Count,
		brand.Country,
		brand.EngName,
		brand.Name,
		brand.Slang,
		brand.MarkaID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBrandNotFound
	}

	return nil
}

func (s *SqliteRepository) Delete(ctx context.Context, brandID int) error {
	const op = "storage.brand.DeleteBrand"

	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM brands WHERE marka_id = ?",
		brandID,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBrandNotFound
	}

	return nil
}

func (s *SqliteRepository) Create(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.NewBrand"

	const query = `
		INSERT INTO brands (
			category_id,
			cnt,
			country_id,
			eng,
			marka_id,
			name,
			slang,
			value
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		brand.Category,
		brand.Count,
		brand.Country,
		brand.EngName,
		brand.MarkaID,
		brand.Name,
		brand.Slang,
		brand.Value,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return storage.ErrBrandExists
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if affected == 0 {
		return storage.ErrBrandExists
	}

	return nil
}

func (s *SqliteRepository) RefreshBrands() error {
	return nil
}
