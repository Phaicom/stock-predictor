package model

import (
	"time"
)

type Point struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	AdjClose float64
	Volume   uint64
}
