package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/model"
	"github.com/phaicom/stock-predictor/pkg/repository/point"
	"github.com/phaicom/stock-predictor/pkg/service"
)

func main() {
	fmt.Println("--File List--")
	files := listFile()
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("-------------")
	total := ""
	fmt.Print("Enter: ")
	fmt.Scanln(&total)
	if total == "" {
		total = "4923"
	}
	csvStruct, _ := driver.OpenCSV("assets/kbank-" + total + ".csv")
	pointRepo := point.NewPointRepo(csvStruct)
	pointService := service.NewPointService(&pointRepo)
	count, probHight, probLow, _ := pointService.GetClosePriceProb(5, 7.0)
	fmt.Printf("Total: %s\tCount: %v record(s)\nHigher: %.2f percent\tLower: %.2f percent\n", total, count, probHight, probLow)

	// Mock for creating difference file size
	// points, _ := pointRepo.Fetch()
	// createCSV(points, 100)
	// createCSV(points, 500)
	// createCSV(points, 1000)
	// createCSV(points, 2500)
	// createCSV(points, 4923)
}

func listFile() []string {
	var files []string

	root := "assets"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		r, _ := regexp.Compile("kbank-\\d*")
		if ok := r.MatchString(path); ok {
			file := r.FindString(path)

			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
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
