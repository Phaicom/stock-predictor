package service

type PointService interface {
	GetClosePriceProb(size int, diff float64) (probHight float64, probLow float64)
}
