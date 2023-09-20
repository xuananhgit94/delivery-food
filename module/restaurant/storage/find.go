package restaurantstorage

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindRestaurantWidthCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var dataRestaurant restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&dataRestaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &dataRestaurant, nil
}
