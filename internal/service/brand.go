package service

import (
	"net/http"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/brand"
	brandRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/brands"
)

func RefreshBrand(cfg *config.Config, repository brand.Repository, r *http.Request) error {

	brands, err := brandRequests.GetBrands(cfg.AutoriaKey)
	if err != nil {
		return err
	}

	for _, oneBrand := range brands {
		err = repository.InsertOrUpdate(r.Context(), oneBrand)
		if err != nil {
			return err
		}
	}

	return nil

}
