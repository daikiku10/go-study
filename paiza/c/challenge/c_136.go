package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ダイエットの連続記録
func ChallengeC_136() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) //

	resNoD := 0 // 最大数
	resD := 0   // 最大数
	noD := 0    // 連続記録
	d := 0      // 連続記録
	preT := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		xi, _ := strconv.Atoi(sc.Text()) // i日目の体重

		// 初日
		if i == 0 {
			preT = xi
			continue
		}

		if preT > xi {
			d++
			noD = 0
			if resD < d {
				resD = d
			}
			preT = xi
			continue
		}

		noD++
		d = 0
		if resNoD < noD {
			resNoD = noD
		}

		preT = xi
	}

	resStr := fmt.Sprintf("%v %v", resD, resNoD)
	fmt.Println(resStr)
}
