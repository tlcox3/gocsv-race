package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gocarina/gocsv"
)

type dataStruct struct {
	Column1 string `csv:"COLUMN1"`
	Column2 string `csv:"COLUMN2"`
}

func main() {
	csv1 := "COLUMN1,COLUMN2\ntest1,test2"
	csv2 := "COLUMN1|COLUMN2\ntest3|test4"
	s, err := doStuff(csv1, csv2)

	if err != nil {
		log.Fatalf("Failed to parse data, err: %v", err)
	}
	fmt.Println(s)
}

func doStuff(d1, d2 string) (string, error) {
	d1Chan := make(chan []dataStruct)
	err1Chan := make(chan error)
	go parse(d1, ',', d1Chan, err1Chan)

	// Uncomment following line to remove effects of race (tests results pass)
	// time.Sleep(200 * time.Millisecond)

	d2Chan := make(chan []dataStruct)
	err2Chan := make(chan error)
	go parse(d2, '|', d2Chan, err2Chan)

	var s string
	select {
	case d := <-d1Chan:
		s = fmt.Sprintf("Data1 Col1: %v", d[0].Column1)
	case err := <-err1Chan:
		return "", fmt.Errorf("failed to read csv1: %w", err)
	}

	select {
	case d := <-d2Chan:
		s = fmt.Sprintf("%v Data2 Col2: %v", s, d[0].Column2)
	case err := <-err2Chan:
		return "", fmt.Errorf("failed to read csv1: %w", err)
	}

	return s, nil
}

func parse(csvString string, delimiter rune, dataChan chan<- []dataStruct, errChan chan<- error) {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = delimiter
		return r
	})
	time.Sleep(100 * time.Millisecond)
	var data []dataStruct

	err := gocsv.UnmarshalString(csvString, &data)
	if err != nil {
		errChan <- err
		return
	}
	fmt.Printf("CSV: %v\nData: %+v\n\n", csvString, data)

	dataChan <- data
}
