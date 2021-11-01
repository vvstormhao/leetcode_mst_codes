package main


import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	ret := make([][]int, n, n)
	for i:=0;i<n;i++ {
		ret[i] = make([]int, n, n)
	}

	var (
		startx int = 0
		starty int = 0
	)

	loop := n/2 // 需要循环几次
	count := 1
	offset := 1
	mid := n / 2;

	for;loop > 0;loop-- {
		i := startx
		j := starty

		for ;j < starty + n - offset;j++ {
			ret[i][j] = count
			count++
		}

		for ;i < startx + n - offset; i++ {
			ret[i][j] = count
			count++
		}

		for ;j > startx; j-- {
			ret[i][j] = count
			count++
		}

		for ;i > starty; i-- {
			ret[i][j] = count
			count++
		}

		startx++
		starty++
		offset += 2
	}

	if n % 2 != 0 {
		ret[mid][mid] = count
	}

	return ret
}

func main() {
	ret := generateMatrix(4)
	for _, v := range ret {
		fmt.Println(v)
	}
}