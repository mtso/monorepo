package main

import "math"
import "log"

const (
	POINTMAX = 3000
	POINTMIN = 0
	BASE     = 200
)

func TransferPoints(win, lose, base int) int {
	winfloat := float64(win)
	losefloat := float64(lose)

	delta := math.Abs(winfloat - losefloat)
	avg := math.Min(winfloat, losefloat) + (delta / 2)

	factor := 1 - (avg / POINTMAX)
	points := factor * float64(base)

	winloss := ((losefloat - winfloat) / POINTMAX) * points
	points += winloss
	return int(math.Floor(points))
}

func main() {
	pt := TransferPoints(2000, 1000, BASE)
	log.Println(pt)
}
