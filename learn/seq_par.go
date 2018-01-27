package main

import (
	"fmt"
	"math"
	"time"
)

func add(s []float64, c chan float64) {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := make([]float64, 100000000)
	c := make(chan float64, 3)
	for i := range s {
		s[i] = math.Sqrt(float64(i))
	}
	start := time.Now()
	//	add(s, c)
	//	end := time.Since(start)
	//	fmt.Println(<-c, end)
	go add(s[:len(s)/3], c)
	go add(s[len(s)/3:2*len(s)/3], c)
	go add(s[2*len(s)/3:], c)
	res1 := <-c
	//end1 := time.Since(start) - end
	//end2 := time.Since(start) - end1 - end
	res2 := <-c
	res3 := res1 + res2
	//fmt.Println(res3, end1, end2, end1+end2, time.Since(start)-end)
	fmt.Println(res3, time.Since(start))

}
