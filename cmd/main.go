package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/model"
	"github.com/phaicom/stock-predictor/pkg/repository/point"
	"github.com/phaicom/stock-predictor/pkg/service"
)

func main() {
	csvStruct, _ := driver.OpenCSV("assets/kbank-4923.csv")
	pointRepo := point.NewPointRepo(csvStruct)
	pointService := service.NewPointService(&pointRepo)
	count, probHight, probLow, _ := pointService.GetClosePriceProb(5, 7.0)
	fmt.Printf("Count: %v, Higher: %v percent, Lower: %v percent\n", count, probHight, probLow)

	// Mock for creating difference file size
	// points, _ := pointRepo.Fetch()
	// createCSV(points, 100)
	// createCSV(points, 500)
	// createCSV(points, 1000)
	// createCSV(points, 2500)
	// createCSV(points, 4923)
}

// Mock for creating difference file size
func createCSV(points []*model.Point, size int) {
	name := fmt.Sprintf("assets/kbank-%s.csv", strconv.Itoa(size))
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, point := range points[:size] {
		open := fmt.Sprintf("%f", point.Open)
		high := fmt.Sprintf("%f", point.High)
		low := fmt.Sprintf("%f", point.Low)
		close := fmt.Sprintf("%f", point.Close)
		adjClose := fmt.Sprintf("%f", point.AdjClose)
		volumn := fmt.Sprintf("%d", point.Volume)
		err := writer.Write([]string{
			point.Date.String()[:10],
			open,
			high,
			low,
			close,
			adjClose,
			volumn,
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
