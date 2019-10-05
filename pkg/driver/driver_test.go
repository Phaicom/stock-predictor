package driver_test

import (
	"fmt"
	"io/ioutil"
	"os"
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
		t.Run(fmt.Sprintf("Pass-%s", file.name), func(t *testing.T) {
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
		t.Run(fmt.Sprintf("Fail-%s", file.name), func(t *testing.T) {
			_, err := driver.OpenCSV("../../assets/" + file.name + ".csv")
			if err == nil {
				t.Errorf(err.Error())
			}
		})
	}

	t.Run("Fail-BadCSV", func(t *testing.T) {
		name := "kbank-string.csv"
		desc := []byte("name,age\njohn,13,1995")
		f, _ := os.Create(name)
		defer f.Close()
		ioutil.WriteFile(name, desc, 0644)
		_, err := driver.OpenCSV(name)
		os.Remove(name)
		if err == nil {
			t.Errorf(err.Error())
		}
	})
}
