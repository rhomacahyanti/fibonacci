package main

import "fmt"

func main() {
	var n int64

	fmt.Print("Enter the number of terms : ")
	// fmt.Scan(&n)

	n = 5

	fibSeries := fibonacciSeries(n)

	fmt.Println("Fibonacci Series :", fibSeries)
}

func fibonacciSeries(n int64) (fibSeries []int64) {

	var t1, t2 = 0, 1

	for i := int64(1); i <= n; i++ {
		if i == 1 {
			fibSeries = append(fibSeries, int64(t1))
			continue
		}
		if i == 2 {
			fibSeries = append(fibSeries, int64(t2))
			continue
		}

		nextTerm := t1 + t2
		t1 = t2
		t2 = nextTerm

		fibSeries = append(fibSeries, int64(nextTerm))
	}

	return fibSeries
}
