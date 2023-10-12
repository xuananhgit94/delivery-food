package retaurantbiz

import (
	"context"
	"errors"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"testing"
)

type mokeCreateStore struct{}

func (mokeCreateStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Anh" {
		return common.ErrDB(errors.New("something went wrong in DB"))
	}
	data.Id = 200
	return nil
}

func TestNewCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurantBiz(mokeCreateStore{})

	dataTest := restaurantmodel.RestaurantCreate{Name: ""}
	err := biz.CreateRestaurant(context.Background(), &dataTest)
	if err == nil || err.Error() != "name cannot is empty" {
		t.Errorf("Failed")
		return
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Anh"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)
	if err == nil {
		t.Errorf("Failed")
		return
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Xuan"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)
	if err != nil {
		t.Errorf("Failed")
		return
	}
	//t.Logf("TestNewCreateRestaurantBiz Passed")
}
