package main

import (
	"fmt"
	"math"
)

var n int
var a, b float64

func my_f(x, eps float64) float64 {
	// x = (math.Abs(arctg(x, eps)) * 2) / math.Pi
	result := float64(1)
	k := 1
	step := result
	for math.Abs(step) >= eps {
		step *= -2 * float64(k) * x * (float64(3-2*k) / float64(4*math.Pow(float64(k), 2)*
			float64(1-2*k)))
		result += step
		k++
	}
	return result
}
func arctg(x, eps float64) float64 {
	new_x := x
	if math.Abs(x) > 1 {
		new_x = 1.0 / x
	}
	result := new_x
	k := 2
	step := new_x
	for math.Abs(step) >= eps {
		step *= -1 * math.Pow(new_x, 2) * float64(2*k-3) / float64(2*k-1)
		k++
		result += step
	}
	if math.Abs(x) > 1 {
		return math.Pi/2 - result
	}
	return result
}

func matrix_print(to_print [][]float64, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(to_print[i])
	}
}

func find_biggest(to_anal [][]float64, n int) [2]int {
	// var result []int
	var to_return [2]int
	to_return[0] = -1
	to_return[1] = 1000000000
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = float64(0)
	}
	for i := 0; i < n; i++ {
		for g := 0; g < n; g++ {
			result[g] += float64(to_anal[i][g])
		}
	}
	for i := 0; i < n; i++ {
		if result[i] > float64(to_return[0]) {
			to_return[0] = i
		}
		if result[i] < float64(to_return[1]) {
			to_return[1] = i
		}
	}
	return to_return
}

func odd_array(n int, digits []float64, eps float64) [][]float64 {
	result := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
	}
	k := 0
	m := 1
	result[n/2][n/2] = digits[0]
	now_x := n / 2
	now_y := n / 2
	count := 0
	for m < n-1 {
		if k == 0 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y+i][now_x] = my_f(digits[count], eps)
			}
			now_y += m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x+i] = my_f(digits[count], eps)
			}
			now_x += m
			k = 1
		} else if k == 1 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y-i][now_x] = my_f(digits[count], eps)
			}
			now_y -= m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x-i] = my_f(digits[count], eps)
			}
			now_x -= m
			k = 0
		}
		m++
	}
	for i := 1; i <= m; i++ {
		count++
		result[now_y-i][now_x] = my_f(digits[count], eps)
	}
	now_y -= m
	for i := 1; i <= m; i++ {
		count++
		result[now_y][now_x-i] = my_f(digits[count], eps)
	}
	now_x -= m
	for i := 1; i <= m; i++ {
		count++
		result[now_y+i][now_x] = my_f(digits[count], eps)
	}
	now_y += m
	return result
}
func honest_array(n int, digits []float64, eps float64) [][]float64 {
	result := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		result[i] = make([]float64, n)
	}
	k := 0
	m := 1
	result[n/2-1][n/2-1] = digits[0]
	now_x := n/2 - 1
	now_y := n/2 - 1
	count := 0
	for m < n-1 {
		if k == 0 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y+i][now_x] = my_f(digits[count], eps)
			}
			now_y += m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x+i] = my_f(digits[count], eps)
			}
			now_x += m
			k = 1
		} else if k == 1 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y-i][now_x] = my_f(digits[count], eps)
			}
			now_y -= m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x-i] = my_f(digits[count], eps)
			}
			now_x -= m
			k = 0
		}
		m++
	}
	// fmt.Println("second", m)
	// fmt.Println(now_y, now_x)
	for i := 1; i <= m; i++ {
		count++
		result[now_y+i][now_x] = my_f(digits[count], eps)
	}
	now_y += m
	for i := 1; i <= m; i++ {
		count++
		result[now_y][now_x+i] = my_f(digits[count], eps)
	}
	now_x += m
	for i := 1; i <= m; i++ {
		count++
		result[now_y-i][now_x] = my_f(digits[count], eps)
	}
	now_y -= m
	return result
}

//	func main() {
//		fmt.Println("Enter size of square array")
//		fmt.Scan(&n)
//		var result [][]int
//		if n%2 == 1 {
//			result = odd_array(n)
//		} else {
//			result = honest_array(n)
//		}
//		matrix_print(result, n)
//		indexes := find_biggest(result, n)
//		fmt.Println("biggest - ", indexes[0], "; least - ", indexes[1])
//	}
func main() {
	var result [][]float64
	eps := float64(0.00001)
	fmt.Println("Enter size of square matrix")
	fmt.Scan(&n)
	fmt.Println("Enter left and right border")
	fmt.Scan(&a, &b)
	digits := make([]float64, n*n)
	now_digit := a
	for i := 0; i < n*n; i++ {
		digits[i] = now_digit
		now_digit += (b - a) / float64(n-1)
	}
	fmt.Println(len(digits))
	if n%2 == 1 {
		result = odd_array(n, digits, eps)
	} else {
		result = honest_array(n, digits, eps)
	}
	matrix_print(result, n)
	indexes := find_biggest(result, n)
	fmt.Println("biggest - ", indexes[0], "; least - ", indexes[1])
}
