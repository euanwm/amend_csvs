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
		removeCountryColumn(path.Join(csvDir, data.Name()))
	}
}

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
