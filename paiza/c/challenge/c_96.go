package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 夏休み 不正解(特殊なケースの場合の1問を間違えた)
func ChallengeC_96() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // メンバーの人数(子供は除く)

	list := make([][]int, 0, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		iArr := strings.Split(sc.Text(), " ")
		start, _ := strconv.Atoi(iArr[0])
		end, _ := strconv.Atoi(iArr[1])

		l := make([]int, 0, end-start+1)
		for i := start; i <= end; i++ {
			l = append(l, i)
		}
		list = append(list, l)
	}

	for i, vl := range list {
		if i == 0 {
			continue
		}
		ok := false
		for _, prev := range list[i-1] {
			for _, v := range vl {
				if prev == v {
					ok = true
				}
			}
		}

		if ok {
			continue
		}
		fmt.Println("NG")
		return
	}
	fmt.Println("OK")
}
