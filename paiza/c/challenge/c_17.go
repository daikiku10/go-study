package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_17() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(iArr[0]) // 親カード1つ目
	b, _ := strconv.Atoi(iArr[1]) // 親カード2つ目

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	res := make([]string, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		an, _ := strconv.Atoi(i2Arr[0]) // aの子カード
		bn, _ := strconv.Atoi(i2Arr[1]) // bの子カード

		// 同数ならbを比べる
		if a == an {
			if b < bn {
				res[i] = "High"
				continue
			}
			res[i] = "Low"
			continue
		}

		if a > an {
			res[i] = "High"
			continue
		}

		if a < an {
			res[i] = "Low"
			continue
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}

}
