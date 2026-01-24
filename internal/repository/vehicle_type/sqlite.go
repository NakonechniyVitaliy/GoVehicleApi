package brand

import "database/sql"

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

func (s *SqliteRepository) GetVehicleType() error {
	return nil
}

func (s *SqliteRepository) NewVehicleType() error {
	return nil
}
func (s *SqliteRepository) DeleteVehicleType() error {
	return nil
}
