package driver_test

import (
	"fmt"
	"testing"

	"github.com/phaicom/stock-predictor/pkg/driver"
)

func TestOpenCSV(t *testing.T) {
	filesPass := []struct {
		name string
	}{
		{"kbank-100"},
		{"kbank-500"},
		{"kbank-1000"},
		{"kbank-2500"},
		{"kbank-4923"},
	}
	for _, file := range filesPass {
		t.Run(fmt.Sprintf("%s", file.name), func(t *testing.T) {
			_, err := driver.OpenCSV("../../assets/" + file.name + ".csv")
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}

	filesFail := []struct {
		name string
	}{
		{"kbank-200"},
		{"kbank-1050"},
	}

	for _, file := range filesFail {
		t.Run(fmt.Sprintf("%s", file.name), func(t *testing.T) {
			_, err := driver.OpenCSV("../../assets/" + file.name + ".csv")
			if err == nil {
				t.Errorf(err.Error())
			}
		})
	}
}
