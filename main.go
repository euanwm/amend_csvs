package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	const csvDir string = "AUS"
	csvFiles, err := os.ReadDir(csvDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, data := range csvFiles {
		removeEmptyResults(path.Join(csvDir, data.Name()))
		removeCountryColumn(path.Join(csvDir, data.Name()))
	}
}

//removeEmptyResults Removes any csv file that doesn't contain any results, not sure why it happens
func removeEmptyResults(csvFp string) {
	csvFile, err := os.Open(csvFp)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	bigSlice, _ := reader.ReadAll()
	if len(bigSlice) == 1 {
		err := os.Remove(csvFp)
		if err != nil {
			log.Fatal(err)
		}

	}
}

//removeCountryColumn When creating the Australia weightlifting results I stupidly added in AUS to each result line.
//this removes that annoying column because it fucks up stuff higher in the stack
func removeCountryColumn(csvFp string) {
	csvFile, err := os.Open(csvFp)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	bigSlice, _ := reader.ReadAll()
	for index, line := range bigSlice {
		if len(line) == 15 {
			bigSlice[index] = line[:14]
		}
	}
	newCsvFile, err := os.Create(csvFp)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(newCsvFile)
	writeData := writer.WriteAll(bigSlice)
	if writeData != nil {
		fmt.Println(writeData)
	}
}
