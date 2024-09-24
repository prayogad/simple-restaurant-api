package web

type FoodCreateRequest struct {
	Name  string  `validate:"required,max=100,min=1" json:"name"`
	Price float32 `validate:"required,min=1" json:"price"`
	Stock int     `validate:"required,min=1" json:"stock"`
}

type FoodUpdateRequest struct {
	Id    int     `validate:"required" json:"id"`
	Name  string  `validate:"max=100" json:"name"`
	Price float32 `validate:"max=100" json:"price"`
	Stock int     `validate:"max=100" json:"stock"`
}

type FoodResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}
