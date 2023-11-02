// задачи которые будут рассмотрены
// сложение матриц одинаковой размерности
// подсчет диагонали
// транспонирование
package main

import "fmt"

const rows = 3
const cols = 4

// сложение матриц одинакового размера
func sumMatrix(A [rows][cols]int, B [rows][cols]int) [rows][cols]int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] = A[i][j] + B[i][j]
		}
	}
	return A
}

// подсчет диагонали (суммы по дигонали) матрицы
func diagonalSum(A [rows][cols]int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i][i]
	}
	return sum
}

// траспонирование матрицы
func transpose(A [rows][cols]int) [cols][rows]int {
	var transposed [cols][rows]int
	for i := 0; i < len(A[0]); i++ {
		for j := 0; j < len(A); j++ {
			transposed[i][j] = A[j][i]
		}

	}
	return transposed
}

// использование одного цикла для прохода по матрице:
// не всегда, когда итерироние по матрице нужно использовать вложенные циклы
// существует "лайф хак"б рассмотрим вариант когда массив обходится с помощью одного индекса
func printInOneCycle(m [rows][cols]int) {
	for i := 0; i < rows*cols; i++ { //rows*cols позволяет рассмотреть матрицу как один большой массив
		row := i / cols //индекс строки
		col := i % cols //индекс колонок
		fmt.Println(m[row][col])
	}
}

func main() {
	matrix := [rows][cols]int{
		{10, 10, 10, 10},
		{10, 20, 10, 20},
		{-10, -20, -10, -10},
	}
	fmt.Println(sumMatrix(matrix, matrix))
	fmt.Println(diagonalSum(matrix))
	fmt.Println(transpose(matrix))
	printInOneCycle(matrix)
}
