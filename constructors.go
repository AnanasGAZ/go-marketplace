package main

import "strings"

func NewUser(
	id int, name string, balance int,
) (User, bool) {
	if id <= 0 {
		return User{}, false
	}
	// 	Имя пользователя после удаления пробелов по краям
	// не должны быть пустыми
	name = strings.TrimSpace(name)
	if name == "" {
		return User{}, false
	}
	//Баланс и остаток товара не могут быть отрицательными.
	if balance < 0 {
		return User{}, false
	}
	return User{
		ID:      id,
		Name:    name,
		Balance: balance,
	}, true
}

func NewProduct(
	id int, name string, price int, stock int,
) (Product, bool) {
	//название товара после удаления пробелов
	// по краям не должны быть пустыми.
	if id <= 0 {
		return Product{}, false
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return Product{}, false
	}
	//Цена товара должна быть больше нуля
	if price <= 0 {
		return Product{}, false
	}
	// Количество товара в OrderItem должно быть больше нуля
	if stock <= 0 {
		return Product{}, false
	}
	return Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}, true
}

func NewCart(userID int) (Cart, bool) {
	if userID <= 0 {
		return Cart{}, false
	}

	return Cart{
		UserID: userID,
		Items:  make(map[int]int),
	}, true
}

func NewOrderItem(product Product, quantity int) (OrderItem, bool) {
	if product.ID <= 0 {
		return OrderItem{}, false
	}
	productName := strings.TrimSpace(product.Name)
	if productName == "" {
		return OrderItem{}, false
	}
	if product.Price <= 0 {
		return OrderItem{}, false
	}
	if quantity <= 0 {
		return OrderItem{}, false
	}
	return OrderItem{
		ProductID:   product.ID,
		ProductName: productName,
		Price:       product.Price,
		Quantity:    quantity,
	}, true
}

func NewOrder(
	id int, userID int, items []OrderItem,
) (Order, bool) {
	if id <= 0 {
		return Order{}, false
	}
	if userID <= 0 {
		return Order{}, false
	}
	if len(items) == 0 {
		return Order{}, false
	}
	return Order{
		ID:     id,
		UserID: userID,
		Items:  items,
		Total:  CalculateOrderTotal(items),
		Status: "peid",
	}, true
}

func CalculateOrderTotal(items []OrderItem) int {
	total := 0
	for _, item := range items {
		total += item.Price * item.Quantity
	}
	return total
}

func CopyOrderItems(items []OrderItem) []OrderItem {
	copiedItems := make([]OrderItem, len(items))
	copy(copiedItems, items) //все поля значения, поэтому можно так
	return copiedItems
}
