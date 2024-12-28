package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ロボット芸人
func ChallengeC_103() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0]) // 作動する時間
	M, _ := strconv.Atoi(iArr[1]) // 規則数

	kl := make([]struct {
		aj int
		bj string
	}, 0, M)
	for i := 0; i < M; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		aj, _ := strconv.Atoi(i2Arr[0]) // 動作の規則
		bj := i2Arr[1]                  // 文字列
		kl = append(kl, struct {
			aj int
			bj string
		}{
			aj: aj,
			bj: bj,
		})
	}

	for i := 0; i < N; i++ {
		res := ""
		isNone := true
		for _, v := range kl {
			if (i+1)%v.aj == 0 {
				if res == "" {
					res = v.bj
					isNone = false
					continue
				}
				res = fmt.Sprintf("%v %v", res, v.bj)
				isNone = false
				continue
			}
		}
		if isNone {
			fmt.Println(i + 1)
		} else {
			fmt.Println(res)
		}
	}

}
