package driver_type

import (
	"context"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	gearboxRepo "github.com/NakonechniyVitaliy/GoVehicleApi/internal/repository/gearbox"
	gearboxRequets "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/gearboxes"
)

func RefreshGearboxes(ctx context.Context, cfg *config.Config, repository gearboxRepo.RepositoryInterface) error {

	gearboxes, err := gearboxRequets.GetGearboxes(cfg.AutoriaKey)
	if err != nil {
		return err
	}

	for _, oneGearbox := range gearboxes {
		err = repository.InsertOrUpdate(ctx, oneGearbox)
		if err != nil {
			return err
		}
	}

	return nil
}
