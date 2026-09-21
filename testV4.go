package main

import "fmt"

func TestTopUp() {
	user := User{ID: 1, Name: "Alex", Balance: 10000}
	user.TopUp(5000)
	fmt.Println(user.Balance) // 15000
}
