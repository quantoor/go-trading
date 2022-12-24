package core

type Filter struct {
	Min  float64
	Max  float64
	Tick float64
}

type ProductInfo struct {
	Symbol        string
	BaseCurrency  string
	QuoteCurrency string
	ContractSize  int
	PriceFilter   Filter
	QtyFilter     Filter
}

type ProductsInfo map[string]ProductInfo
