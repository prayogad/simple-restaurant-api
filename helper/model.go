package helper

import (
	"simple-restaurant-web/model/domain"
	"simple-restaurant-web/model/web"
)

func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
	}
}

func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}

func ToCustomerLoginResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:       customer.Id,
		Username: customer.Username,
		Token:    customer.Token,
	}
}

func ToFoodResponse(food domain.Food) web.FoodResponse {
	return web.FoodResponse{
		Id:    food.Id,
		Name:  food.Name,
		Price: food.Price,
		Stock: food.Stock,
	}
}

func ToFoodResponses(foods []domain.Food) []web.FoodResponse {
	var foodResponses []web.FoodResponse
	for _, food := range foods {
		foodResponses = append(foodResponses, ToFoodResponse(food))
	}

	return foodResponses
}

// func ToOrderDetailResponse(orderDetail domain.OrderDetail) web.OrderDetailResponse {
// 	return web.OrderDetailResponse{
// 		FoodName:  orderDetail.FoodName,
// 		FoodPrice: orderDetail.FoodPrice,
// 		Quantity:  orderDetail.Quantity,
// 	}
// }

func ToOrderResponse(order domain.Orders) web.OrderResponse {
	var orderDetailResponses []web.OrderDetailResponse
	for _, orderDetail := range order.OrderDetails {
		newOrderDetail := web.OrderDetailResponse{
			FoodName:  orderDetail.FoodName,
			FoodPrice: orderDetail.FoodPrice,
			Quantity:  orderDetail.Quantity,
		}
		orderDetailResponses = append(orderDetailResponses, newOrderDetail)
	}

	return web.OrderResponse{
		IdOrder:       order.Id,
		TotalQuantity: order.Quantity,
		TotalPrice:    float32(order.TotalPrice),
		OrderDetail:   orderDetailResponses,
	}
}
