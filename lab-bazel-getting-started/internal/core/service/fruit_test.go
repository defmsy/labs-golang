package service_test

import (
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"

	"github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/port/mock"
	"github.com/defmsy/labs-golang/lab-bazel-getting-started/internal/core/service"
)

func TestFruitService_ListFruits(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFruitRepository := mock.NewMockFruitRepository(ctrl)

	mockFruitRepository.EXPECT().ListFruits().Return([]string{"apple", "banana", "orange"}, nil)

	fruitService := service.NewFruitService(mockFruitRepository)
	fruits, err := fruitService.List()

	if err != nil {
		t.Errorf("Failed to list fruits: %s", err.Error())
	}

	if !reflect.DeepEqual(fruits, []string{"apple", "banana", "orange"}) {
		t.Errorf("Not implemented")
	}
}
