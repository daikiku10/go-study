package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 当選通知
func ChallengeC_159() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	// iArr := strings.Split(sc.Text(), " ")
	// N, _ := strconv.Atoi(iArr[0]) // 当選番号の数
	// K, _ := strconv.Atoi(iArr[1]) // 投票の数
	sc.Scan()
	nArr := strings.Split(sc.Text(), " ")
	sc.Scan()
	kArr := strings.Split(sc.Text(), " ")

	isNT := true

	for i, kv := range kArr {
		isT := false
		id := 0
		for _, nv := range nArr {
			if nv == kv {
				isT = true
				id = i + 1
				continue
			}
		}

		if isT {
			fmt.Println(id)
			isNT = false
		}
	}

	if isNT {
		fmt.Println(-1)
	}
}
