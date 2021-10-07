package main

import "github.com/Tike-Myson/darwin/pkg/models"

func (app *application) GetExchangeInvolvement(signals []models.Signal) map[string]float64 {
	m := make(map[string]float64)
	for _, signal := range signals {
		m[signal.ExchangeBuy]++
		m[signal.ExchangeCell]++
	}

	for key, value := range m {
		m[key] = value / float64(len(signals) * 2) * 100
	}

	return m
}