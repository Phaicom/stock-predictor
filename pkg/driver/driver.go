package driver

import (
	"encoding/csv"
	"log"
	"os"
)

// DB ...
type CSV struct {
	Records [][]string
}

// DBConn ...
var csvOpen = &CSV{}

func OpenCSV(name string) *CSV {
	csvfile, err := os.Open(name)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	records, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	csvOpen.Records = records[1:]
	return csvOpen
}
