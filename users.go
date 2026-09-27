package main

//???????? импорты нужны в каждый пакет?
import (
	"fmt"
	"strings"
)

func AddUser(users map[int64]string, id int64, name string) bool {
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
func AddBalances(id int64) {
	balances[id] = 0
}
func GetUser(users map[int64]string, id int64) (string, bool) {
	val, exists := users[id]
	return val, exists
}
func RenameUser(users map[int64]string, id int64, newName string) bool {
	if _, exists := users[id]; !exists {
		return false
	}
	users[id] = newName
	return true
}
func DeleteUser(users map[int64]string, id int64) bool {
	if _, exists := users[id]; !exists {
		return false
	}
	delete(users, id)
	DeleteBalances(id)
	return true
}
func DeleteBalances(id int64) {
	delete(balances, id)
}
func FindUsersByName(users map[int64]string, query string) []int64 {
	query1 := strings.ToLower(query)
	var result []int64
	for id, name := range users {
		if strings.Contains(strings.ToLower(name), query1) {
			result = append(result, id)
		}
	}
	return result
}
func TopUpBalance(
	users map[int64]string,
	balances map[int64]int64,
	userID int64,
	amount int64,
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
func GetBalancePrint(in int64) string {
	return fmt.Sprintf("%d руб. %02d коп.", in/100, in%100)
}
func GetBalance(
	users map[int64]string,
	balances map[int64]int64,
	userID int64,
) (int64, bool) {
	if _, exists := users[userID]; !exists {
		return 0, false
	}
	return balances[userID], true
}
