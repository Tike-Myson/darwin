package main

import (
	"github.com/Tike-Myson/darwin/pkg/models"
	"strconv"
	"time"
)

func (app *application) IsValueExist(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func (app *application) IsDateExist(dates []models.Date, date models.Date) bool {
	for _, d := range dates {
		if d.Day == date.Day && d.Month == date.Month && d.Year == date.Year {
			return true
		}
	}
	return false
}

func (app *application) GetAverageValue(timeMarkers []models.TimeMarker) time.Duration {
	var allTimeMarkersSum time.Duration
	for _, timeMarker := range timeMarkers {
		allTimeMarkersSum += timeMarker.TimeCounter
	}

	return time.Duration(allTimeMarkersSum.Nanoseconds() / int64(len(timeMarkers)))
}

func (app *application) GetExportFilename(filename string, date models.Date) string {
	var updatedFilename string
	bFilename := []byte(filename)
	for i := 0; i < 15; i++ {
		bFilename = bFilename[:len(bFilename) - 1]
	}
	strDay := strconv.Itoa(date.Day)
	strMonth := strconv.Itoa(int(date.Month))
	updatedFilename = string(bFilename) + strDay + "." + strMonth + ".csv"
	return updatedFilename
}
