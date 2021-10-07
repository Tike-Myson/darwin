package main

import (
	"github.com/Tike-Myson/darwin/pkg/models"
	"sort"
	"strconv"
)

type timeSlice []models.Signal

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].Date.Before(p[j].Date)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (app *application) SortByDate(signals []models.Signal) []models.Signal {
	var signalsMap = make(map[string]models.Signal)
	for i := range signals {
		signalsMap[strconv.Itoa(i)] = signals[i]
	}
	sortedSignals := make(timeSlice, 0, len(signalsMap))
	for _, d := range signalsMap {
		sortedSignals = append(sortedSignals, d)
	}
	sort.Sort(sortedSignals)
	return sortedSignals
}
