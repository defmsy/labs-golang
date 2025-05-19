package service

type FruitService struct{}

func NewFruitService() *FruitService {
	return &FruitService{}
}

func (s *FruitService) List() ([]string, error) {
	return []string{"apple", "banana", "orange"}, nil
}
