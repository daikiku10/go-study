package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ローソク足
func ChallengeC_22() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	resS := 0
	resE := 0
	resH := 0
	resL := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		h, _ := strconv.Atoi(i2Arr[2])
		l, _ := strconv.Atoi(i2Arr[3])

		if i == 0 {
			s, _ := strconv.Atoi(i2Arr[0])
			resS = s
			resL = l
		}

		if i == n-1 {
			e, _ := strconv.Atoi(i2Arr[1])
			resE = e
		}

		if resH < h {
			resH = h
		}

		if resL > l {
			resL = l
		}
	}
	f := fmt.Sprintf("%v %v %v %v", resS, resE, resH, resL)
	fmt.Println(f)
}
