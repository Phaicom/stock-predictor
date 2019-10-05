package model_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/phaicom/stock-predictor/pkg/model"
)

func TestModelPoint(t *testing.T) {
	pointPass := []struct {
		name     string
		Date     interface{}
		Open     interface{}
		High     interface{}
		Low      interface{}
		Close    interface{}
		AdjClose interface{}
		Volume   interface{}
	}{
		{
			name:     time.Now().String(),
			Date:     time.Now(),
			Open:     12.3,
			High:     13.3,
			Low:      10.2,
			Close:    13.5,
			AdjClose: 5.66,
			Volume:   uint64(12900),
		},
	}
	for _, point := range pointPass {
		t.Run(fmt.Sprintf("%s", point.name), func(t *testing.T) {
			result := model.Point{
				Date:     point.Date.(time.Time),
				Open:     point.Open.(float64),
				High:     point.High.(float64),
				Low:      point.Low.(float64),
				Close:    point.Close.(float64),
				AdjClose: point.AdjClose.(float64),
				Volume:   point.Volume.(uint64),
			}
			if reflect.DeepEqual(result, model.Point{}) {
				t.Errorf("Can not create point model")
			}
		})
	}
}
