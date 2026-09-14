package main

import "fmt"

// 1. Создаём пустые мапы.
var users = map[int]string{}
var balances = map[int]int{}

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
