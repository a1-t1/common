package httputils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
	StatusBadRequest   = 400
	StatusNotFound     = 404
	StatusUnauthorized = 401
	StatusForbidden    = 403
	StatusInternal     = 500
	StatusOK           = 200
)

type Router interface {
	SetupRoutes() *chi.Mux
	SetupPublicRoutes() *chi.Mux
}

func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ParseStringPTRParamFromQuery(r *http.Request, param string) *string {
	val := r.URL.Query().Get(param)
	if val == "" || val == "null" || val == "undefined" {
		return nil
	}
	return &val
}

func ParseBoolPTRParamFromQuery(r *http.Request, param string) *bool {
	val := r.URL.Query().Get(param)
	if val == "" || val == "null" || val == "undefined" {
		return nil
	}

	parsedVal, err := strconv.ParseBool(val)
	if err != nil {
		return nil
	}
	return &parsedVal
}

func ErrorBadRequest(w http.ResponseWriter, err error) {
	write(w, StatusBadRequest, map[string]string{"error": err.Error()})
}

func ErrorUnauthorized(w http.ResponseWriter, err error) {
	write(w, StatusUnauthorized, map[string]string{"error": err.Error()})
}

func ErrorNotFound(w http.ResponseWriter, err error) {
	write(w, StatusNotFound, map[string]string{"error": err.Error()})
}

func ErrorInternal(w http.ResponseWriter, err error) {
	write(w, StatusInternal, map[string]string{"error": err.Error()})
}

func Ok(w http.ResponseWriter, data interface{}) {
	write(w, StatusOK, data)
}

func ErrorForbidden(w http.ResponseWriter, err error) {
	write(w, StatusForbidden, map[string]string{"error": err.Error()})
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func write(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ParseInt64Param(r *http.Request, key string) (int64, error) {
	p := chi.URLParam(r, key)
	return strconv.ParseInt(p, 10, 64)
}

func ParseInt64PTRParamFromQuery(r *http.Request, key string) *int64 {
	p := r.URL.Query().Get(key)
	if p == "" || p == "null" || p == "undefined" {
		return nil
	}
	val, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return nil
	}
	return &val
}

func ParseIntParam(r *http.Request, key string) (int, error) {
	p := chi.URLParam(r, key)
	return strconv.Atoi(p)
}
