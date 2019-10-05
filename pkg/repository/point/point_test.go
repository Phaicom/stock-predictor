package point_test

import (
	"fmt"
	"testing"

	"github.com/phaicom/stock-predictor/pkg/repository/point"

	"github.com/phaicom/stock-predictor/pkg/driver"
)

func TestRepoPoint(t *testing.T) {
	csvStruct, _ := driver.OpenCSV("../../../assets/kbank-100.csv")

	t.Run("Pass", func(t *testing.T) {
		pointRepo := point.NewPointRepo(csvStruct)
		_, err := pointRepo.Fetch()
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	for i := 0; i < len(csvStruct.Records[0]); i++ {
		t.Run(fmt.Sprintf("Fail-%d", i), func(t *testing.T) {
			temp := csvStruct.Records[0][i]
			csvStruct.Records[0][i] = "2538%$--0"
			pointRepo := point.NewPointRepo(csvStruct)
			_, err := pointRepo.Fetch()
			if err == nil {
				t.Errorf(err.Error())
			}
			csvStruct.Records[0][i] = temp
		})
	}
}
