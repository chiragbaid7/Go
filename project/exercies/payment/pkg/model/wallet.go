package model

import "google.golang.org/genproto/googleapis/type/decimal"

type Wallet struct {
	Name    string
	Balance decimal.Decimal
}
