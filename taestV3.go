package main

import "fmt"

//Создайте корректного пользователя и выведите все его поля
func TestCreateValidUser() {
	user, ok := NewUser(1, "  Анна  ", 50000)

	if !ok {
		fmt.Println("Пользователь не создан")
		return
	}

	fmt.Println("Пользователь создан:")
	fmt.Println("ID:", user.ID)
	fmt.Println("Имя:", user.Name)
	fmt.Println("Баланс:", user.Balance)
}

//Попробуйте создать пользователя
// с пустым именем и отрицательным балансом.
func TestCreateInvalidUser() {
	user, ok := NewUser(0, "", -1000)

	if !ok {
		fmt.Println("Пользователь не создан")
		return
	}

	fmt.Println("Пользователь создан:")
	fmt.Println("ID:", user.ID)
	fmt.Println("Имя:", user.Name)
	fmt.Println("Баланс:", user.Balance)
}

//Создайте товар с ценой 19900 и остатком 10.
func TestCreateValidProduct() {
	product, ok := NewProduct(101, "Ноутбук", 19900, 10)

	if !ok {
		fmt.Println("Товар не создан")
		return
	}

	fmt.Println("Товар создан:")
	fmt.Println("ID:", product.ID)
	fmt.Println("Название:", product.Name)
	fmt.Println("Цена:", product.Price)
	fmt.Println("Остаток:", product.Stock)
}

//Создайте пустую корзину и добавьте товар в Cart.Items.
func TestAddProductToCart() {
	cart, ok := NewCart(1)
	if !ok {
		fmt.Println("Корзина не создана")
		return
	}
	product, ok := NewProduct(101, "Ноутбук", 19900, 10)
	if !ok {
		fmt.Println("Товар не создан")
		return
	}

	AddToCartV3(cart, product, 2)

	fmt.Println("Корзина:", cart.Items)
}

// Создайте две позиции заказа из разных товаров.
func TestCreateOrderItems() {
	firstProduct, ok := NewProduct(
		101,
		"Ноутбук",
		19900,
		10,
	)
	if !ok {
		fmt.Println("Первый товар не создан")
		return
	}

	secondProduct, ok := NewProduct(
		102,
		"Мышь",
		1500,
		20,
	)
	if !ok {
		fmt.Println("Второй товар не создан")
		return
	}

	firstItem, ok := NewOrderItem(firstProduct, 2)
	if !ok {
		fmt.Println("Первая позиция заказа не создана")
		return
	}

	secondItem, ok := NewOrderItem(secondProduct, 1)
	if !ok {
		fmt.Println("Вторая позиция заказа не создана")
		return
	}

	fmt.Println("Первая позиция:", firstItem)
	fmt.Println("Вторая позиция:", secondItem)
}

//Создайте заказ и проверьте его итоговую стоимость и статус
func TestCreateOrder() {
	firstProduct, ok := NewProduct(
		101,
		"Ноутбук",
		19900,
		10,
	)
	if !ok {
		fmt.Println("Первый товар не создан")
		return
	}

	secondProduct, ok := NewProduct(
		102,
		"Мышь",
		1500,
		20,
	)
	if !ok {
		fmt.Println("Второй товар не создан")
		return
	}

	firstItem, ok := NewOrderItem(firstProduct, 2)
	if !ok {
		fmt.Println("Первая позиция не создана")
		return
	}

	secondItem, ok := NewOrderItem(secondProduct, 1)
	if !ok {
		fmt.Println("Вторая позиция не создана")
		return
	}

	items := []OrderItem{
		firstItem,
		secondItem,
	}

	order, ok := NewOrder(1, 1, items)
	if !ok {
		fmt.Println("Заказ не создан")
		return
	}

	expectedTotal := 19900*2 + 1500*1

	fmt.Println("Заказ создан:", ok)
	fmt.Println("Итоговая стоимость:", order.Total)
	fmt.Println("Ожидаемая стоимость:", expectedTotal)
	fmt.Println("Стоимость правильная:", order.Total == expectedTotal)

	fmt.Println("Статус заказа:", order.Status)
	fmt.Println("Статус правильный:", order.Status == "paid")
}

// Измените исходный слайс позиций после создания заказа
func TestChangeOriginalItems() {
	product, ok := NewProduct(
		101,
		"Ноутбук",
		19900,
		10,
	)
	if !ok {
		fmt.Println("Товар не создан")
		return
	}
	item, ok := NewOrderItem(product, 2)
	if !ok {
		fmt.Println("Позиция не создана")
		return
	}
	items := []OrderItem{item}
	order, ok := NewOrder(1, 1, items)
	if !ok {
		fmt.Println("Заказ не создан")
		return
	}
	fmt.Println("До изменения:")
	fmt.Println("Исходный слайс:", items)
	fmt.Println("Позиции заказа:", order.Items)

	items[0].Quantity = 10
	items[0].ProductName = "Изменённый товар"

	fmt.Println("После изменения исходного слайса:")
	fmt.Println("Исходный слайс:", items)
	fmt.Println("Позиции заказа:", order.Items)
	fmt.Println("Заказ защищён от изменения:",
		order.Items[0].Quantity == 2 &&
			order.Items[0].ProductName == "Ноутбук",
	)

}
