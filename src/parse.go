package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strconv"
	"strings"
)

func readLines() ([]string, error) {

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	if _, _, err := reader.ReadLine(); err != nil {
		return nil, err
	}

	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		lines = append(lines, string(line))
	}

	return lines, nil
}

func parseCoordinates(s string) (float64, float64) {
	var x, y float64
	n, err := fmt.Sscanf(s, "(%f,%f)", &x, &y)
	if n != 2 || err != nil {
		fmt.Println("Error parsing coordinates:", err)
	}
	return x, y
}

func parseLoads(lines []string) []Load {
	loads := []Load{}

	// Add origin to the list of points
	initDist := make([]float64, len(lines)+1)
	loads = append(loads, Load{ID: 0, x1: 0, y1: 0, x2: 0, y2: 0, deliveryTime: 0, distBetween: initDist})

	for _, line := range lines {
		fields := strings.Fields(line)

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error parsing ID:", err)
			continue
		}

		x1, y1 := parseCoordinates(fields[1])
		x2, y2 := parseCoordinates(fields[2])
		myDistance := calculateDistance(x1, y1, x2, y2)
		distancesTo := make([]float64, (len(lines) + 1))

		loads = append(loads, Load{
			ID:           id,
			x1:           x1,
			y1:           y1,
			x2:           x2,
			y2:           y2,
			deliveryTime: myDistance,
			distBetween:  distancesTo,
		})
	}

	return loads
}
