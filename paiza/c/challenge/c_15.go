package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ポイントカードの計算
func ChallengeC_15() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	point := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		di := i2Arr[0]                  // 日付
		pi, _ := strconv.Atoi(i2Arr[1]) // 購入金額

		if strings.Contains(di, "5") {
			p := float64(pi) * 0.05
			point = point + int(p)
			continue
		}

		if strings.Contains(di, "3") {
			p := float64(pi) * 0.03
			point = point + int(p)
			continue
		}

		p := float64(pi) * 0.01
		point = point + int(p)
		continue
	}
	fmt.Println(point)
}
