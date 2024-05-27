package main

import (
	"fmt"
	"os"

	"math-skills/stats"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]
	data, err := stats.DataReader(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Average: %.f \n", stats.Average(data))
	fmt.Printf("Median: %.f \n", stats.Median(data))
	fmt.Printf("Variance: %.f \n", stats.Variance(data))
	fmt.Printf("Standard Deviation: %.f \n", stats.StandardDeviation(data))
}
