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
