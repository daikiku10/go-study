package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_154() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")
	N, _ := strconv.Atoi(inputArr[0]) // 個数
	L, _ := strconv.Atoi(inputArr[1]) // 半額適用金額
	sc.Scan()
	input2 := sc.Text()
	input2Arr := strings.Split(input2, " ")

	item := 0 // 半額対象商品
	goukei := 0
	for i := 0; i < N; i++ {
		enStr := input2Arr[i]
		en, _ := strconv.Atoi(enStr)
		if en >= L && en > item {
			item = en
		}
		goukei = goukei + en
	}

	fmt.Println(goukei - (item / 2))
}
