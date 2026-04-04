package user

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitaliy/GoVehicleApi/internal/http-server/dto/user"
	response "github.com/NakonechniyVitaliy/GoVehicleApi/internal/lib/api/response"
	service "github.com/NakonechniyVitaliy/GoVehicleApi/internal/services/user"
	"github.com/go-chi/render"
)

// SignUp godoc
// @Summary      Реєстрація користувача
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body  body      SignUpPayload  true  "Дані нового користувача"
// @Success      200   {object}  response.Response
// @Failure      400   {object}  response.Response
// @Failure      409   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /user/sign-up [post]
func SignUp(log *slog.Logger, srv *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.user.sign_up"))

		var req dto.SignUpDTO
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error(dtoErrors.InvalidJSONorWrongFieldType, slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, dtoErrors.InvalidJSONorWrongFieldType)
			return
		}

		err = req.Validate()
		if err != nil {
			log.Error("validation error", slog.String("error", err.Error()))
			response.RenderError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		err = srv.SignUp(r.Context(), req)

		if errors.Is(err, service.ErrUserExists) {
			response.RenderError(w, r, http.StatusConflict, service.ErrUserExists.Error())
			return
		}
		if errors.Is(err, service.ErrSaveUser) {
			response.RenderError(w, r, http.StatusInternalServerError, service.ErrSaveUser.Error())
			return
		}

		render.JSON(w, r, response.OK())
	}
}
