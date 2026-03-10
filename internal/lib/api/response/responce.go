package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func RenderError(w http.ResponseWriter, r *http.Request, httpStatus int, msg string) {
	render.Status(r, httpStatus)
	render.JSON(w, r, Error(msg))
}
