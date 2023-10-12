package restaurantmodel

import (
	"errors"
	"testing"
)

type testData struct {
	Input  RestaurantCreate
	Expect error
}

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTests := []testData{
		{Input: RestaurantCreate{Name: ""}, Expect: ErrNameIsEmpty},
		{Input: RestaurantCreate{Name: "Xuan anh"}, Expect: nil},
	}

	for _, item := range dataTests {
		err := item.Input.Validate()
		if errors.Is(err, item.Expect) {
			t.Logf("Validate restaurant pass")
			continue
		}
		t.Errorf("Validate restaurant. Input: %v, Expect: %v, Output: %v", item.Input, item.Expect, err)
	}
}
