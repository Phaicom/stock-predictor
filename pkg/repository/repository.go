package repository

import "github.com/phaicom/stock-predictor/pkg/model"

type PointRepo interface {
	Fetch() ([]*model.Point, error)
}
