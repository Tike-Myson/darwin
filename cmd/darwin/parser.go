package main

import (
	"encoding/csv"
	"errors"
	"github.com/Tike-Myson/darwin/pkg/models"
	"io"
	"os"
	"strconv"
	"time"
)

var dir = "./assets/in/csv_07.09-01.10/"

func (app *application) parseFile(filename string) ([]models.Signal, error) {
	var signal models.Signal
	var signals []models.Signal
	csvfile, err := os.Open(dir + filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		signal.ActiveCurrency = record[0]
		signal.BaseCurrency = record[1]
		signal.ExchangeCell = record[2]
		signal.ExchangeBuy = record[3]
		priceCell, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			if !errors.Is(err, strconv.ErrSyntax) {
				return nil, err
			}
			priceCell = 0.0
		}
		signal.PriceCell = priceCell
		priceBuy, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			if !errors.Is(err, strconv.ErrSyntax) {
				return nil, err
			}
			priceBuy = 0.0
		}
		signal.PriceBuy = priceBuy
		amount, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			if !errors.Is(err, strconv.ErrSyntax) {
				return nil, err
			}
			amount = 0.0
		}
		signal.Amount = amount
		profit, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			if !errors.Is(err, strconv.ErrSyntax) {
				return nil, err
			}
			profit = 0.0
		}
		signal.Profit = profit
		ROI, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			if !errors.Is(err, strconv.ErrSyntax) {
				return nil, err
			}
			ROI = 0.0
		}
		signal.ROI = ROI
		//date, err := time.Parse("2006-01-02 15:04:05.000", record[9])
		date, err := time.Parse("2006-01-02 15:04:05.000", record[9])
		if err != nil {
			return nil, err
		}
		signal.Date = date
		signals = append(signals, signal)
	}
	return signals, nil
}
