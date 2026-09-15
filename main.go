package main

import (
	"fmt"
	"strings"
)

// 1. Создаём пустые мапы.
var users = map[int]string{}
var balances = map[int]int{}
var productNames = map[int]string{}
var productPrices = map[int]int{}
var productStocks = map[int]int{}

var carts = map[int]map[int]int{}

// состав товара
var orderItems = map[int]map[int]int{}

// цены товаров на момент оформления
var orderItemPrices = map[int]map[int]int{}

// хранит статус заказа
var orderStatuses = map[int]string{}

// хранит историю операций пользователя
var operationHistory = map[int][]string{}

// склад
var stocks = map[int]int{}

// владелец заказа
var orderOwners = map[int]int{}

// orderTotals := map[int]int{}

func main() {
	// 2. Добавляем  пользователей и начальные балансы.
	AddUser(users, 1, "Anna")
	AddUser(users, 2, "Anna")
	AddUser(users, 3, "Masha")
	AddUser(users, 4, "Anna")
	fmt.Println(users)

	// 3. Проверяем, что повторяющийся идентификатор не добавляется.
	fmt.Println("Повторное добавление:", AddUser(users, 2, "Pavel"))

	// 4. Ищем пользователя по идентификатору.
	if name, ok := GetUser(users, 3); ok {
		fmt.Printf("Get user id 3: %v\n", name)
	}

	// 5. Переименовываем одного пользователя.
	RenameUser(users, 2, "Pavel")
	fmt.Println(users)

	// 6. Ищем пользователей по части имени.
	fmt.Println(FindUsersByName(users, "Anna"))
	// 7. Пополняем баланс двух пользователей.
	TopUpBalance(users, balances, 1, 12550)
	TopUpBalance(users, balances, 2, 5000)
	fmt.Println(balances)

	// 8. Выводим балансы в рублях и копейках.
	amount, ok := GetBalance(users, balances, 1)
	fmt.Println(amount)
	if ok {
		fmt.Println(GetBalancePrint(amount))
	}
	// 9.
	DeleteUser(users, 3)
	fmt.Println(users)
	fmt.Println(balances)

}

func AddUser(users map[int]string, id int, name string) bool {
	if id <= 0 {
		return false
	}
	if _, exists := users[id]; exists {
		return false
	}
	users[id] = name
	AddBalances(id)
	return true
}
func AddBalances(id int) {
	balances[id] = 0
}
func GetUser(users map[int]string, id int) (string, bool) {
	val, exists := users[id]
	return val, exists
}
func RenameUser(users map[int]string, id int, newName string) bool {
	if _, exists := users[id]; !exists {
		return false
	}
	users[id] = newName
	return true
}
func DeleteUser(users map[int]string, id int) bool {
	if _, exists := users[id]; !exists {
		return false
	}
	delete(users, id)
	DeleteBalances(id)
	return true
}
func DeleteBalances(id int) {
	delete(balances, id)
}
func FindUsersByName(users map[int]string, query string) []int {
	var result []int
	for id, name := range users {
		if name == query {
			result = append(result, id)
		}
	}
	return result
}
func TopUpBalance(
	users map[int]string,
	balances map[int]int,
	userID int,
	amount int,
) bool {
	if _, exists := users[userID]; !exists {
		return false
	}
	if amount <= 0 {
		return false
	}
	balances[userID] += amount
	return true

}
func GetBalancePrint(in int) string {
	return fmt.Sprintf("%d руб. %02d коп.", in/100, in%100)
}
func GetBalance(
	users map[int]string,
	balances map[int]int,
	userID int,
) (int, bool) {
	if _, exists := users[userID]; !exists {
		return 0, false
	}
	return balances[userID], true
}

func AddProduct(
	names map[int]string, prices map[int]int, stocks map[int]int,
	id int, name string, price int, stock int,
) bool {
	if id <= 0 {
		return false
	}
	if _, exists := names[id]; exists {
		return false
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return false
	}
	if price <= 0 {
		return false
	}
	if stock < 0 { //количество товара на складе
		return false
	}

	names[id] = name
	prices[id] = price
	stocks[id] = stock

	return true

}

func GetProduct(
	names map[int]string, prices map[int]int, stocks map[int]int,
	productID int,
) (string, int, int, bool) {
	name, exists := names[productID]
	if !exists {
		return "", 0, 0, false
	}
	//Получение товара должно возвращать название, цену, остаток и признак существования.
	return name, prices[productID], stocks[productID], true

}

func UpdateProductStock(
	names map[int]string, stocks map[int]int,
	productID int, stock int,
) bool {
	if _, exists := names[productID]; !exists {
		return false
	}
	if stock < 0 {
		return false
	}
	stocks[productID] = stock
	return true
}

func SearchProducts(names map[int]string, query string) []int {
	result := []int{}
	if query == "" {
		return result
	}

	query = strings.ToLower(strings.TrimSpace(query))

	for id, name := range names {
		name = strings.ToLower(name)

		if strings.Contains(name, query) {
			result = append(result, id)
		}
	}
	return result

}

// carts[userID][productID] = количество
// ??? корзину создаем тоже сразу, при создании пользователя?
func AddToCart(
	users map[int]string, products map[int]string,
	carts map[int]map[int]int,
	userID int, productID int, quantity int,
) bool {
	if _, exists := users[userID]; !exists {
		return false
	}
	if _, exists := products[productID]; !exists {
		return false
	}
	if quantity <= 0 {
		return false
	}
	if _, exists := carts[userID]; !exists {
		carts[userID] = map[int]int{}
	}
	carts[userID][productID] += quantity

	return true
}

// устанавливает количество уже добавленного товара в корзине
func SetCartQuantity(
	carts map[int]map[int]int,
	userID int, productID int, quantity int,
) bool {
	// две проверки: есть такой пользователь? Есть такой товар?
	userCart, userExists := carts[userID]
	if !userExists {
		return false
	}
	if _, productExists := userCart[productID]; !productExists {
		return false
	}
	if quantity < 0 {
		return false
	}
	if quantity == 0 {
		delete(userCart, productID)
		return true
	}
	userCart[productID] = quantity // ??? установим новое
	return true
}

// func IsExistUserAndProduct(
//
//	carts map[int]map[int]int, userID int, productID int,
//
//	) bool {
//		userCart, userExists := carts[userID]
//		if !userExists {
//			return false
//		}
//		_, productExists := userCart[productID]
//		return productExists
//	}
//
// удалить товар из корзины пользователя
func RemoveFromCart(
	carts map[int]map[int]int, userID int, productID int,
) bool {
	userCart, userExists := carts[userID]
	if !userExists {
		return false
	}
	if _, productExists := userCart[productID]; !productExists {
		return false
	}
	delete(userCart, productID)
	return true
}
func ClearCart(carts map[int]map[int]int, userID int) bool {
	_, userExists := carts[userID]
	if !userExists {
		return false
	}

	carts[userID] = map[int]int{}
	return true
}
func CalculateCartTotal(
	carts map[int]map[int]int, prices map[int]int, userID int,
) (int, bool) {
	userCart, userExists := carts[userID]
	if !userExists {
		return 0, false
	}
	total := 0
	for productID, quantity := range userCart {
		price, productExists := prices[productID]
		if !productExists {
			return 0, false
		}

		total += price * quantity
	}

	return total, true
}
func Checkout(
	users map[int]string,
	balances map[int]int,
	productPrices map[int]int,
	productStocks map[int]int,
	carts map[int]map[int]int,
	orderOwners map[int]int,
	orderItems map[int]map[int]int,
	orderItemPrices map[int]map[int]int,
	orderTotals map[int]int,
	orderStatuses map[int]string,
	operationHistory map[int][]string,
	userID int,
) (int, bool) {
	// 1. Проверяем пользователя.
	if _, userExists := users[userID]; !userExists {
		return 0, false
	}
	// 2. Проверяем корзину: существует, не пустая
	userCart, cartExists := carts[userID]
	if !cartExists || len(userCart) == 0 {
		return 0, false
	}
	// 3–5. Проверяем товары, остатки и  рассчитываем сумму.
	total := 0
	for productID, quantity := range userCart {
		if quantity <= 0 {
			return 0, false
		}

		price, productExists := productPrices[productID]
		if !productExists {
			return 0, false
		}

		stock, stockExists := productStocks[productID]
		if !stockExists || quantity > stock {
			return 0, false
		}

		total += price * quantity
	}
	// 6. Проверяем баланс.
	balance, balanceExists := balances[userID]
	if !balanceExists || balance < total {
		return 0, false
	}
	//---проверки закончились, начинаем обрабатывать заказ---------

	// 7. Создаём новый идентификатор заказа.
	orderID := 1
	for existingOrderID := range orderOwners {
		if existingOrderID >= orderID {
			orderID = existingOrderID + 1
		}
	}
	// 8–9. Копируем товары и сохраняем цены.
	items := map[int]int{}
	itemPrices := map[int]int{}
	//Заказ должен хранить отдельную копию из корзины ??? так?
	for productID, quantity := range userCart {
		items[productID] = quantity
		itemPrices[productID] = productPrices[productID]
	}

	orderItems[orderID] = items
	orderItemPrices[orderID] = itemPrices

	// 10. Списываем деньги.
	balances[userID] -= total

	// 11. Уменьшаем остатки на складе.
	for productID, quantity := range userCart {
		productStocks[productID] -= quantity
	}

	// 12. Сохранить владельца, сумму и статус paid
	orderOwners[orderID] = userID
	orderTotals[orderID] = total
	orderStatuses[orderID] = "paid"

	// 13. Очищаем корзину.
	carts[userID] = map[int]int{}

	// 14. Добавляем запись в историю.
	operationHistory[userID] = append(
		operationHistory[userID],
		fmt.Sprintf("order %d paid: %d", orderID, total),
	)

	return orderID, true
}

func CancelOrder(
	balances map[int]int, productStocks map[int]int,
	orderOwners map[int]int, orderItems map[int]map[int]int,
	orderTotals map[int]int, orderStatuses map[int]string,
	operationHistory map[int][]string,
	userID int, orderID int,
) bool {
	//  1,2. Проверяем существование заказа и пользователь его владелец.
	owner, orderExists := orderOwners[orderID]
	if !orderExists || owner != userID {
		return false
	}
	// 3. Проверить, что заказ имеет статус paid(оплаченный заказ)
	// уже отмененный заказ "cancelled" нельзя отменить повторно.
	// if orderStatuses[orderID] == "cancelled" {
	// 	return false
	// }
	if orderStatuses[orderID] != "paid" {
		return false
	}
	// 4. Возвращаем товары на склад.
	for productID, quantity := range orderItems[orderID] {
		productStocks[productID] += quantity
	}
	// 5. Возвращаем пользователю полную стоимость заказа.
	total := orderTotals[orderID]
	balances[userID] += total
	// 6. Изменить статус заказа на cancelled.
	orderStatuses[orderID] = "cancelled"

	// 7. Добавляем запись в историю операций.
	operationHistory[userID] = append(
		operationHistory[userID],
		fmt.Sprintf("order %d cancelled: %d", orderID, total),
	)
	return true
}

func GetUserOrders(
	orderOwners map[int]int, userID int,
) []int {
	orders := make([]int, 0)
	//orderOwners это map[order]owner
	for order, owner := range orderOwners {
		if owner == userID {
			orders = append(orders, order)
		}
	}
	return orders
}

func GetOrderItems(
	orderItems map[int]map[int]int, orderID int,
) (map[int]int, bool) {
	items, exists := orderItems[orderID]
	if !exists {
		return nil, false
	}
	itemsCopy := make(map[int]int, len(items))
	for productID, quantity := range items {
		itemsCopy[productID] = quantity
	}
	return itemsCopy, true
}

func GetUserHistory(
	history map[int][]string, userID int,
) []string {
	userHistory, exists := history[userID]
	if !exists {
		return nil
	}
	historyCopy := make([]string, len(userHistory))
	copy(historyCopy, userHistory) // копируем строки
	return historyCopy
}
