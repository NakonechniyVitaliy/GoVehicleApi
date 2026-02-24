package category

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	categoryRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/category"
	categoryRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/categories"
)

func RefreshCategories(ctx context.Context, cfg *config.Config, repository categoryRepo.RepositoryInterface) error {

	categories, err := categoryRequests.GetCategories(cfg.AutoriaKey)
	if err != nil {
		return err
	}

	for _, oneCategory := range categories {
		err = repository.InsertOrUpdate(ctx, oneCategory)
		if err != nil {
			return err
		}
	}

	return nil
}
