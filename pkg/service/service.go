package service

type PointService interface {
	GetClosePriceProb(size int, diff float64) (count int, probHight float64, probLow float64, err error)
}
