package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 【2021年Xmas問題】 ラッキーデイ
func ChallengeC_101() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	X := sc.Text()

	cnt := 0
	for i := 0; i < 365; i++ {
		if strings.Contains(strconv.Itoa(i), X) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
