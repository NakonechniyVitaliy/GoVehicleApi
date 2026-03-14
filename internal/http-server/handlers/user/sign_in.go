package user

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/user"
	resp "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/user"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response resp.Response
	tokenJWT *string
}

func SignIn(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.user.sign_in"))

		var req dto.SignInRequest
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

		tokenJWT, err := srv.SignIn(r.Context(), req.SignData)

		if errors.Is(err, service.ErrUserNotFound) {
			resp.RenderError(w, r, http.StatusNotFound, service.ErrUserNotFound.Error())
			return
		}
		if err != nil {
			resp.RenderError(w, r, http.StatusInternalServerError, service.ErrGetUser.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response: resp.OK(),
			tokenJWT: tokenJWT,
		})
	}
}
