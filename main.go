package main

import (
	"fmt"
	"math"
	"time"
)

const (
	eps  = 1e-6 // Порог точности
	kmax = 1000 // Максимальное количество итераций
	N    = 7    // Размерность матрицы
)

// Функция для генерации матрицы A
func GenerateMatrixA(n int) [][]float64 {
	A := make([][]float64, n)
	for i := range A {
		A[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		A[i][i] = 10.0 * math.Sqrt(math.Pow(float64(i+1), 3))
		for j := 0; j < n; j++ {
			if i != j {
				A[i][j] = 0.001 / math.Pow(float64(j+1), 0.3)
			}
		}
	}

	return A
}

// Функция для генерации вектора b
func GenerateVectorB(A [][]float64, n int) []float64 {
	xStar := make([]float64, n)
	for i := 0; i < n; i++ {
		xStar[i] = float64(i + 3) // Создаем вектор x* = (3, 4, 5, ...)
	}

	b := make([]float64, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b[i] += A[i][j] * xStar[j]
		}
	}

	return b
}

func IsStrictlyDiagonallyDominant(A [][]float64) bool {
	n := len(A)

	for i := 0; i < n; i++ {
		diagonal := math.Abs(A[i][i])
		sum := 0.0

		for j := 0; j < n; j++ {
			if i != j {
				sum += math.Abs(A[i][j])
			}
		}

		if diagonal <= sum {
			return false // Не строго диагонально доминирующая матрица
		}
	}

	return true // Строго диагонально доминирующая матрица
}

// Функция для вычисления суммы элементов A * X, исключая элемент с индексом excludeIndex
func SumAx(A [][]float64, X []float64, excludeIndex int) float64 {
	sum := 0.0
	for i := 0; i < len(A); i++ {
		if i != excludeIndex {
			sum += A[excludeIndex][i] * X[i]
		}
	}
	return sum
}

// Метод Якоби
func JacobiMethod(A [][]float64, b []float64, n int) ([]float64, time.Duration) {
	startTime := time.Now()
	X := make([]float64, n) // Начальное приближение

	k := 0
	for k < kmax {
		XNew := make([]float64, n)
		for i := 0; i < n; i++ {
			XNew[i] = (b[i] - SumAx(A, X, i)) / A[i][i]
		}

		maxDiff := 0.0
		for i := 0; i < n; i++ {
			maxDiff = math.Max(maxDiff, math.Abs(XNew[i]-X[i]))
		}

		X = XNew
		k++

		if maxDiff < eps {
			elapsedTime := time.Since(startTime)
			fmt.Printf("Jacobi method: the required accuracy per iteration has been achieved on iter %d\n", k)
			return X, elapsedTime
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Jacobi method: the maximum number of iterations reached (%d)\n", kmax)
	return X, elapsedTime
}

// Метод Гаусса-Зейделя (w = 1)
func GaussSeidelMethod(A [][]float64, b []float64, n int) ([]float64, time.Duration) {
	startTime := time.Now()
	X := make([]float64, n) // Начальное приближение

	k := 0
	for k < kmax {
		for i := 0; i < n; i++ {
			sum := SumAx(A, X, i)
			X[i] = (b[i] - sum) / A[i][i]
		}

		maxDiff := 0.0
		for i := 0; i < n; i++ {
			maxDiff = math.Max(maxDiff, math.Abs(X[i]))
		}

		k++

		if maxDiff < eps {
			elapsedTime := time.Since(startTime)
			fmt.Printf("Gauss-Seidel method: the required accuracy per iteration is achieved on iter %d\n", k)
			return X, elapsedTime
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Gauss-Seidel method: the maximum number of iterations reached (%d)\n", kmax)
	return X, elapsedTime
}

// Метод релаксации
func RelaxationMethod(A [][]float64, b []float64, w float64, n int) []float64 {
	X := make([]float64, n) // Начальное приближение

	k := 0
	for k < kmax {
		XNew := make([]float64, n)
		for i := 0; i < n; i++ {
			sum1 := SumAx(A, XNew, i)
			sum2 := SumAx(A, X, i)
			XNew[i] = X[i] + w*((b[i]-sum1)/A[i][i]-X[i]+sum2)
		}

		maxDiff := 0.0
		for i := 0; i < n; i++ {
			maxDiff = math.Max(maxDiff, math.Abs(XNew[i]-X[i]))
		}

		X = XNew
		k++

		if maxDiff < eps {
			fmt.Printf("Relaxation method (w = %.2f): the required accuracy per iteration has been achieved on iter %d\n", w, k)
			return X
		}
	}

	fmt.Printf("Relaxation method (w = %.2f): the maximum number of iterations has been reached (%d)\n", w, kmax)
	return X
}

func printMatrix(matrix [][]float64, b []float64, n int) {
	maxRows := 8
	maxCols := 8

	// fmt.Printf("%v\n", b)

	for i := 0; i < min(n, maxRows); i++ {
		for j := 0; j < min(n, maxCols); j++ {
			fmt.Printf("%10.6f ", matrix[i][j])
		}

		if n > maxCols {
			fmt.Print(" ...")
		}

		fmt.Printf(" |     %10.6f\n", b[i])
	}

	if n > maxRows {
		fmt.Print("...")
		for j := 1; j < maxCols; j++ {
			fmt.Print(" ...")
		}

		fmt.Print(" | ...")
	}

	fmt.Println()
}

// Функция для вычисления кубической нормы разности между двумя векторами
func CubicNorm(x1, x2 []float64) float64 {
	n := len(x1)
	norm := 0.0
	for i := 0; i < n; i++ {
		norm += math.Pow(math.Abs(x1[i]-x2[i]), 3)
	}
	return math.Cbrt(norm)
}

// Функция для вычисления относительной погрешности
func RelativeError(xTrue, xApprox []float64) float64 {
	n := len(xTrue)
	maxAbsXTrue := 0.0
	for i := 0; i < n; i++ {
		maxAbsXTrue = math.Max(maxAbsXTrue, math.Abs(xTrue[i]))
	}

	error := 0.0
	for i := 0; i < n; i++ {
		error = math.Max(error, math.Abs(xTrue[i]-xApprox[i])/maxAbsXTrue)
	}
	return error
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Задайте матрицу A и вектор b
	A := GenerateMatrixA(N)
	b := GenerateVectorB(A, N)

	// Вывод матрицы A и вектора b
	fmt.Println("Matrix A:")
	printMatrix(A, b, N)

	isStrictlyDiagonallyDominant := IsStrictlyDiagonallyDominant(A)
	if isStrictlyDiagonallyDominant {
		fmt.Println("The matrix is strictly diagonally dominant.")
	} else {
		fmt.Println("The matrix is not strictly diagonally dominant.")
	}
	fmt.Println()
	fmt.Println()

	// Решение системы методами
	solutionJacobi, timeJacobi := JacobiMethod(A, b, N)
	solutionGaussSeidel, timeGS := GaussSeidelMethod(A, b, N)
	solutionRelaxation1 := RelaxationMethod(A, b, 0.5, N)
	solutionRelaxation2 := RelaxationMethod(A, b, 1.5, N)

	// Вычисление точного решения (x*)
	xTrue := make([]float64, len(b))
	copy(xTrue, b)
	fmt.Println()
	fmt.Println()

	// Вывод точного решения
	fmt.Print("Precise solution (x*): ")
	for _, x := range xTrue {
		fmt.Printf("%10.6f ", x)
	}
	fmt.Println()

	// Вывод приближенных решений
	fmt.Print("Approximate solution by the Jacobi method: ")
	for _, x := range solutionJacobi {
		fmt.Printf("%10.6f ", x)
	}
	fmt.Println()

	fmt.Print("Approximate solution by the Gauss-Seidel method: ")
	for _, x := range solutionGaussSeidel {
		fmt.Printf("%10.6f ", x)
	}
	fmt.Println()

	fmt.Print("Approximate solution by the relaxation method (w = 0.5): ")
	for _, x := range solutionRelaxation1 {
		fmt.Printf("%10.6f ", x)
	}
	fmt.Println()

	fmt.Print("Approximate solution by the relaxation method (w = 1.5): ")
	for _, x := range solutionRelaxation2 {
		fmt.Printf("%10.6f ", x)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// Вывод времени работы методов
	fmt.Printf("Jacobi method: %v\n", timeJacobi)
	fmt.Printf("Gauss-Seidel method: %v\n", timeGS)
	fmt.Println()
	fmt.Println()

	// Вычисление кубической нормы разности и относительной погрешности
	normJacobi := CubicNorm(xTrue, solutionJacobi)
	normGaussSeidel := CubicNorm(xTrue, solutionGaussSeidel)
	normRelaxation1 := CubicNorm(xTrue, solutionRelaxation1)
	normRelaxation2 := CubicNorm(xTrue, solutionRelaxation2)

	errorJacobi := RelativeError(xTrue, solutionJacobi)
	errorGaussSeidel := RelativeError(xTrue, solutionGaussSeidel)
	errorRelaxation1 := RelativeError(xTrue, solutionRelaxation1)
	errorRelaxation2 := RelativeError(xTrue, solutionRelaxation2)

	// Вывод результатов
	fmt.Printf("Cubic norm of difference (Jacobi method): %10.6f\n", normJacobi)
	fmt.Printf("Relative error (Jacobi method): %10.6f\n", errorJacobi)

	fmt.Printf("Cubic norm of difference (Gauss-Seidel method): %10.6f\n", normGaussSeidel)
	fmt.Printf("Relative error (Gauss-Seidel method): %10.6f\n", errorGaussSeidel)

	fmt.Printf("Cubic norm of difference (relaxation method, w = 0.5): %10.6f\n", normRelaxation1)
	fmt.Printf("Relative error (relaxation method, w = 0.5): %10.6f\n", errorRelaxation1)

	fmt.Printf("Cubic norm of difference (relaxation method, w = 1.5): %10.6f\n", normRelaxation2)
	fmt.Printf("Relative error (relaxation method, w = 1.5): %10.6f\n", errorRelaxation2)
}
