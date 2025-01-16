package main

import (
	"fmt"
)

var n int

func odd_array(n int) [][]int {
	result := make([][]int, n, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	k := 0
	m := 1
	result[n/2][n/2] = 1
	now_x := n / 2
	now_y := n / 2
	count := 1
	for m < n-1 {
		if k == 0 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y+i][now_x] = count
			}
			now_y += m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x+i] = count
			}
			now_x += m
			k = 1
		} else if k == 1 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y-i][now_x] = count
			}
			now_y -= m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x-i] = count
			}
			now_x -= m
			k = 0
		}
		m++
	}
	for i := 1; i <= m; i++ {
		count++
		result[now_y-i][now_x] = count
	}
	now_y -= m
	for i := 1; i <= m; i++ {
		count++
		result[now_y][now_x-i] = count
	}
	now_x -= m
	for i := 1; i <= m; i++ {
		count++
		result[now_y+i][now_x] = count
	}
	now_y += m
	return result
}
func honest_array(n int) [][]int {
	result := make([][]int, n, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	k := 0
	m := 1
	now_x := (n / 2) - 1
	now_y := (n / 2) - 1
	result[now_x][now_y] = 1
	count := 1
	for m < n-1 {
		if k == 0 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y+i][now_x] = count
			}
			now_y += m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x+i] = count
			}
			now_x += m
			k = 1
		} else if k == 1 {
			for i := 1; i <= m; i++ {
				count++
				result[now_y-i][now_x] = count
			}
			now_y -= m
			for i := 1; i <= m; i++ {
				count++
				result[now_y][now_x-i] = count
			}
			now_x -= m
			k = 0
		}
		m++
	}
	for i := 1; i <= m; i++ {
		count++
		result[now_y+i][now_x] = count
	}
	now_y += m
	for i := 1; i <= m; i++ {
		count++
		result[now_y][now_x+i] = count
	}
	now_x += m
	for i := 1; i <= m; i++ {
		count++
		result[now_y-i][now_x] = count
	}
	now_y -= m
	return result
}
func main() {
	fmt.Println("Enter size of square array")
	fmt.Scan(&n)
	if n%2 == 1 {
		fmt.Println(odd_array(n))
	} else {
		fmt.Println(honest_array(n))
	}
}
