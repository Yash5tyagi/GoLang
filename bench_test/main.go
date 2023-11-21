package main

import (
	"fmt"
	"math"
)

func Primenum(max int) []int {
	var prime []int

	for n := 2; n <= max; n++ {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, n)
		}
	}

	return prime
}

func PrimenumFaster(max int) []int {
	b := make([]bool, max+1)
	var prime []int
	for i := 2; i <= max; i++ {
		if b[i] {
			continue
		}
		prime = append(prime, i)
		for k := i * i; k <= max; k += i {
			b[k] = true
		}
	}
	return prime
}

func main() {
	// primeNum := Primenum(10)
	primeNum := PrimenumFaster(10)
	for _, num := range primeNum {
		fmt.Println(num)
	}
}
