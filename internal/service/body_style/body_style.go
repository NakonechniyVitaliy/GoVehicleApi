package vehicle_category

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	bodyStyleRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/body_style"
	bodyStyleRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/body_styles"
)

func RefreshBodyStyles(ctx context.Context, cfg *config.Config, repository bodyStyleRepo.RepositoryInterface) error {

	vehicleCategories, err := bodyStyleRequests.GetBodyStyles(cfg.AutoriaKey)
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
