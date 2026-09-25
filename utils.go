package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Sum(channel chan float64, chunks []string) {
	sum := float64(0)
	for _, chunk := range chunks {
		val := strings.Split(chunk, ";")
		parseVal, err := strconv.ParseFloat(val[1], 64)
		if err != nil {
			fmt.Println(err.Error())
			panic("sum error")
		}
		sum += parseVal
	}

	channel <- sum
}

func Min(channel chan float64, chunks []string, minVal float64) {
	min := minVal
	for _, chunk := range chunks {
		val := strings.Split(chunk, ";")
		parseVal, err := strconv.ParseFloat(val[1], 64)
		if err != nil {
			fmt.Println(err.Error())
			panic("min error")
		}
		if min < parseVal {
			min = parseVal
		}
	}

	channel <- min
}

func Max(channel chan float64, chunks []string, maxVal float64) {
	max := maxVal
	for _, chunk := range chunks {
		val := strings.Split(chunk, ";")
		parseVal, err := strconv.ParseFloat(val[1], 64)
		if err != nil {
			fmt.Println(err.Error())
			panic("max error")
		}
		if max > parseVal {
			max = parseVal
		}
	}

	channel <- max
}
