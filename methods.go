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
