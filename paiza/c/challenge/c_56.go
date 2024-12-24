package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// テストの採点
func ChallengeC_56() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0]) // 人数
	M, _ := strconv.Atoi(iArr[1]) // 合格点

	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		ai, _ := strconv.Atoi(i2Arr[0]) // 点数
		bi, _ := strconv.Atoi(i2Arr[1]) // 欠席回数

		k := bi * 5
		t := ai - k
		if 0 > ai-k {
			t = 0
		}
		if M <= t {
			fmt.Println(i + 1)
		}
	}

}
