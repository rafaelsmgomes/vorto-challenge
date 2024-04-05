package main

import (
	"math"
)

// calculateDistance calculates the distance between two points
// this is the naive way of calculating the distance between two points
// we may be able to get some performance improvements by not squaring the values...?
func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
}

func calcAllDistances(loads []Load) []Load {
	for i, load := range loads {
		for j, otherLoad := range loads {
			if i == j {
				continue
			}
			loads[i].distBetween[j] = calculateDistance(load.x2, load.y2, otherLoad.x1, otherLoad.y1)
		}
	}

	return loads
}
