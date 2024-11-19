package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeC_102() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()

	M, _ := strconv.Atoi(i) // Aライブの日数
	aArr := make(map[int]int, M)

	for i := 0; i < M; i++ {
		sc.Scan()
		a := sc.Text()
		aInt, _ := strconv.Atoi(a)
		aArr[aInt] = aInt
	}

	sc.Scan()
	i2 := sc.Text()
	N, _ := strconv.Atoi(i2) // Bライブの日数
	bArr := make(map[int]int, N)

	for i := 0; i < N; i++ {
		sc.Scan()
		b := sc.Text()
		bInt, _ := strconv.Atoi(b)
		bArr[bInt] = bInt
	}

	months := make([]int, 31)
	for i := 0; i < 31; i++ {
		months[i] = i + 1
	}

	turn := "A"
	res := make([]string, 31)

	for i, v := range months {
		isA := false
		isB := false
		_, existA := aArr[v]
		if existA {
			isA = true
		}
		_, existB := bArr[v]
		if existB {
			isB = true
		}

		if !isA && !isB {
			res[i] = "x"
			continue
		}

		if isA && isB {
			res[i] = turn
			if turn == "A" {
				turn = "B"
				continue
			}
			turn = "A"
			continue
		}

		if isA {
			res[i] = "A"
			continue
		}

		if isB {
			res[i] = "B"
			continue
		}
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
