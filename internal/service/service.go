package service

import (
	"Kokos/internal/repo"
)

type Service struct {
}

func NewService(repo *repo.Repository) *Service {
	return &Service{}
}
