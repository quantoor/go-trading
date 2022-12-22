package tradingo

import "fmt"

type OrderBook struct {
	Symbol string
	Asks   OrderBookOrders
	Bids   OrderBookOrders
}

func (ob *OrderBook) Update(price float64, quantity float64, isBuy bool) {
	if quantity == 0 {
		if isBuy {
			ob.Bids.delete(price)
		} else {
			ob.Asks.delete(price)
		}
	} else {
		if isBuy {
			ob.Bids.insertBid(price, quantity)
			bestAsk := ob.BestAsk()
			for bestAsk != nil && price >= bestAsk.price {
				ob.Asks.deleteHead()
				bestAsk = ob.BestAsk()
			}
		} else {
			ob.Asks.insertAsk(price, quantity)
			bestBid := ob.BestBid()
			for bestBid != nil && price <= bestBid.price {
				ob.Bids.deleteHead()
				bestBid = ob.BestBid()
			}
		}
	}
}

func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{Symbol: symbol}
}

func (ob *OrderBook) BestAsk() *OrderBookOrder {
	return ob.Asks.head
}

func (ob *OrderBook) BestBid() *OrderBookOrder {
	return ob.Bids.head
}

func (ob *OrderBook) IsEmpty() bool {
	return ob.Asks.head == nil && ob.Bids.head == nil
}

func (ob *OrderBook) Reset() {
	ob.Asks = OrderBookOrders{}
	ob.Bids = OrderBookOrders{}
}

func (ob *OrderBook) PrintBestBidAsk() {
	if ob.BestAsk() != nil && ob.BestBid() != nil {
		//fmt.Printf("%+v -> %+v\n", ob.BestBid().price, ob.BestAsk().price)
		fmt.Printf("%+v, %+v -> %+v, %+v\n", ob.BestBid().price, ob.BestBid().quantity, ob.BestAsk().price, ob.BestAsk().quantity)
	}
}

type OrderBookOrder struct {
	prev     *OrderBookOrder
	next     *OrderBookOrder
	price    float64
	quantity float64
}

type OrderBookOrders struct {
	head *OrderBookOrder
	tail *OrderBookOrder
}

func (o *OrderBookOrders) insertBid(price float64, quantity float64) {
	if o.head == nil {
		o.insertHead(price, quantity)
	}

	if price > o.head.price {
		o.insertHead(price, quantity)
	} else if price < o.tail.price {
		o.insertTail(price, quantity)
	} else {

		l := o.head
		for l != nil {

			if l.price == price {
				l.quantity = quantity
				return
			} else if price < l.price && l.next != nil && price > l.next.price {
				bid := &OrderBookOrder{
					prev:     l,
					next:     l.next,
					price:    price,
					quantity: quantity,
				}

				l.next.prev = bid
				l.next = bid
				return
			}

			l = l.next
		}
	}
}

func (o *OrderBookOrders) insertAsk(price float64, quantity float64) {
	if o.head == nil {
		o.insertHead(price, quantity)
	}

	if price < o.head.price {
		o.insertHead(price, quantity)
	} else if price > o.tail.price {
		o.insertTail(price, quantity)
	} else {

		l := o.head
		for l != nil {

			if l.price == price {
				l.quantity = quantity
				return
			} else if price > l.price && l.next != nil && price < l.next.price {
				bid := &OrderBookOrder{
					prev:     l,
					next:     l.next,
					price:    price,
					quantity: quantity,
				}

				l.next.prev = bid
				l.next = bid
				return
			}

			l = l.next
		}
	}
}

func (o *OrderBookOrders) delete(price float64) {
	if o.head == nil {
		return
	}
	if price == o.head.price {
		o.deleteHead()
		return
	}
	if price == o.tail.price {
		o.deleteTail()
		return
	}

	l := o.head
	for l != nil {
		if l.price == price {

			if l.prev != nil {
				l.prev.next = l.next
			}
			if l.next != nil {
				l.next.prev = l.prev
			}
			l = nil
			return
		}
		l = l.next
	}
}

func (o *OrderBookOrders) insertHead(price float64, quantity float64) {
	list := &OrderBookOrder{
		next:     o.head,
		price:    price,
		quantity: quantity,
	}

	if o.head == nil {
		o.tail = list
	} else {
		o.head.prev = list
	}

	o.head = list
}

func (o *OrderBookOrders) insertTail(price float64, quantity float64) {
	list := &OrderBookOrder{
		prev:     o.tail,
		price:    price,
		quantity: quantity,
	}

	if o.tail == nil {
		o.head = list
	} else {
		o.tail.next = list
	}

	o.tail = list
}

func (o *OrderBookOrders) deleteHead() {
	if o.head.next != nil {
		o.head = o.head.next
		o.head.prev = nil
	} else {
		o.head = nil
		o.tail = nil
	}
}

func (o *OrderBookOrders) deleteTail() {
	if o.tail.prev != nil {
		o.tail = o.tail.prev
		o.tail.next = nil
	} else {
		o.head = nil
		o.tail = nil
	}
}
