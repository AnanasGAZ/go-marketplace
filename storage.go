package main

var users = map[int64]string{}        //ид, имя
var balances = map[int64]int64{}      //ид пользователя, деньги
var productNames = map[int64]string{} //ид, название продукта
var productPrices = map[int64]int64{} //ид, цена
var productStocks = map[int64]int64{} //ид, количество

var carts = map[int64]map[int64]int64{} //carts[userID][productID] += quantity

// состав товара
var orderItems = map[int64]map[int64]int64{}

// цены товаров на момент оформления
var orderItemPrices = map[int64]map[int64]int64{}

// хранит статус заказа
var orderStatuses = map[int64]string{}

// хранит историю операций пользователя
var operationHistory = map[int64][]string{}

// склад
var stocks = map[int64]int64{}

// владелец заказа
var orderOwners = map[int64]int64{}

var orderTotals = map[int64]int64{}
