package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func processFiles(rows []string) Result {
	sum := float64(0)
	maxVal := math.Inf(-1)
	minVal := math.Inf(1)

	for _, row := range rows {
		val := strings.Split(row, ";")
		parseVal, err := strconv.ParseFloat(val[1], 64)
		if err != nil {
			fmt.Println(err.Error())
			panic("sum error")
		}
		minVal = math.Min(minVal, parseVal)
		maxVal = math.Max(maxVal, parseVal)
		sum += parseVal

	}

	return Result{
		sum:   sum,
		min:   minVal,
		max:   maxVal,
		count: len(rows),
	}
}
