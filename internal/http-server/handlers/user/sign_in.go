package user

import (
	"errors"
	"log/slog"
	"net/http"

	dtoErrors "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/_errors"
	dto "github.com/NakonechniyVitalii/GoVehicleApi/internal/http-server/dto/user"
	response "github.com/NakonechniyVitalii/GoVehicleApi/internal/lib/api/response"
	jwtService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/jwt"
	userService "github.com/NakonechniyVitalii/GoVehicleApi/internal/services/user"
	"github.com/go-chi/render"
)

type GetResponse struct {
	Response response.Response
	TokenJWT string
}

// SignIn godoc
// @Summary      Вхід користувача
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body  body      SignInPayload  true  "Логін та пароль"
// @Success      200   {object}  GetResponse
// @Failure      400   {object}  response.Response
// @Failure      401   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /user/sign-in [post]
func SignIn(log *slog.Logger, uService *userService.Service, jwtService *jwtService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log = log.With(slog.String("op", "handlers.user.sign_in"))

		var req dto.SignInDTO
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

		tokenJWT, err := uService.SignIn(r.Context(), req, jwtService)
		if errors.Is(err, userService.ErrIncorrectLoginOrPass) {
			response.RenderError(w, r, http.StatusUnauthorized, userService.ErrIncorrectLoginOrPass.Error())
			return
		}
		if err != nil {
			response.RenderError(w, r, http.StatusInternalServerError, userService.ErrGetUser.Error())
			return
		}

		render.JSON(w, r, GetResponse{
			Response: response.OK(),
			TokenJWT: tokenJWT,
		})
	}
}
