package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func NewUser(
	id int64, name string,
	// balance int,
) (User, bool) {
	if id <= 0 {
		return User{}, false
	}
	// 	Имя пользователя после удаления пробелов по краям
	// не должны быть пустыми
	name, ok := normalizeName(name)
	if !ok {
		return User{}, false
	}

	// //Баланс и остаток товара не могут быть отрицательными.
	// if balance < 0 {
	// 	return User{}, false
	// }
	return User{
		ID:   id,
		Name: name,
		// Balance: balance,
	}, true
}

func normalizeName(name string) (string, bool) {
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("Пустое имя")
		return "", false
	}

	if !utf8.ValidString(name) {
		fmt.Println("Не корректный текст")
		return "", false
	}

	length := utf8.RuneCountInString(name)
	if length < 1 || length > 200 {
		return "", false
	}

	return name, true
}

func NewWallet(userID int64) (Wallet, bool) {
	if userID <= 0 {
		return Wallet{}, false
	}
	return Wallet{
		UserID:  userID,
		Balance: 0,
	}, true
}
func NewProduct(
	id int64, name string, price int64, stock int64,
) (Product, bool) {
	//название товара после удаления пробелов
	// по краям не должны быть пустыми.
	if id <= 0 {
		return Product{}, false
	}
	name, ok := normalizeName(name)
	if !ok {
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

func NewCart(userID int64) (Cart, bool) {
	if userID <= 0 {
		return Cart{}, false
	}

	return Cart{
		UserID: userID,
		Items:  make(map[int64]int64),
	}, true
}

func NewOrderItem(product Product, quantity int64) (OrderItem, bool) {
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
	id int64, userID int64, items []OrderItem,
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
		//скопировать слайс, чтобы не менялся
		// список в заказе, если изменится список внешний
		Items:  CopyOrderItems(items),
		Total:  CalculateOrderTotal(items),
		Status: "peid",
	}, true
}

func CalculateOrderTotal(items []OrderItem) int64 {
	var total int64 //нулевое значение автоматом
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
