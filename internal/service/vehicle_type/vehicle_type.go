package vehicle_category

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	vehicleTypeRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_type"
	vehicleTypeRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/vehicle_types"
)

func RefreshVehicleTypes(ctx context.Context, cfg *config.Config, repository vehicleTypeRepo.Repository) error {

	vehicleCategories, err := vehicleTypeRequests.GetVehicleTypes(cfg.AutoriaKey)
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
