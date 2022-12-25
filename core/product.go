package core

type ProductInfo struct {
	Type          string
	Symbol        string
	BaseCurrency  string
	QuoteCurrency string
	PriceTick     float64
	MinQty        float64
	QtyTick       float64
}

type ProductsInfo map[string]ProductInfo
