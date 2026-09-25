package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	fmt.Println("inside this damn shit")

	file, err := os.Open("weather_stations.csv")
	if err != nil {
		fmt.Println("error bro!!")
		fmt.Print(err.Error())
	}

	r := csv.NewReader(file)
	var chunk []string
	c := make(chan float64)
	sumValue := float64(0)
	var wg sync.WaitGroup
	minValue := float64(0)
	maxValue := float64(0)

	count := 0
	for len(chunk) < 10 {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		chunk = append(chunk, record[0])
		wg.Go(func() { Sum(c, chunk) })
		wg.Go(func() { Max(c, chunk, maxValue) })
		wg.Go(func() { Min(c, chunk, minValue) })

		sumValue = <-c
		maxValue = <-c
		minValue = <-c

		count++

		if len(chunk) == 10 {
			chunk = chunk[:0]
		}
	}

	fmt.Printf("sum: %f\n", sumValue)
	fmt.Printf("min: %f\n", minValue)
	fmt.Printf("max: %f\n", maxValue)
	fmt.Printf("count: %d\n", count)

	// fmt.Printf("sum :%f", <-sumValue)
	fmt.Println("inside this damn shit2")

}
