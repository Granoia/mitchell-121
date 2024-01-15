package hw2

import "math"

// IsPrime returns true if the argument is prime.
func IsPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i < x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func SumCubes(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func Fibonacci3(n int) int {
	x := 1
	y := 1
	z := 1
	for i := 3; i < n; i++ {
		nextNumber := x + y + z
		x = y
		y = z
		z = nextNumber
	}
	return z
}

func poly(x, y float64) float64 {
	return -math.Pow(x, 4) + 3*math.Pow(x, 2) - math.Pow(y, 4) + 5*math.Pow(y, 2)
}
func Polymax(xmin, xmax, ymin, ymax int) int {
	max := math.MinInt
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			n := int(poly(float64(x), float64(y)))
			if n > max {
				max = n
			}
		}
	}
	return max
}

func Collatz(n int) int {
	c := 1
	for n != 1 {
		c++
		if n%2 == 1 {
			n = 3*n + 1
		} else {
			n = n / 2
		}
	}
	return c
}
