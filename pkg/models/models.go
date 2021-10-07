package models

import "time"

type Signal struct {
	ActiveCurrency string
	BaseCurrency   string
	ExchangeCell   string
	ExchangeBuy    string
	PriceCell      float64
	PriceBuy       float64
	Amount         float64
	Profit         float64
	ROI            float64
	Date           time.Time
}

type TimeMarker struct {
	Profit      float64
	TimeCounter time.Duration
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}
