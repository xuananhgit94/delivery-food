package restaurantmodel

type Filter struct {
	UserId int   `json:"owner_id,omitempty" form:"owner_id"`
	Status []int `json:"-"`
}
