package main

import (
	"fmt"
)

type Order interface {
	execute()
}

type BuyStock struct {
	stock Stock
}

type SellStock struct {
	stock Stock
}

func (bs BuyStock) execute() {
	bs.stock.buy()
}

func (ss SellStock) execute() {
	ss.stock.sell()
}

type Stock struct {
	name     string
	price    float64
	quantity int
}

func (stock *Stock) buy() {
	fmt.Printf("Stock [ name : %v, price : %v quantity : %v ] bought.\n",
		stock.name, stock.price, stock.quantity)
}

func (stock *Stock) sell() {
	fmt.Printf("Stock [ name : %v, price : %v quantity : %v ] sold.\n",
		stock.name, stock.price, stock.quantity)
}

type Broker struct {
	orders []Order
}

func (broker *Broker) takeOrder(order Order) {
	broker.orders = append(broker.orders, order)
}

func (broker *Broker) placeOrder() {
	for _, order := range broker.orders {
		order.execute()
	}
}

func main() {
	stock1 := Stock{"goog", 34.22, 50}
	stock2 := Stock{"facebook", 200, 70}
	bs1 := BuyStock{stock1}
	bs2 := BuyStock{stock2}
	ss1 := SellStock{stock1}
	ss2 := SellStock{stock2}
	broker := Broker{}
	broker.takeOrder(bs1)
	broker.takeOrder(bs2)
	broker.takeOrder(ss1)
	broker.takeOrder(ss2)
	broker.placeOrder()

	fmt.Println(`In command pattern a request is wrapped under a command object. Now the
	invoker processes it as per the command.`)
}
