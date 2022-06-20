package main

import (
	"bufio"
	"math"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func avarage(nums []int) float64 {

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	a := len(nums)
	total := float64(sum) / float64(a)
	return (total)
}

func median(nums []int) int {

	sort.Ints(nums)
	numbers := len(nums)

	var median int

	if numbers == 0 {
		return 0
	} else if numbers%2 == 0 {
		median = numbers / 2
		median -= 1
		sum := nums[median] + nums[median+1]
		return sum / 2
	} else {
		return nums[len(nums)/2]
	}
}

func variance(nums []int) float64 {

	var dev []float64
	var res float64
	var sq []float64
	var sum float64

	//find the mean
	mean := avarage(nums)

	// find each score’s deviation from the mean
	for _, num := range nums {
		s := float64(num) - float64(mean)
		dev = append(dev, float64(s))
	}

	// square each deviation from the mean
	for i := 0; i < len(dev); i++ {
		res = math.Pow(dev[i], 2)
		sq = append(sq, res)
	}

	//	find the sum of squares
	sum = 0

	for i := 0; i < len(sq); i++ {
		sum += sq[i]
	}

	//divide the sum of squares by lenght of array of numbers -1
	numI := len(nums) - 1
	s := sum / float64(numI)
	
	return s
	}

func standardDeviation(nums []int) float64 {

	// find the variance
	v := variance(nums)

	//find the square root of the variance
	res := math.Sqrt(v)

	return res

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints, scanner.Err()
}

func main() {

	file, err := readLines("data.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	a := avarage(file)
	m := median(file)
	v := variance(file)
	s := standardDeviation(file)
	
	fmt.Println("Avarage:", math.Round(a))
	fmt.Println("Median:", (m))
	fmt.Println("Variance:", math.Round(v))
	fmt.Println("Standard Deviation:", math.Round(s))

}
