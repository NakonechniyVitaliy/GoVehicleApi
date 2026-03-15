package user

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/user"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	jwtService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/jwt"
	userService "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/user"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response resp.Response
	TokenJWT string
}

func SignIn(log *slog.Logger, uService *userService.Service, jwtService *jwtService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.user.sign_in"))

		var req dto.SignInDTO
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(dtoErrors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, dtoErrors.InvalidJSONorWrongFieldType)
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			resp.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		tokenJWT, err := uService.SignIn(r.Context(), req, jwtService)
		if errors.Is(err, userService.ErrIncorrectLoginOrPass) {
			resp.RenderError(w, r, http.StatusUnauthorized, userService.ErrIncorrectLoginOrPass.Error())
			return
		}
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, userService.ErrGetUser.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response: resp.OK(),
			TokenJWT: tokenJWT,
		})
	}
}
