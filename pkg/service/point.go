package service

import (
	"log"
	"math"
	"sync"

	"github.com/phaicom/stock-predictor/pkg/model"

	"github.com/phaicom/stock-predictor/pkg/repository"
)

type pointService struct {
	Repo *repository.PointRepo
}

func NewPointService(repo *repository.PointRepo) PointService {
	return &pointService{
		Repo: repo,
	}
}

func (s *pointService) GetClosePriceProb(size int, diff float64) (probHight float64, probLow float64) {
	points, err := (*s.Repo).Fetch()
	if err != nil {
		log.Fatalln(err)
	}

	closePoints := [][]*model.Point{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := range points {
		wg.Add(1)
		go func(i int) {
			pointSet := findClosePriceSet(points[i:], size, diff)
			if len(pointSet) > 0 {
				mu.Lock()
				closePoints = append(closePoints, pointSet)
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Toubleshoot
	// for _, ps := range closePoints {
	// 	for _, p := range ps {
	// 		fmt.Printf("%+v\n", p)
	// 	}
	// 	fmt.Println("-1", ps[size].Close)
	// 	fmt.Println("-2", ps[size-1].Close)
	// 	fmt.Println("------------------------")
	// }
	// fmt.Printf("length: %v\n", len(closePoints))

	var higher, lower int
	for _, closePoint := range closePoints {
		if closePoint[size].Close > closePoint[size-1].Close {
			higher++
		} else {
			lower++
		}
	}

	if len(closePoints) == 0 {
		return
	}
	probHight = (float64(higher) * 100) / float64(len(closePoints))
	probLow = 100 - probHight
	return
}

func findClosePriceSet(points []*model.Point, size int, d float64) (result []*model.Point) {
	for i := 0; i < len(points); i++ {
		if len(result) == size {
			// append predict value
			result = append(result, points[i+1])
			return
		}

		for j := i + 1; j < len(points); j++ {
			if diff := points[j].Close - points[i].Close; math.Round(diff) == d {
				if len(result) == 0 {
					result = append(result, points[i])
				}
				result = append(result, points[j])
				i = j
				break
			}
		}
	}

	return []*model.Point{}
}
