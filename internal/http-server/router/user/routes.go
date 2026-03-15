package user

import (
	"log/slog"

	userHandler "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/handlers/user"
	"github.com/NakonechniyVitaliy/GoVehicleApi/internal/services"
	"github.com/go-chi/chi/v5"
)

func SetupUserRoutes(
	router chi.Router,
	log *slog.Logger,
	services services.Container,

) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/sign-in", userHandler.SignIn(log, services.User, services.JWT))
		r.Post("/sign-up", userHandler.SignUp(log, services.User))
	})
}
