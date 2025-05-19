package service

import "github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/port"

type FruitService struct {
	repo port.FruitRepository
}

func NewFruitService(
	repo port.FruitRepository,
) *FruitService {
	return &FruitService{
		repo,
	}
}

func (s *FruitService) List() ([]string, error) {
	return s.repo.ListFruits()
}
