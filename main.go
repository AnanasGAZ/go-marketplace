package main

import (
	"fmt"
)

func main() {
	// TestCreateValidUser()
	// TestCreateInvalidUser()
	// TestCreateValidProduct()
	// TestAddProductToCart()
	// TestCreateOrderItems()
	// TestCreateOrder()
	// TestChangeOriginalItems()
	TestTopUp()

	// TestV1()
	// TestV2()

}
func TestV2() {
	fmt.Println("=== 1. Добавление пользователей и товаров ===")

	AddUser(users, 1, "Анна")
	AddUser(users, 2, "Маша")

	AddProduct(
		productNames,
		productPrices,
		productStocks,
		101,
		"Ноутбук",
		100000,
		5,
	)
	AddProduct(
		productNames,
		productPrices,
		productStocks,
		102,
		"Мышь",
		1500,
		10,
	)
	AddProduct(
		productNames,
		productPrices,
		productStocks,
		103,
		"Телефон",
		100000,
		2,
	)
	fmt.Println("Пользователи:", users)
	fmt.Println("Товары:", productNames)
	fmt.Println("Цены:", productPrices)
	fmt.Println("Остатки:", productStocks)

	fmt.Println("\n=== 2. Поиск товара по части названия ===")

	searchResult := SearchProducts(productNames, "  ")
	fmt.Println("Результат поиска по запросу «ноут»:", searchResult)

	fmt.Println("\n=== 3. Добавление товара в корзину несколько раз ===")

	AddToCart(users, productNames, carts, 1, 101, 1)
	AddToCart(users, productNames, carts, 1, 101, 2)

	fmt.Println("Корзина после повторного добавления:", carts[1])

	fmt.Println("\n=== 4. Изменение количества и расчёт стоимости ===")

	SetCartQuantity(carts, 1, 101, 2)

	cartTotal, totalOK := CalculateCartTotal(
		carts,
		productPrices,
		1,
	)

	fmt.Println("Количество товара в корзине:", carts[1])
	fmt.Println("Стоимость корзины:", cartTotal)
	fmt.Println("Расчёт выполнен:", totalOK)
	fmt.Println("\n=== 5. Пополнение баланса и оформление заказа ===")

	TopUpBalance(users, balances, 1, 500000) // если 150000 то потом заказ не оформится, так как мало

	orderID, checkoutOK := Checkout(
		users,
		balances,
		productPrices,
		productStocks,
		carts,
		orderOwners,
		orderItems,
		orderItemPrices,
		orderTotals,
		orderStatuses,
		operationHistory,
		1,
	)

	fmt.Println("Номер заказа:", orderID)
	fmt.Println("Заказ оформлен:", checkoutOK)
	fmt.Println("\n=== 6. Проверка результата оформления заказа ===")

	balance, _ := GetBalance(users, balances, 1)
	userOrders := GetUserOrders(orderOwners, 1)
	items, itemsOK := GetOrderItems(orderItems, orderID)

	fmt.Println("Новый баланс:", balance)
	fmt.Println("Остатки на складе:", productStocks)
	fmt.Println("Заказы пользователя:", userOrders)
	fmt.Println("Товары в заказе:", items)
	fmt.Println("Товары заказа найдены:", itemsOK)
	fmt.Println("Корзина после оформления:", carts[1])
	fmt.Println("Корзина пустая:", len(carts[1]) == 0)

	fmt.Println("\n=== 7. Оформление заказа при недостаточном балансе ===")

	AddToCart(users, productNames, carts, 1, 103, 5)

	poorBalanceOrderID, poorBalanceCheckoutOK := Checkout(
		users,
		balances,
		productPrices,
		productStocks,
		carts,
		orderOwners,
		orderItems,
		orderItemPrices,
		orderTotals,
		orderStatuses,
		operationHistory,
		1,
	)

	fmt.Println("Заказ при недостаточном балансе:", poorBalanceOrderID)
	fmt.Println("Оформление выполнено:", poorBalanceCheckoutOK)

	// Чистим руками для теста (После неудачного оформления корзина не очищается.
	ClearCart(carts, 1)
	fmt.Println("\n=== 8. Покупка товара в количестве больше складского остатка ===")

	// После первого заказа ноутбуков осталось 3.
	AddToCart(users, productNames, carts, 1, 101, 4)

	lowStockOrderID, lowStockCheckoutOK := Checkout(
		users,
		balances,
		productPrices,
		productStocks,
		carts,
		orderOwners,
		orderItems,
		orderItemPrices,
		orderTotals,
		orderStatuses,
		operationHistory,
		1,
	)

	fmt.Println("Заказ при нехватке товара:", lowStockOrderID)
	fmt.Println("Оформление выполнено:", lowStockCheckoutOK)
	//Чистим руками для теста
	ClearCart(carts, 1)

	fmt.Println("\n=== 9. Отмена оплаченного заказа ===")

	balanceBeforeCancel := balances[1]
	stockBeforeCancel := productStocks[101]

	cancelOK := CancelOrder(
		balances,
		productStocks,
		orderOwners,
		orderItems,
		orderTotals,
		orderStatuses,
		operationHistory,
		1,
		orderID,
	)

	fmt.Println("Заказ отменён:", cancelOK)
	fmt.Println("Баланс до отмены:", balanceBeforeCancel)
	fmt.Println("Баланс после отмены:", balances[1])
	fmt.Println("Остаток товара до отмены:", stockBeforeCancel)
	fmt.Println("Остаток товара после отмены:", productStocks[101])
	fmt.Println("Статус заказа:", orderStatuses[orderID])
	fmt.Println("История операций:", operationHistory[1])

	fmt.Println("\n=== 10. Повторная отмена того же заказа ===")

	balanceAfterCancel := balances[1]
	stockAfterCancel := productStocks[101]

	repeatCancelOK := CancelOrder(
		balances,
		productStocks,
		orderOwners,
		orderItems,
		orderTotals,
		orderStatuses,
		operationHistory,
		1,
		orderID,
	)

	fmt.Println("Повторная отмена выполнена:", repeatCancelOK)
	fmt.Println("Баланс не изменился:", balances[1] == balanceAfterCancel)
	fmt.Println("Остаток товара не изменился:", productStocks[101] == stockAfterCancel)

}
func TestV1() {
	// 2. Добавляем  пользователей и начальные балансы.
	AddUser(users, 1, "Anna")
	AddUser(users, 2, "Anna")
	AddUser(users, 3, "Masha")
	AddUser(users, 4, "Anna")
	AddUser(users, 5, "anna")
	AddUser(users, 6, "aNNA")
	AddUser(users, 7, "aNNA Иванова")

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
	fmt.Println("Много Анн:", FindUsersByName(users, "Anna"))
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
