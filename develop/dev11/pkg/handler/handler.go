package handler

import "rest/pkg/repository"

type Handler struct {
	repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{repository: r}
}
