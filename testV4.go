package main

import "fmt"

func TestTopUp() {
	//Go сможет автоматически взять адрес результата???
	user := User{ID: 1, Name: "Alex", Balance: 10000}
	user.TopUp(5000)
	fmt.Println(user.Balance) // 15000
}

func TestTopUpCopy() {
	user := User{ID: 1, Name: "Alex", Balance: 10000}
	TopUpCopy(user, 5000) // передали значение user и в нем значение Balance int, поэтому не изменилось
	fmt.Println(user.Balance)
	TopUpPointer(&user, 7000) // передали адрес
	fmt.Println(user.Balance)

}

// Проверка работы
func TestV4() {

	TestV4withdraw(User{ID: 1, Name: "Alex", Balance: 10000}, 5000, 7000)
	TestV4withdrawTooMuch(User{ID: 1, Name: "Alex", Balance: 10000})
	TestV4reserveProduct(Product{ID: 1, Name: "Light sabr", Price: 100000, Stock: 0})
	TestV4quantityTooMuch(Product{ID: 1, Name: "Light sabr", Price: 100000, Stock: 0})
	u := User{ID: 1, Name: "Alex", Balance: 10000}
	p := Product{ID: 1, Name: "Light sabr", Price: 100000, Stock: 0}
	c := Cart{
		UserID: u.ID,
		Items:  make(map[int]int),
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
func TestV4withdraw(user User, topup int, withdraw int) {
	fmt.Printf("1 Попробуем пополнить баланс на %v: ", topup)
	fmt.Println(user.TopUp(topup))
	fmt.Printf("Попробуем списать сумму %v: ", withdraw)
	fmt.Println(user.Withdraw(withdraw))
}

//  Попробуйте списать больше денег, чем есть на балансе.
func TestV4withdrawTooMuch(user User) {
	withdraw := 15000
	fmt.Printf("2 Попробуйте списать сумму %v больше текущего баланса %v: ", withdraw, user.Balance)
	fmt.Println(user.Withdraw(withdraw))
}

//Добавьте товар на склад и зарезервируйте несколько единиц:
func TestV4reserveProduct(p Product) {
	fmt.Print("3  Добавьте товар на склад и зарезервируйте несколько единиц: ")
	p.AddStock(3)
	fmt.Println(p.Reserve(1))
}

//  Попробуйте зарезервировать больше товара, чем есть на складе.
func TestV4quantityTooMuch(p Product) {
	quantityTooMuch := p.Stock + 1
	fmt.Printf("4 Попробуйте зарезервировать %v единиц, когда на складе осталось %v: ", quantityTooMuch, p.Stock)
	fmt.Println(p.Reserve(quantityTooMuch))
	fmt.Printf("Остаток на складе после попытки: %v\n", p.Stock)
}

//  Добавьте товар в корзину два раза и проверьте количество.
func TestV4cart(c Cart, p Product, quantity int) {

	c.Add(p.ID, quantity)
	c.Add(p.ID, quantity)

	fmt.Printf("5 Добавьте товар в корзину два раза и проверьте количество: %v\n", c.Items[p.ID])
}

//  Измените количество на 0 и убедитесь, что товар удалился.
func TestV4cartZeroDel(c Cart, p Product) {
	fmt.Println("6 Если количество 0 , то удалить товар:")
	fmt.Println("было в корзине:")
	fmt.Println(c.Items)
	c.SetQuantity(p.ID, 0)
	fmt.Println("стало в корзине:")
	fmt.Println(c.Items)
}

//  Очистите корзину и проверьте, что ее мапа не равна nil.
func TestV4ClearCart(c Cart) {
	fmt.Println("Очистите корзину и проверьте, что ее мапа не равна nil")
	fmt.Printf("Кориза до очисткиЖ %v\n", c.Items)
	c.Clear()
	fmt.Printf("Кориза почле очисткиЖ %v\n", c.Items)
}

//  Отмените заказ и попробуйте отменить его повторно.
func TestV4MarkCancelled(o Order) {
	fmt.Println(" Отмените заказ и попробуйте отменить его повторно.")
	fmt.Printf("Отменяем заказ: %v\n", o.Status)
	fmt.Printf("Отменяем заказ: %v\n", o.MarkCancelled())
	fmt.Printf("2 Отменяем заказ: %v\n", o.Status)
	fmt.Printf("Отмена еще раз: %v\n", o.MarkCancelled())

}
