package main

import (
	"fmt"
	"log"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/repository/point"
)

func main() {
	csv := driver.OpenCSV("assets/KBANK.BK.csv")
	pointRepo := point.NewPointRepo(csv)
	points, err := pointRepo.Fetch()
	if err != nil {
		log.Fatalln(err)
	}

	for _, point := range points {
		fmt.Printf("point: %+v\n", point)
	}
}
