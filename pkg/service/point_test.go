package service_test

import (
	"fmt"
	"testing"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/repository/point"
	"github.com/phaicom/stock-predictor/pkg/service"
)

func TestServicePoint(t *testing.T) {
	files := []struct {
		name  string
		count int
	}{
		{
			name:  "kbank-500",
			count: 0,
		},
		{
			name:  "kbank-1000",
			count: 239,
		},
		{
			name:  "kbank-2500",
			count: 2254,
		},
		{
			name:  "kbank-4923",
			count: 4735,
		},
	}

	for _, file := range files {
		t.Run(fmt.Sprintf("Pass-%s", file.name), func(t *testing.T) {
			csvStruct, _ := driver.OpenCSV("../../assets/" + file.name + ".csv")
			pointRepo := point.NewPointRepo(csvStruct)
			pointService := service.NewPointService(&pointRepo)
			count, _, _, err := pointService.GetClosePriceProb(5, 7.0)
			if err != nil {
				t.Errorf(err.Error())
			}
			if count != file.count {
				t.Errorf("Count value mismatch")
			}
		})
	}

	t.Run("Fail-RepoFetch", func(t *testing.T) {
		csvStruct, _ := driver.OpenCSV("../../assets/kbank-500.csv")
		csvStruct.Records[0][0] = "hello, world!"
		pointRepo := point.NewPointRepo(csvStruct)
		pointService := service.NewPointService(&pointRepo)
		_, _, _, err := pointService.GetClosePriceProb(5, 7.0)
		if err == nil {
			t.Errorf(err.Error())
		}
	})
}

func BenchmarkPointService(b *testing.B) {
	files := []struct {
		name string
	}{
		{"kbank-1000"},
		{"kbank-2500"},
		{"kbank-4923"},
	}

	for _, file := range files {
		b.Run(fmt.Sprintf("%s", file.name), func(b *testing.B) {
			csvStruct, _ := driver.OpenCSV("../../assets/" + file.name + ".csv")
			pointRepo := point.NewPointRepo(csvStruct)
			pointService := service.NewPointService(&pointRepo)
			for i := 0; i < b.N; i++ {
				pointService.GetClosePriceProb(5, 7.0)
			}
		})
	}
}
