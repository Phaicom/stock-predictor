package service_test

import (
	"fmt"
	"testing"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/repository/point"
	"github.com/phaicom/stock-predictor/pkg/service"
)

func BenchmarkPointService(b *testing.B) {
	files := []struct {
		name string
	}{
		{"kbank-1000"},
		{"kbank-2500"},
		{"kbank-4923"},
	}

	for _, file := range files {
		// for k := 0.; k <= 10; k++ {
		// 	n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%s", file.name), func(b *testing.B) {
			csvStock := driver.OpenCSV("../../assets/" + file.name + ".csv")
			pointRepo := point.NewPointRepo(csvStock)
			pointService := service.NewPointService(&pointRepo)
			for i := 0; i < b.N; i++ {
				pointService.GetClosePriceProb(5, 7.0)
			}
		})
		// }
	}
	// csvStock := driver.OpenCSV("../../assets/kbank-1000.csv")
	// pointRepo := point.NewPointRepo(csvStock)
	// pointService := service.NewPointService(&pointRepo)
	// for i := 0; i < b.N; i++ {
	// 	pointService.GetClosePriceProb(5, 7.0)
	// }
}
