package main

var users = map[int]string{}
var balances = map[int]int{}
var productNames = map[int]string{}
var productPrices = map[int]int{}
var productStocks = map[int]int{}

var carts = map[int]map[int]int{}

// состав товара
var orderItems = map[int]map[int]int{}

// цены товаров на момент оформления
var orderItemPrices = map[int]map[int]int{}

// хранит статус заказа
var orderStatuses = map[int]string{}

// хранит историю операций пользователя
var operationHistory = map[int][]string{}

// склад
var stocks = map[int]int{}

// владелец заказа
var orderOwners = map[int]int{}

var orderTotals = map[int]int{}
