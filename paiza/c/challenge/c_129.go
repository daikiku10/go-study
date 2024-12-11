package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 工場の検品 不正解
func ChallengeC_129() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	// N, _ := strconv.Atoi(iArr[0]) // 寿司の種類の数
	M, _ := strconv.Atoi(iArr[1]) // パックされる寿司の個数

	model := make([]int, 0, M)
	for i := 0; i < M; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		model = append(model, ai)
	}

	to := make([]int, 0, M)
	for i := 0; i < M; i++ {
		sc.Scan()
		bi, _ := strconv.Atoi(sc.Text())
		to = append(to, bi)
	}

	for _, v := range to {
		for j, vj := range model {
			if v == vj {
				model = model[:j+copy(model[j:], model[j+1:])]
				break // ここを記入していなかった
			}
		}
	}

	if len(model) > 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}
