package handler

import (
	"github.com/mayron-dev/letsgoadmin/config"
)

type Handler struct {
	Logger *config.Logger
}

func NewHandler(logger *config.Logger) *Handler {
	handler := &Handler{
		Logger: logger,
	}
	return handler
}
