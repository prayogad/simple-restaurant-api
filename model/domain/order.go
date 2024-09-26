package domain

type Orders struct {
	Id           int
	Quantity     int
	TotalPrice   float32
	IdCustomer   int
	OrderDetails []OrderDetail
}

type OrderDetail struct {
	OrderId   int
	FoodId    int
	Quantity  int
	FoodName  string
	FoodPrice float32
}
