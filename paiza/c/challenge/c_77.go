package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 【30万人記念問題】 レポートの評価
func ChallengeC_77() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(iArr[0]) // 学生数
	n, _ := strconv.Atoi(iArr[1]) // レポート問題数

	for i := 0; i < k; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		di, _ := strconv.Atoi(i2Arr[0]) // 提出日
		ai, _ := strconv.Atoi(i2Arr[1]) // 正解した問題数

		aif := float64(ai)
		nf := float64(n)
		t := (aif / nf) * 100

		st := int(t)

		if di > 0 {
			if di > 9 {
				st = 0
			} else {
				// 8割得点
				st = int(t * 0.8)
			}
		}

		if st >= 80 {
			fmt.Println("A")
		} else if st >= 70 {
			fmt.Println("B")
		} else if st >= 60 {
			fmt.Println("C")
		} else {
			fmt.Println("D")
		}
	}
}
