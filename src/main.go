package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
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
	totalDist float64
	route     []int
}

func eachDriver(loads []Load, delivered []bool) Driver {
	driver := Driver{totalDist: 0, route: []int{}}
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
			// Condition 0) if cargo has already been delivered, skip it
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
		if cur == 0 && nextDeliveryIdx == -1 {
			for i, val := range delivered {
				if !val {
					log.Fatal("Can't make this delivery", i)
				}
			}
		}

		// if we found a point to deliver to, add it to the queue
		if nextDeliveryIdx != -1 {
			// 1. Go to pickup
			driver.totalDist += loads[cur].distBetween[nextDeliveryIdx]
			// add next point on itinerary
			q.Enqueue(nextDeliveryIdx)

		}

		if q.IsEmpty() {
			// means we are returning to origin
			driver.totalDist += loads[cur].distBetween[0]
		}
	}

	return driver
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("This program requires a file path as an argument")
		os.Exit(1)
	}

	lines, err := readLines()
	if err != nil {
		log.Fatal(err)
	}

	loads := parseLoads(lines)
	loads = calcAllDistances(loads)

	delivered := make([]bool, len(loads))

	for {
		if allTrue(delivered) {
			break
		}

		driver := eachDriver(loads, delivered)
		str := fmt.Sprint(driver.route)
		str = strings.ReplaceAll(str, " ", ",")
		fmt.Println(str)
	}

}
