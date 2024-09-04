package db

import (
	"time"

	"github.com/ars0915/glossika-exercise/entity"
)

func (s *AppRepo) RecommendProduct() ([]entity.Product, error) {
	time.Sleep(3 * time.Second)

	products := []entity.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
	}
	return products, nil
}
