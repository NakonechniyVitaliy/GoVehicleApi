package vehicle_category

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	vehicleCategoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/vehicle_category"
	vehicleCategoryRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/vehicle_categories"
)

func RefreshVehicleCategories(ctx context.Context, cfg *config.Config, repository vehicleCategoryRepo.Repository) error {

	vehicleCategories, err := vehicleCategoryRequests.GetVehicleCategories(cfg.AutoriaKey)
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
