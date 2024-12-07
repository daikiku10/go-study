package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 株の売買
func ChallengeC_78() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0])
	c1, _ := strconv.Atoi(iArr[1])
	c2, _ := strconv.Atoi(iArr[2])

	k := 0
	kCnt := 0
	b := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		pi, _ := strconv.Atoi(sc.Text())

		// N日 すべて売る
		if i == N-1 {
			b = b + (pi*kCnt - k)
			continue
		}

		// 持ち株を買う
		if c1 >= pi {
			k = k + pi
			kCnt++
			fmt.Println(k)
			fmt.Println()
			continue
		}

		// 持ち株をすべて売る
		if c2 <= pi {
			b = b + (pi*kCnt - k)
			k = 0
			kCnt = 0
			continue
		}
	}
	fmt.Println(b)
}
