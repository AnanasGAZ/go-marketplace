package main

func (w Wallet) HasEnoughMoney(amount int64) bool {
	return w.Balance >= amount
}

// ресивер???
func (w *Wallet) TopUp(amount int64) bool {
	// этот метод должен изменить u.Balance
	//Баланс и остаток товара не могут быть отрицательными.
	if amount <= 0 || w == nil { // надо ли проверить UserID и баланс не отрицательный?
		return false
	}
	//максимум тоже может стоит проверить? ну вдруг ошибка..
	w.Balance += amount
	return true
}

// Создайте пользователя с балансом 10000.
//
//	Передайте его в TopUpCopy и проверьте исходный баланс.
//	Передайте его адрес в TopUpPointer и снова проверьте баланс.
//	Коротко опишите в комментарии, почему результаты отличаются
func TopUpCopy(w Wallet, amount int64) {
	w.TopUp(amount)
}
func TopUpPointer(w *Wallet, amount int64) {
	w.TopUp(amount)
}

// Сумма должна быть больше нуля.
//  TopUp увеличивает баланс пользователя.
//  Withdraw не должен допускать отрицательный баланс.
//  Если денег недостаточно, баланс не изменяется и метод возвращает false.

// func (u User) HasEnoughMoney(amount int) bool //уже есть...
// func (u *User) TopUp(amount int) bool //уже есть...

// снятие
func (w *Wallet) Withdraw(amount int64) bool {
	if amount <= 0 {
		return false
	}
	if !w.HasEnoughMoney(amount) {
		return false
	}
	w.Balance -= amount
	return true
}

// Количество должно быть больше нуля.
//	IsAvailable проверяет, достаточно ли товара на складе.
//	AddStock увеличивает остаток.
//	Reserve уменьшает остаток, только если товара достаточно.
//	Если товара недостаточно, остаток не меняется

func (p Product) IsAvailable(quantity int64) bool {
	if quantity <= 0 {
		return false
	}
	if p.Stock < quantity {
		return false
	}
	return true
}
func (p *Product) AddStock(quantity int64) bool {
	if quantity <= 0 {
		return false
	}
	p.Stock += quantity
	return true
}
func (p *Product) Reserve(quantity int64) bool {
	if quantity <= 0 {
		return false
	}
	if !p.IsAvailable(quantity) {
		return false
	}
	p.Stock -= quantity
	return true
}

// Идентификатор товара и добавляемое количество должны быть больше нуля.
//
//	Если товар уже есть в корзине, Add увеличивает его количество.
//	Если в SetQuantity передано количество 0, товар удаляется.
//	Remove возвращает false, если товара в корзине нет.
//	Clear оставляет в корзине созданную пустую мапу.
//
// Особенность
// мапы
// Даже value receiver может изменить содержимое мапы внутри
// структуры, потому что копия мапы ссылается на те же данные. Для
// изменяющих методов Cart все равно используйте *Cart, чтобы
// намерение было явным.
func (c Cart) IsEmpty() bool {
	return len(c.Items) == 0
}
func (c *Cart) Add(productID int64, quantity int64) bool {
	if quantity <= 0 {
		return false
	}
	c.Items[productID] += quantity
	return true
}
func (c *Cart) SetQuantity(
	productID int64, quantity int64,
) bool {
	if quantity < 0 {
		return false
	}
	if quantity == 0 {
		delete(c.Items, productID)
	} else {
		c.Items[productID] = quantity
	}
	return true
}
func (c *Cart) Remove(productID int64) bool {
	if _, ok := c.Items[productID]; !ok {
		return false
	}
	delete(c.Items, productID)
	return true
}
func (c *Cart) Clear() {
	c.Items = make(map[int64]int64)
}

// IsPaid проверяет статус paid.
//	IsCancelled проверяет статус cancelled.
//	MarkCancelled меняет только статус paid на cancelled.
//	Повторная отмена не меняет заказ и возвращает false.

func (o Order) IsPaid() bool {
	return o.Status == "paid"
}
func (o Order) IsCancelled() bool {
	return o.Status == "cancelled"
}
func (o *Order) MarkCancelled() bool {
	if o.Status == "cancelled" {
		return false
	}
	o.Status = "cancelled"
	return true
}
