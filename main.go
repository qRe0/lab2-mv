package main

import (
	"fmt"
	"math"
)

func generateMatrix(n int) [][]float64 {
	A := make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				A[i][j] = 10 * math.Sqrt(float64((i+1)*(i+1)*(i+1)))
			} else {
				A[i][j] = 0.001 / math.Pow(float64(j+1), 0.3)
			}
		}
	}
	return A
}

func generateVector(n int) []float64 {
	b := make([]float64, n)
	for i := 0; i < n; i++ {
		b[i] = float64(i + 2)
	}
	return b
}

func solveSystem(n int, A [][]float64, b []float64, eps float64, kmax int, w float64) []float64 {
	x := make([]float64, n)
	xPrev := make([]float64, n)

	for k := 0; k < kmax; k++ {
		maxDiff := 0.0
		for i := 0; i < n; i++ {
			xPrev[i] = x[i]
			sum := 0.0
			for j := 0; j < n; j++ {
				if j != i {
					sum += A[i][j] * xPrev[j]
				}
			}
			x[i] = (1-w)*xPrev[i] + (w/b[i])*(b[i]-sum)
			diff := math.Abs(x[i] - xPrev[i])
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		if maxDiff < eps {
			fmt.Printf("Метод завершен на итерации %d с точностью %e\n", k, maxDiff)
			break
		}
	}
	return x
}

func printSolution(x []float64) {
	fmt.Println("Приближенное решение:")
	for i, value := range x {
		fmt.Printf("x[%d] = %f\n", i+1, value)
	}
}

func main() {
	n := 5
	eps := 1e-6
	kmax := 1000
	wValues := []float64{1, 0.5, 1.5}

	A := generateMatrix(n)
	b := generateVector(n)
	x := make([]float64, n)
	// xPrev := make([]float64, n)

	for _, w := range wValues {
		x = solveSystem(n, A, b, eps, kmax, w)
		fmt.Printf("Метод релаксации с w = %f\n", w)
		printSolution(x)
	}
}
