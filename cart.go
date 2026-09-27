package main

// carts[userID][productID] = количество
func AddToCart(
	users map[int64]string, products map[int64]string,
	carts map[int64]map[int64]int64,
	userID int64, productID int64, quantity int64,
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
		carts[userID] = map[int64]int64{}
	}
	carts[userID][productID] += quantity

	return true
}

//???исоздала другой вариант функции AddToCart,
//  которая принимает Cart и Product
// вместо map пользователей и продуктов.
func AddToCartV3(
	cart Cart,
	product Product,
	quantity int64,
) bool {
	if cart.UserID <= 0 {
		return false
	}

	if product.ID <= 0 {
		return false
	}

	if quantity <= 0 {
		return false
	}

	cart.Items[product.ID] += quantity

	return true
}

// устанавливает количество уже добавленного товара в корзине
func SetCartQuantity(
	carts map[int64]map[int64]int64,
	userID int64, productID int64, quantity int64,
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
func ClearCart(carts map[int64]map[int64]int64, userID int64) bool {
	_, userExists := carts[userID]
	if !userExists {
		return false
	}

	carts[userID] = map[int64]int64{}
	return true
}
func CalculateCartTotal(
	carts map[int64]map[int64]int64, prices map[int64]int64, userID int64,
) (int64, bool) {
	userCart, userExists := carts[userID]
	if !userExists {
		return 0, false
	}
	var total int64
	for productID, quantity := range userCart {
		price, productExists := prices[productID]
		if !productExists {
			return 0, false
		}

		total += price * quantity
	}

	return total, true
}
