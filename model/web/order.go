package web

type OrderDetailCreateRequest struct {
	FoodId   int `validate:"required,min=1" json:"food_id"`
	Quantity int `validate:"required,min=1" json:"quantity"`
}

type OrderCreateRequest struct {
	OrderDetails []OrderDetailCreateRequest `json:"order_details"`
}

type OrderDetailResponse struct {
	FoodName  string  `json:"name"`
	FoodPrice float32 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type OrderResponse struct {
	IdOrder       int     `json:"id"`
	TotalQuantity int     `json:"total_quantity"`
	TotalPrice    float32 `json:"total_price"`
	OrderDetail   []OrderDetailResponse
}
