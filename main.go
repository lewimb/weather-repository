package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
	"time"
)

type Result struct {
	sum   float64
	min   float64
	max   float64
	count int
}

const LIMIT = 1000

func main() {
	fmt.Println("Waiting...")
	startTime := time.Now()

	file, err := os.Open("weather_stations.csv")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer file.Close()

	res := make(chan Result)
	chunk := make([]string, 0)

	r := csv.NewReader(file)
	var wg sync.WaitGroup

	process := func(c []string) {
		wg.Go(func() {
			res <- processFiles(c)
		})
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		chunk = append(chunk, record[0])
		if len(chunk) == LIMIT {
			process(chunk)
			chunk = make([]string, 0, LIMIT)
		}
	}

	if len(chunk) > 0 {
		process(chunk)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	var sum, min, max float64 = 0, math.Inf(1), math.Inf(-1)
	count := 0

	for result := range res {
		sum += result.sum
		min = math.Min(min, result.min)
		max = math.Max(max, result.max)
		count += result.count
	}

	fmt.Printf("sum : %f\n", sum)
	fmt.Printf("min : %f\n", min)
	fmt.Printf("max : %f\n", max)
	fmt.Printf("count : %d\n", count)
	fmt.Printf("avg: %f\n", sum/float64(count))
	fmt.Printf("Finished at: %f\n", (time.Since(startTime).Seconds()))

}
