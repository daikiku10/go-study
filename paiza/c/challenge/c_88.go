package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RPGでお買い物
func ChallengeC_88() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // 道具数
	sc.Scan()
	i2Arr := strings.Split(sc.Text(), " ") // 各道具単価リスト
	itemOf := make(map[int]int, N)
	for i, v := range i2Arr {
		vInt, _ := strconv.Atoi(v)
		itemOf[i+1] = vInt
	}

	sc.Scan()
	i3Arr := strings.Split(sc.Text(), " ")
	T, _ := strconv.Atoi(i3Arr[0]) // 所持金
	Q, _ := strconv.Atoi(i3Arr[1]) // 注文回数

	res := T
	for i := 0; i < Q; i++ {
		sc.Scan()
		i4Arr := strings.Split(sc.Text(), " ")
		x, _ := strconv.Atoi(i4Arr[0]) // 道具番号
		k, _ := strconv.Atoi(i4Arr[1]) // 個数

		buy := itemOf[x] * k

		if res >= buy {
			res = res - buy
			continue
		}
	}
	fmt.Println(res)
}
