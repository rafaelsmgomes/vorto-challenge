package main

import (
	"fmt"
	"math"
	"os"
)

type Load struct {
	ID int

	x1           float64
	y1           float64
	x2           float64
	y2           float64
	deliveryTime float64
	distBetween  []float64
}

type Driver struct {
	ID        int
	totalDist float64
	route     []int
}

// what do we need? We need to build an array of points that we're gonna go to
// forget about the restrictions for a while (no more than 12 hours driving)
// So, for each point we go to, we need to to remove it from the list of available points
// then we do the load and after we go to the next point

func oneDriver(loads []Load, delivered []bool) Driver {
	driver := Driver{ID: 1, totalDist: 0, route: []int{}}
	upperBound := 720.0

	// I need to start at point 0 and move into each next nearest point in the list
	q := NewQueue()
	q.Enqueue(0)

	for {
		if q.IsEmpty() {
			break
		}
		// get the next point to deliver
		cur := q.Dequeue().(int)
		if cur != 0 {
			driver.route = append(driver.route, cur)
		}

		// lastIdx is the last point delivered
		driver.totalDist += loads[cur].deliveryTime
		delivered[cur] = true
		// fmt.Printf("Driver totalDist: %f\n", driver.totalDist)

		minDist := math.Inf(1)
		nextDeliveryIdx := -1
		for i, dist := range loads[cur].distBetween {
			// if cargo has already been delivered, skip it
			if delivered[i] {
				continue
			}

			// Condition 1) Is this a valid delivery?
			// We need to check if the driver can make the delivery and come back to the origin
			totalDistAfterDelivery := driver.totalDist + dist + loads[i].distBetween[0] + loads[i].deliveryTime
			if totalDistAfterDelivery > upperBound {
				continue
			}
			// Condition 2) Is this the shortest distance?
			// is so, update
			if dist < minDist {
				minDist = dist
				nextDeliveryIdx = i
			}
		}

		// if we found a point to deliver to, add it to the queue
		if nextDeliveryIdx != -1 {
			// 1. Go to pickup
			driver.totalDist += loads[cur].distBetween[nextDeliveryIdx]
			// fmt.Printf("Go to pickup: Driver totalDist: %f\n", driver.totalDist)

			q.Enqueue(nextDeliveryIdx)

		}

		if q.IsEmpty() {
			// fmt.Printf("Drive back to origin: %.2f\n", loads[cur].distBetween[0])
			driver.totalDist += loads[cur].distBetween[0]
		}
	}

	return driver
}

// let's calculate the smallest time for a single driver to do all of this

func main() {
	if len(os.Args) < 1 {
		fmt.Println("This program requires a file path as an argument")
		os.Exit(1)
	}

	lines := readLines()
	loads := parseLoads(lines)
	loads = calcAllDistances(loads)

	delivered := make([]bool, len(loads))
	fmt.Println(delivered)

	driver := oneDriver(loads, delivered)
	fmt.Println(driver)
	for {
		if allTrue(delivered) {
			break
		}

		driver := oneDriver(loads, delivered)
		fmt.Println(driver.route)
		fmt.Println(delivered)
	}

}
