package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeC_141() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()

	N, _ := strconv.Atoi(i)

	h := make(map[string]int, N)

	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()
		_, exist := h[i2]
		if exist {
			h[i2]++
			continue
		}
		h[i2] = 1
	}

	res := ""
	resCount := 0
	for k, v := range h {
		if resCount < v {
			resCount = v
			res = k
		}
	}
	fmt.Println(res)
}
