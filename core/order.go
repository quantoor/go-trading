package core

type OrderId string
type Timestamp uint64

type OrderType string

const (
	OrderTypeLimit  OrderType = "limit"
	OrderTypeMarket           = "market"
	OrderTypeStop             = "stop"
)

type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "good_till_cancelled"
	TimeInForceFOK             = "fill_or_fill"
	TimeInForceIOC             = "immediate_or_cancel"
)

type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "new"
	OrderStatusPartiallyFilled             = "partially_filled"
	OrderStatusFilled                      = "filled"
	OrderStatusCancelled                   = "cancelled"
)

type PlaceOrderParams struct {
	Symbol      string
	Qty         float64
	Price       float64
	IsBuy       bool
	OrderType   OrderType
	TimeInForce TimeInForce
	StopPrice   float64
	ReduceOnly  bool
	PostOnly    bool
}

type GetOrderResponse struct {
	OrderID      string      `json:"order_id"`
	Symbol       string      `json:"symbol"`
	IsBuy        bool        `json:"is_buy"`
	OrderType    OrderType   `json:"order_type"`
	Price        float64     `json:"price"`
	Qty          float64     `json:"qty"`
	TimeInForce  TimeInForce `json:"time_in_force"`
	OrderStatus  OrderStatus `json:"order_status"`
	FilledQty    float64     `json:"filled_qty"`
	CumExecFee   float64     `json:"cum_exec_fee"`
	Label        string      `json:"label"`
	PlacedTs     Timestamp   `json:"created_at"`
	UpdatedTs    Timestamp   `json:"updated_at"`
	TriggerPrice float64     `json:"stop_loss"`
}
