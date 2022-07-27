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
		fpath := path.Join(csvDir, data.Name())
		removeCountryColumn(fpath)
		fillEmptyColumns(fpath)
		removeEmptyResults(fpath)
	}
}

//fillEmptyColumns A lot of the results files - where people have bombed out or not taken attempts - do not have any contents.
//This could throw some issues in pulling them into the larger database.
func fillEmptyColumns(csvfp string) {
	csvFile, err := os.Open(csvfp)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	bigSlice, _ := reader.ReadAll()
	for _, data := range bigSlice {
		for index, contents := range data {
			if len(contents) == 0 {
				data[index] = "0"
			}
		}
	}
	writeCSV(csvfp, bigSlice)
}

//writeCSV Writes CSV file, first arg is the filepath/name. Second is the bigSlice data.
func writeCSV(csvFp string, bigSlice [][]string) {
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
	writeCSV(csvFp, bigSlice)
}
