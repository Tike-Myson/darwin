package main

import (
	"encoding/csv"
	"fmt"
	"github.com/Tike-Myson/darwin/pkg/models"
	"os"
)

var exportDir = "./assets/out/csv_07.09-01.10/"

func (app *application) WriteDataToCSVFile(signals []models.Signal, date models.Date, filename string) error {
	updatedFilename := exportDir + app.GetExportFilename(filename, date)
	csvFile, err := os.Create(updatedFilename)
	if err != nil {
		return err
	}

	csvWriter := csv.NewWriter(csvFile)

	for _, signal := range signals {
		var convertedSignals []string
		convertedSignals = append(convertedSignals, signal.ActiveCurrency)
		convertedSignals = append(convertedSignals, signal.BaseCurrency)
		convertedSignals = append(convertedSignals, signal.ExchangeCell)
		convertedSignals = append(convertedSignals, signal.ExchangeBuy)
		convertedSignals = append(convertedSignals, fmt.Sprintf("%f", signal.PriceCell))
		convertedSignals = append(convertedSignals, fmt.Sprintf("%f", signal.PriceBuy))
		convertedSignals = append(convertedSignals, fmt.Sprintf("%f", signal.Amount))
		convertedSignals = append(convertedSignals, fmt.Sprintf("%f", signal.Profit))
		convertedSignals = append(convertedSignals, fmt.Sprintf("%f", signal.ROI))
		convertedSignals = append(convertedSignals, signal.Date.String())
		err = csvWriter.Write(convertedSignals)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	csvFile.Close()
	return nil
}