package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type XorshiftState[T uint32 | uint64] struct {
	a, b, c, d T
}

// Random number generation with xorshift algorithm.
func xorshiftRand[T uint32 | uint64](state *XorshiftState[T]) T {
	x := state.a
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	state.a = x
	return x
}

// Calculates pi with Monte Carlo method.
func calcPiXorshift(n int) float64 {
	state := XorshiftState[uint64]{123456789, 362436069, 521288629, 88675123}
	insideCount := 0
	for i := 0; i < n; i++ {
		// convert uint64 -> float64 (0.0 <= x < 1.0)
		x := float64(xorshiftRand(&state)) / (1 << 64)
		y := float64(xorshiftRand(&state)) / (1 << 64)

		// check if (x, y) is inside the circle
		if x*x+y*y <= 1 {
			insideCount++
		}
	}
	return 4 * float64(insideCount) / float64(n)
}

func calcPiMathRand(n int) float64 {
	insideCount := 0
	for i := 0; i < n; i++ {
		x := float64(rand.Uint64()) / (1 << 31)
		y := float64(rand.Uint64()) / (1 << 31)
		if x*x+y*y <= 1 {
			insideCount++
		}
	}
	return 4 * float64(insideCount) / float64(n)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go [mathrand|xorshift]")
		return
	}

	mode := os.Args[1]
	calcPi := calcPiXorshift

	if mode == "mathrand" {
		calcPi = calcPiMathRand
		fmt.Println("use math/rand")
	} else if mode == "xorshift" {
		fmt.Println("use xorshift")
	} else {
		fmt.Println("usage: go run main.go [mathrand|xorshift]")
		return
	}

	start := time.Now()
	pi := calcPi(100000000)
	end := time.Now()
	fmt.Println(pi)
	fmt.Println(end.Sub(start))
}
