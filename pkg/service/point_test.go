package service_test

import (
	"fmt"
	"testing"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/repository/point"
	"github.com/phaicom/stock-predictor/pkg/service"
)

func TestServicePoint(t *testing.T) {
	csvStruct, _ := driver.OpenCSV("../../assets/kbank-1000.csv")

	t.Run("Pass", func(t *testing.T) {
		pointRepo := point.NewPointRepo(csvStruct)
		pointService := service.NewPointService(&pointRepo)
		count, _, _, err := pointService.GetClosePriceProb(5, 7.0)
		if err != nil {
			t.Errorf(err.Error())
		}
		if count != 239 {
			t.Errorf("Count value mismatch")
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
