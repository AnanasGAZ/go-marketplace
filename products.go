package main

import (
	"fmt"
	"strings"
)

func AddProduct(
	names map[int]string, prices map[int]int, stocks map[int]int,
	id int, name string, price int, stock int,
) bool {
	if id <= 0 {
		return false
	}
	if _, exists := names[id]; exists {
		return false
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return false
	}
	if price <= 0 {
		return false
	}
	if stock < 0 { //количество товара на складе
		return false
	}

	names[id] = name
	prices[id] = price
	stocks[id] = stock

	return true

}

func GetProduct(
	names map[int]string, prices map[int]int, stocks map[int]int,
	productID int,
) (string, int, int, bool) {
	name, exists := names[productID]
	if !exists {
		return "", 0, 0, false
	}
	//Получение товара должно возвращать название, цену, остаток и признак существования.
	return name, prices[productID], stocks[productID], true

}

func UpdateProductStock(
	names map[int]string, stocks map[int]int,
	productID int, stock int,
) bool {
	if _, exists := names[productID]; !exists {
		return false
	}
	if stock < 0 {
		return false
	}
	stocks[productID] = stock
	return true
}

func SearchProducts(names map[int]string, query string) []int {
	fmt.Printf("ЗАпрос:%v\n", len(query))
	result := []int{}
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return result
	}

	fmt.Printf("ЗАпрос:%v\n", len(query))

	for id, name := range names {
		name = strings.ToLower(name)

		if strings.Contains(name, query) {
			result = append(result, id)
		}
	}
	return result

}
