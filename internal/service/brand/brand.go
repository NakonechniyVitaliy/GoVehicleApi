package brand

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	brandRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/brands"
)

func RefreshBrands(ctx context.Context, cfg *config.Config, repository brand.RepositoryInterface) error {

	brands, err := brandRequests.GetBrands(cfg.AutoriaKey)
	if err != nil {
		return err
	}

	for _, oneBrand := range brands {
		err = repository.InsertOrUpdate(ctx, oneBrand)
		if err != nil {
			return err
		}
	}

	return nil
}
