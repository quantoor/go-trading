package core

import "time"

type OrderEvent struct {
	OrderId     OrderId
	Label       string
	Symbol      string
	IsBuy       bool
	OrderType   OrderType
	Price       float64
	Qty         float64
	FilledQty   float64
	CumExecFee  float64
	TimeInForce TimeInForce
	CreateType  string
	CancelType  string
	Status      OrderStatus
	StopPrice   float64
	CreateTime  time.Time
	UpdateTime  time.Time
	ReduceOnly  bool
	PostOnly    bool
}

type TradeEvent struct {
	Symbol       string
	IsBuy        bool
	OrderId      OrderId
	TradeId      string
	Label        string
	Price        float64
	OrderQty     float64
	ExecType     string
	FilledQty    float64
	ExecFee      float64
	RemainingQty float64
	IsMaker      bool
	TradeTime    time.Time
}
