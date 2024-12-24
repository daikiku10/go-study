package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 宿泊費と交通費
func ChallengeC_126() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	A, _ := strconv.Atoi(iArr[0]) // 片道料金
	B, _ := strconv.Atoi(iArr[1]) // 宿泊費
	N, _ := strconv.Atoi(iArr[2]) // インターンシップ回数

	nList := make([][]int, 0, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		xi, _ := strconv.Atoi(i2Arr[0]) // 初日
		yi, _ := strconv.Atoi(i2Arr[1]) // 最終日
		l := make([]int, 0, len(i2Arr))
		l = append(l, xi)
		l = append(l, yi)
		nList = append(nList, l)
	}

	res := A * 2 // 最初と最後の交通費

	for i, v := range nList {
		if i == 0 {
			continue
		}

		preEnd := nList[i-1][1]
		start := v[0]

		between := start - preEnd

		if (A * 2) > B*between {
			res = res + (B * between)
			continue
		}
		res = res + (A * 2)
	}
	fmt.Println(res)
}
