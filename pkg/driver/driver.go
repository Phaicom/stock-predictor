package driver

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type CSV struct {
	Records [][]string
}

func OpenCSV(name string) (csvStruct *CSV, err error) {
	csvfile, err := os.Open(name)
	if err != nil {
		err = errors.New(fmt.Sprintln("Couldn't open the csv file", err))
		return
	}
	defer csvfile.Close()

	records, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		return
	}
	csvStruct = &CSV{}
	csvStruct.Records = records[1:]
	return
}
