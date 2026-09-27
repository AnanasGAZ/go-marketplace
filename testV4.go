package main

import "fmt"

func TestTopUp() {
	//Go сможет автоматически взять адрес результата???
	user := User{ID: 1, Name: "Alex"}
	w := Wallet{UserID: user.ID, Balance: 10000}
	w.TopUp(5000)
	fmt.Println(w.Balance) // 15000
}

func TestTopUpCopy() {
	user := User{ID: 1, Name: "Alex"}
	w := Wallet{UserID: user.ID, Balance: 10000}
	fmt.Println("Ожидаем что баланс 10000, не пополнится")
	TopUpCopy(w, 5000) // передали значение user и в нем значение Balance int,
	//  поэтому не изменилось
	fmt.Println(w.Balance)
	fmt.Println("Ожидаем что баланс 17000, пополнится")
	TopUpPointer(&w, 7000) // передали адрес
	fmt.Println(w.Balance)

}

// Проверка работы
func TestV4() {
	u1 := User{ID: 1, Name: "Alex"}
	w1 := Wallet{u1.ID, 10000}

	TestV4withdraw(u1, w1, 5000, 7000) //Balance: 10000
	TestV4withdrawTooMuch(u1, w1)      //, Balance: 10000
	TestV4reserveProduct(Product{ID: 1, Name: "Light sabr", Price: 100000, Stock: 0})
	TestV4quantityTooMuch(Product{ID: 1, Name: "Light sabr", Price: 100000, Stock: 0})

	u := User{ID: 1, Name: "Alex"}
	// w := Wallet{u.ID, 10000}
	p := Product{ID: 1, Name: "Light sabr", Price: 100000, Stock: 0}
	c := Cart{
		UserID: u.ID,
		Items:  make(map[int64]int64),
	}
	oi := []OrderItem{{
		ProductID:   p.ID,
		ProductName: p.Name,
		Price:       p.Price,
		Quantity:    1}}
	o := Order{ID: 1, UserID: u.ID,
		Items:  oi,
		Total:  CalculateOrderTotal(oi), // в конструкторе...
		Status: "paid"}
	TestV4cart(c, p, 1)
	TestV4cartZeroDel(c, p)
	TestV4ClearCart(c)
	TestV4MarkCancelled(o)

}

// Пополнить и списать деньги с баланса пользователя.
func TestV4withdraw(user User, wallet Wallet, topup int64, withdraw int64) {
	fmt.Println("=== v4 1. Пополнить и списать деньги с баланса пользователя===")
	fmt.Printf("1 Попробуем пополнить баланс на %v: ", topup)
	fmt.Println(wallet.TopUp(topup))
	fmt.Printf("Попробуем списать сумму %v: ", withdraw)
	fmt.Println(wallet.Withdraw(withdraw))
}

//  Попробуйте списать больше денег, чем есть на балансе.
func TestV4withdrawTooMuch(user User, wallet Wallet) {
	fmt.Println("=== v4 2. Попробуйте списать больше денег, чем есть на балансе.===")

	withdraw := int64(15000)
	fmt.Printf("2 Попробуйте списать сумму %v больше текущего баланса %v: ", withdraw, wallet.Balance)
	fmt.Println(wallet.Withdraw(withdraw))
}

//Добавьте товар на склад и зарезервируйте несколько единиц:
func TestV4reserveProduct(p Product) {
	fmt.Println("=== v4 3. Добавьте товар на склад и зарезервируйте несколько единиц===")

	fmt.Print("3  Добавьте товар на склад и зарезервируйте несколько единиц: ")
	p.AddStock(3)
	fmt.Println(p.Reserve(1))
}

//  Попробуйте зарезервировать больше товара, чем есть на складе.
func TestV4quantityTooMuch(p Product) {
	fmt.Println("=== v4 4. Попробуйте зарезервировать больше товара, чем есть на складе.===")

	quantityTooMuch := p.Stock + 1
	fmt.Printf("4 Попробуйте зарезервировать %v единиц, когда на складе осталось %v: ", quantityTooMuch, p.Stock)
	fmt.Println(p.Reserve(quantityTooMuch))
	fmt.Printf("Остаток на складе после попытки: %v\n", p.Stock)
}

//  Добавьте товар в корзину два раза и проверьте количество.
func TestV4cart(c Cart, p Product, quantity int64) {
	fmt.Println("=== v4 5. Добавьте товар в корзину два раза и проверьте количество ===")

	c.Add(p.ID, quantity)
	c.Add(p.ID, quantity)

	fmt.Printf("5 Добавьте товар в корзину два раза и проверьте количество: %v\n", c.Items[p.ID])
}

//  Измените количество на 0 и убедитесь, что товар удалился.
func TestV4cartZeroDel(c Cart, p Product) {
	fmt.Println("=== v4 6. Измените количество на 0 и убедитесь, что товар удалился. ===")

	fmt.Println("6 Если количество 0 , то удалить товар:")
	fmt.Println("было в корзине:")
	fmt.Println(c.Items)
	c.SetQuantity(p.ID, 0)
	fmt.Println("стало в корзине:")
	fmt.Println(c.Items)
}

//  Очистите корзину и проверьте, что ее мапа не равна nil.
func TestV4ClearCart(c Cart) {
	fmt.Println("=== v4 7. Очистите корзину и проверьте, что ее мапа не равна nil. ===")

	fmt.Println("Очистите корзину и проверьте, что ее мапа не равна nil")
	fmt.Printf("Кориза до очисткиЖ %v\n", c.Items)
	c.Clear()
	fmt.Printf("Кориза почле очисткиЖ %v\n", c.Items)
}

//  Отмените заказ и попробуйте отменить его повторно.
func TestV4MarkCancelled(o Order) {
	fmt.Println("=== v4 8. Отмените заказ и попробуйте отменить его повторно. ===")

	fmt.Println(" Отмените заказ и попробуйте отменить его повторно.")
	fmt.Printf("Отменяем заказ: %v\n", o.Status)
	fmt.Printf("Отменяем заказ: %v\n", o.MarkCancelled())
	fmt.Printf("2 Отменяем заказ: %v\n", o.Status)
	fmt.Printf("Отмена еще раз: %v\n", o.MarkCancelled())

}
