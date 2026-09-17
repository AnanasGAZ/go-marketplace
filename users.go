package main

//???????? импорты нужны в каждый пакет?
import (
	"fmt"
	"strings"
)

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
	query1 := strings.ToLower(query)
	var result []int
	for id, name := range users {
		if strings.Contains(strings.ToLower(name), query1) {
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
