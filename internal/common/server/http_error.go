package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dkhrunov/hexagonal-architecture/internal/common/errors"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpResponseWithError(slug, err, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorized(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpResponseWithError(slug, err, w, r, "Unauthorized", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpResponseWithError(slug, err, w, r, "Bad request", http.StatusBadRequest)
}

func NotFound(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpResponseWithError(slug, err, w, r, "Not found", http.StatusNotFound)
}

func RespondWithError(err error, w http.ResponseWriter, r *http.Request) {
	slugError, ok := err.(errors.SlugError)
	if !ok {
		InternalError("internal-server-error", err, w, r)
		return
	}

	switch slugError.ErrorType() {
	case errors.ErrorTypeAuthorization:
		Unauthorized(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeNotFound:
		NotFound(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

func httpResponseWithError(slug string, err error, w http.ResponseWriter, r *http.Request, msg string, status int) {
	log.Printf("error: %s, slug: %s, msg: %s", err, slug, msg)

	resp := ErrorResponse{
		Error: ErrorInfo{slug, err.Error(), status},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

type ErrorInfo struct {
	Message    string `json:"message"`
	Details    string `json:"details"`
	httpStatus int
}

type ErrorResponse struct {
	Error ErrorInfo `json:"error"`
}

func (e ErrorResponse) Render(w http.ResponseWriter, _ *http.Request) error {
	w.WriteHeader(e.Error.httpStatus)
	return nil
}
