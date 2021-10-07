package main

import (
	"github.com/Tike-Myson/darwin/pkg/models"
	"io/ioutil"
	"log"
	"os"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, Green+"INFO\t"+Reset, log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, Red+"ERROR\t"+Reset, log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		app.errorLog.Println(err)
	}

	//var allSignals []models.Signal
	var signals []models.Signal

	for _, f := range files {
		signals, err = app.parseFile(f.Name())
		if err != nil {
			app.errorLog.Println(err)
		}
		signals = app.SortByDate(signals)
		dates := app.GetDaysFromSignals(signals)

		for _, date := range dates {
			signalsByDay := app.GetSignalsFromDay(signals, date)
			err = app.WriteDataToCSVFile(signalsByDay, date, f.Name())
			//infoLog.Println("Write...")
			if err != nil {
				errorLog.Fatal(err)
			}
			infoLog.Println("Write...")
		}
	}
}
