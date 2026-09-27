package main

import "fmt"

func Checkout(
	users map[int64]string,
	balances map[int64]int64,
	productPrices map[int64]int64,
	productStocks map[int64]int64,
	carts map[int64]map[int64]int64,
	orderOwners map[int64]int64,
	orderItems map[int64]map[int64]int64,
	orderItemPrices map[int64]map[int64]int64,
	orderTotals map[int64]int64,
	orderStatuses map[int64]string,
	operationHistory map[int64][]string,
	userID int64,
) (int64, bool) {
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
	var total int64
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
	orderID := int64(1)
	for existingOrderID := range orderOwners {
		if existingOrderID >= orderID {
			orderID = existingOrderID + 1
		}
	}
	// 8–9. Копируем товары и сохраняем цены.
	items := map[int64]int64{}
	itemPrices := map[int64]int64{}
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
	ClearCart(carts, userID)

	// 14. Добавляем запись в историю.
	operationHistory[userID] = append(
		operationHistory[userID],
		fmt.Sprintf("order %d paid: %d", orderID, total),
	)

	return orderID, true
}

func CancelOrder(
	balances map[int64]int64, productStocks map[int64]int64,
	orderOwners map[int64]int64, orderItems map[int64]map[int64]int64,
	orderTotals map[int64]int64, orderStatuses map[int64]string,
	operationHistory map[int64][]string,
	userID int64, orderID int64,
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
	orderOwners map[int64]int64, userID int64,
) []int64 {
	orders := make([]int64, 0)
	//orderOwners это map[order]owner
	for order, owner := range orderOwners {
		if owner == userID {
			orders = append(orders, order)
		}
	}
	return orders
}

func GetOrderItems(
	orderItems map[int64]map[int64]int64, orderID int64,
) (map[int64]int64, bool) {
	items, exists := orderItems[orderID]
	if !exists {
		return nil, false
	}
	itemsCopy := make(map[int64]int64, len(items))
	for productID, quantity := range items {
		itemsCopy[productID] = quantity
	}
	return itemsCopy, true
}
