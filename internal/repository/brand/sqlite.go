package brand

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

func NewSqliteBrandRepo(db *sql.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}

func (s *SqliteRepository) GetByID(ctx context.Context, brandID uint16) (*models.Brand, error) {
	const op = "storage.brands.GetByID"

	const query = `SELECT id, marka_id, category_id, cnt, country_id, eng, name, slang, value FROM brands WHERE id = ?`

	var brand models.Brand
	err := s.db.QueryRowContext(ctx, query, brandID).Scan(
		&brand.ID, &brand.MarkaID, &brand.CategoryID, &brand.Count, &brand.CountryID, &brand.EngName, &brand.Name, &brand.Slang, &brand.Value,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrBrandNotFound
		}
		return nil, fmt.Errorf("%s: Error to return brand %w", op, err)
	}

	return &brand, nil
}

func (s *SqliteRepository) GetAll(ctx context.Context) ([]models.Brand, error) {
	const op = "storage.brands.GetAll"

	const query = `SELECT marka_id, category_id, cnt, country_id, eng, name, slang, value FROM brands`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: exec: %w", op, err)
	}
	defer rows.Close()

	var brands []models.Brand
	for rows.Next() {
		var b models.Brand
		if err := rows.Scan(
			&b.CategoryID, &b.Count, &b.CountryID, &b.MarkaID, &b.EngName, &b.Name, &b.Slang, &b.Value); err != nil {
			return nil, fmt.Errorf("%s: scan: %w", op, err)
		}
		brands = append(brands, b)
	}

	return brands, nil
}

func (s *SqliteRepository) Update(ctx context.Context, brand models.Brand, brandID uint16) (*models.Brand, error) {
	const op = "storage.brand.Update"

	const query = `
		UPDATE brands
		SET
			category_id = ?,
			cnt = ?,
			country_id = ?,
			eng = ?,
			name = ?,
			slang = ?,
			marka_id = ?
		WHERE id = ?`

	res, err := s.db.ExecContext(
		ctx,
		query,
		brand.CategoryID,
		brand.Count,
		brand.CountryID,
		brand.EngName,
		brand.Name,
		brand.Slang,
		brand.MarkaID,
		brandID,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return nil, storage.ErrBrandNotFound
	}

	updatedBrand, err := s.GetByID(ctx, brandID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return updatedBrand, nil
}

func (s *SqliteRepository) Delete(ctx context.Context, brandID uint16) error {
	const op = "storage.brand.DeleteBrand"

	res, err := s.db.ExecContext(
		ctx,
		"DELETE FROM brands WHERE id = ?",
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

func (s *SqliteRepository) Create(ctx context.Context, brand models.Brand) (*models.Brand, error) {
	const op = "storage.brand.NewBrand"

	const query = `
		INSERT INTO brands (
			marka_id,
			category_id,
			cnt,
			country_id,
			eng,
			name,
			slang,
			value
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	res, err := s.db.ExecContext(
		ctx,
		query,
		brand.MarkaID,
		brand.CategoryID,
		brand.Count,
		brand.CountryID,
		brand.EngName,
		brand.Name,
		brand.Slang,
		brand.Value,
	)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return nil, storage.ErrBrandExists
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if affected == 0 {
		return nil, storage.ErrBrandExists
	}

	createdBrandID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	createdBrand, err := s.GetByID(ctx, uint16(createdBrandID))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return createdBrand, nil
}

func (s *SqliteRepository) InsertOrUpdate(ctx context.Context, brand models.Brand) error {
	const op = "storage.brand.InsertOrUpdate"

	const query = `
		INSERT INTO brands (
			marka_id,
			category_id,
			cnt,
			country_id,
			eng,
			name,
			slang,
			value
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(marka_id) DO UPDATE SET 
			category_id = excluded.category_id,
			cnt = excluded.cnt,
			country_id = excluded.country_id,
			eng = excluded.eng,
			name = excluded.name,
			slang = excluded.slang,
			value = excluded.value
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		brand.MarkaID,
		brand.CategoryID,
		brand.Count,
		brand.CountryID,
		brand.EngName,
		brand.Name,
		brand.Slang,
		brand.Value,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
