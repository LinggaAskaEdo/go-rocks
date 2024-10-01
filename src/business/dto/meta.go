package dto

import apperr "github.com/linggaaskaedo/go-rocks/stdlib/errors"

type Meta struct {
	Path       string           `json:"path" extensions:"x-order=0"`
	StatusCode int              `json:"status_code" extensions:"x-order=1"`
	Status     string           `json:"status" extensions:"x-order=2"`
	Message    string           `json:"message" extensions:"x-order=3"`
	Error      *apperr.AppError `json:"error,omitempty" swaggertype:"primitive,object" extensions:"x-order=4"`
	Timestamp  string           `json:"timestamp" extensions:"x-order=5"`
}
