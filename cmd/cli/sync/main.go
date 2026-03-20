package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/app"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/config"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/logger"
	bsRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/body_styles"
	bRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/brands"
	cRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/categories"
	dRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/driver_types"
	gRequests "github.com/NakonechniyVitaliy/GoVehicleApi/internal/requests/autoria/gearboxes"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	application, err := app.New(log, cfg)
	if err != nil {
		log.Error("failed to initialize app", slog.Any("err", err))
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log.Info("sync started")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		bStyles, err := bsRequests.GetBodyStyles(cfg.AutoriaKey)
		if err != nil {
			return err
		}
		return application.Services.BodyStyle.InsertOrUpdate(ctx, bStyles)
	})

	g.Go(func() error {
		brands, err := bRequests.GetBrands(cfg.AutoriaKey)
		if err != nil {
			return err
		}
		return application.Services.Brand.InsertOrUpdate(ctx, brands)
	})

	g.Go(func() error {
		categories, err := cRequests.GetCategories(cfg.AutoriaKey)
		if err != nil {
			return err
		}
		return application.Services.Category.InsertOrUpdate(ctx, categories)
	})

	g.Go(func() error {
		dTypes, err := dRequests.GetDriverTypes(cfg.AutoriaKey)
		if err != nil {
			return err
		}
		return application.Services.DriverType.InsertOrUpdate(ctx, dTypes)
	})

	g.Go(func() error {
		gearboxes, err := gRequests.GetGearboxes(cfg.AutoriaKey)
		if err != nil {
			return err
		}
		return application.Services.Gearbox.InsertOrUpdate(ctx, gearboxes)
	})

	if err := g.Wait(); err != nil {
		log.Error("sync failed", slog.Any("err", err))
		os.Exit(1)
	}

	log.Info("sync finished")
}
