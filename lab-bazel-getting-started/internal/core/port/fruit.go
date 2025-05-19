package port

type FruitRepository interface {
	ListFruits() ([]string, error)
}
