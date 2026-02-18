package driver_type

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	driverTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/driver_type"
	driverTypeRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/driver_types"
)

func RefreshDriverTypes(ctx context.Context, cfg *config.Config, repository driverTypeRepo.RepositoryInterface) error {

	vehicleCategories, err := driverTypeRequests.GetDriverTypes(cfg.AutoriaKey)
	if err != nil {
		return err
	}

	for _, oneCategory := range vehicleCategories {
		err = repository.InsertOrUpdate(ctx, oneCategory)
		if err != nil {
			return err
		}
	}

	return nil
}
