package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_123() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	N, _ := strconv.Atoi(i) // 人数

	// それぞれの年齢
	pList := make([]struct {
		Age   int
		Count int
	}, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()
		Y, _ := strconv.Atoi(i2)
		pList[i] = struct {
			Age   int
			Count int
		}{Age: Y, Count: 0}
	}

	sc.Scan()
	i3 := sc.Text()
	M, _ := strconv.Atoi(i3) // 命令の数

	for i := 0; i < M; i++ {
		sc.Scan()
		i4 := sc.Text()
		i4Arr := strings.Split(i4, " ")
		A, _ := strconv.Atoi(i4Arr[0])
		B, _ := strconv.Atoi(i4Arr[1])
		C, _ := strconv.Atoi(i4Arr[2])

		for f := A; f <= B; f++ {
			pList[f-1] = struct {
				Age   int
				Count int
			}{
				Age:   pList[f-1].Age,
				Count: pList[f-1].Count + C,
			}
		}
	}
	for _, v := range pList {
		if v.Count >= v.Age {
			fmt.Println(v.Age)
			continue
		}
		fmt.Println(v.Count)
	}
}
