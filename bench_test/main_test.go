package main

import (
	"fmt"
	"testing"
)

var num = 10

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Primenum(num)
	}
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 4000},
	{input: 20000},
}

func BenchmarkVariousPrime(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input range=%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Primenum(v.input)
			}
		})
	}
}

func BenchmarkSievePrime(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input range=%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PrimenumFaster(v.input)
			}
		})
	}
}
