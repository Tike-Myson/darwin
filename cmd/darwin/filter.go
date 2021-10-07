package main

import (
	"github.com/Tike-Myson/darwin/pkg/models"
	"time"
)

func (app *application) FilterByROI(signals []models.Signal, highPercent, lowPercent float64) []models.Signal {
	var sortedSignals []models.Signal
	for _, signal := range signals {
		if signal.ROI > lowPercent && signal.ROI < highPercent {
			sortedSignals = append(sortedSignals, signal)
		}
	}
	return sortedSignals
}

func (app *application) GetExchangeList(signals []models.Signal) []string {
	var exchanges []string
	for _, signal := range signals {
		if !app.IsValueExist(exchanges, signal.ExchangeCell) {
			exchanges = append(exchanges, signal.ExchangeCell)
		}
		if !app.IsValueExist(exchanges, signal.ExchangeBuy) {
			exchanges = append(exchanges, signal.ExchangeBuy)
		}
	}
	return exchanges
}

func (app *application) FilterByProfit(profitSum float64, signals []models.Signal) []models.TimeMarker {
	var (
		profitCounter float64
		startTimeMarker time.Time
		finishTimeMarker time.Time
		timeMarker models.TimeMarker
		timeMarkers []models.TimeMarker
	)

	for _, signal := range signals {
		if profitCounter == 0.0 {
			startTimeMarker = signal.Date
		}
		profitCounter += signal.Profit
		if profitCounter > profitSum {
			finishTimeMarker = signal.Date
			timeMarker.Profit = profitCounter
			timeMarker.TimeCounter = finishTimeMarker.Sub(startTimeMarker)
			timeMarkers = append(timeMarkers, timeMarker)
			profitCounter = 0.0
		}
	}
	return timeMarkers
}

func (app *application) FilterByExchange(exchangeName string, signals []models.Signal) []models.Signal {
	var filteredSignals []models.Signal
	for _, signal := range signals {
		if signal.ExchangeBuy == exchangeName || signal.ExchangeCell == exchangeName {
			filteredSignals = append(filteredSignals, signal)
		}
	}
	return filteredSignals
}

func (app *application) FilterByCurrencyPair(activeCurrency, baseCurrency string, signals []models.Signal) []models.Signal {
	var filteredSignals []models.Signal
	for _, signal := range signals {
		if signal.ActiveCurrency == activeCurrency && signal.BaseCurrency == baseCurrency {
			filteredSignals = append(filteredSignals, signal)
		}
	}
	return filteredSignals
}

func (app *application) GetDaysFromSignals(signals []models.Signal) []models.Date {
	var dates []models.Date
	var date models.Date
	for _, signal := range signals {
		 date.Year, date.Month, date.Day = signal.Date.Date()
		 if !app.IsDateExist(dates, date) {
			 dates = append(dates, date)
		 }
	}
	return dates
}

func (app *application) GetSignalsFromDay(signals []models.Signal, currDate models.Date) []models.Signal {
	var filteredSignals []models.Signal
	for _, signal := range signals {
		year, month, day := signal.Date.Date()
		if currDate.Day == day && currDate.Month == month && currDate.Year == year {
			filteredSignals = append(filteredSignals, signal)
		}
	}
	return filteredSignals
}
