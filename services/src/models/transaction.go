package models

type Transaction struct {
	Id         string
	Amount     float64
	Dirtection TransactionDirection
}

type TransactionDirection struct {
	_direction    int
	_isReversible bool
}
