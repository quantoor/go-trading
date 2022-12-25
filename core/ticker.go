package core

type Ticker struct {
	Symbol    string
	BestAsk   float64
	BestBid   float64
	MidPrice  float64
	LastPrice float64
}
