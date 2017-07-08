package main

import "math"

const (
	POINTMAX = 3000
	POINTMIN = 0
	BASE     = 200
)

func TransferPoints(win, lose, base int64) int64 {
	winfloat := float64(win)
	losefloat := float64(lose)

	delta := math.Abs(winfloat - losefloat)
	avg := math.Min(winfloat, losefloat) + (delta / 2)

	factor := 1 - (avg / POINTMAX)
	points := factor * float64(base)

	winloss := ((losefloat - winfloat) / POINTMAX) * points
	points += winloss
	return int64(math.Floor(points))
}
