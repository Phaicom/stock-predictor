package point

import (
	"strconv"
	"time"

	"github.com/phaicom/stock-predictor/pkg/driver"
	"github.com/phaicom/stock-predictor/pkg/model"
	"github.com/phaicom/stock-predictor/pkg/repository"
)

type pointRepo struct {
	CSV *driver.CSV
}

func NewPointRepo(csv *driver.CSV) repository.PointRepo {
	return &pointRepo{
		CSV: csv,
	}
}

func (p *pointRepo) Fetch() ([]*model.Point, error) {
	points := []*model.Point{}
	for _, record := range p.CSV.Records {
		if record[1] == "null" {
			continue
		}

		date, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, err
		}
		open, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		high, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		low, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, err
		}
		close, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, err
		}
		adjClose, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return nil, err
		}
		volume, err := strconv.ParseUint(record[6], 10, 64)
		if err != nil {
			return nil, err
		}

		point := model.Point{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			AdjClose: adjClose,
			Volume:   volume,
		}
		points = append(points, &point)
	}
	return points, nil
}
