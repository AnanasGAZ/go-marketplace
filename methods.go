package main

func (u User) HasEnoughMoney(amount int) bool {
	return u.Balance >= amount
}

// ресивер???
func (u *User) TopUp(amount int) bool {
	// этот метод должен изменить u.Balance
	u.Balance += amount
	return true
}

// Сумма должна быть больше нуля.
//  TopUp увеличивает баланс пользователя.
//  Withdraw не должен допускать отрицательный баланс.
//  Если денег недостаточно, баланс не изменяется и метод возвращает false.

// func (u User) HasEnoughMoney(amount int) bool //уже есть...
// func (u *User) TopUp(amount int) bool //уже есть...

// снятие
func (u *User) Withdraw(amount int) bool {
	if amount <= 0 {
		return false
	}
	if !u.HasEnoughMoney(amount) {
		return false
	}
	u.Balance -= amount
	return true
}

// Количество должно быть больше нуля.
//	IsAvailable проверяет, достаточно ли товара на складе.
//	AddStock увеличивает остаток.
//	Reserve уменьшает остаток, только если товара достаточно.
//	Если товара недостаточно, остаток не меняется

func (p Product) IsAvailable(quantity int) bool {
	if quantity <= 0 {
		return false
	}
	if p.Stock < quantity {
		return false
	}
	return true
}
func (p *Product) AddStock(quantity int) bool {
	if quantity <= 0 {
		return false
	}
	p.Stock += quantity
	return true
}
func (p *Product) Reserve(quantity int) bool {
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
func (c *Cart) Add(productID int, quantity int) bool {
	if quantity <= 0 {
		return false
	}
	c.Items[productID] += quantity
	return true
}
func (c *Cart) SetQuantity(
	productID int, quantity int,
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
func (c *Cart) Remove(productID int) bool {
	if _, ok := c.Items[productID]; !ok {
		return false
	}
	delete(c.Items, productID)
	return true
}
func (c *Cart) Clear() {
	c.Items = make(map[int]int)
}

// IsPaid проверяет статус paid.
//	IsCancelled проверяет статус cancelled.
//	MarkCancelled меняет только статус paid на cancelled.
//	Повторная отмена не меняет заказ и возвращает false.

func (o Order) IsPaid() bool {
	return o.Status == "paid"
}
func (o Order) IsCancelled() bool {
	return o.Status == "canceled"
}
func (o *Order) MarkCancelled() bool {
	if o.Status == "canceled" {
		return false
	}
	return true
}
