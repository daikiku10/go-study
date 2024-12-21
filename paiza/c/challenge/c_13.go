package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 嫌いな数字
func ChallengeC_13() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n := sc.Text()
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	isNone := true
	for i := 0; i < m; i++ {
		sc.Scan()
		r := sc.Text()

		if strings.Contains(r, n) {
			continue
		}
		rI, _ := strconv.Atoi(r)
		isNone = false
		fmt.Println(rI)
	}
	if isNone {
		fmt.Println("none")
	}

}
