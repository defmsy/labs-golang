package service_test

import (
	"reflect"
	"testing"

	"github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/service"
)

func TestFruitService_ListFruits(t *testing.T) {
	fruitService := service.NewFruitService()
	fruits, err := fruitService.List()

	if err != nil {
		t.Errorf("Failed to list fruits: %s", err.Error())
	}

	if !reflect.DeepEqual(fruits, []string{"apple", "banana", "orange"}) {
		t.Errorf("Not implemented")
	}
}
