package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeC_120() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // 花の数
	sc.Scan()
	S1 := sc.Text()
	sc.Scan()
	S2 := sc.Text()
	res := "No"
	v := S2
	for i := 0; i < N; i++ {
		last := v[len(v)-1:]
		h := fmt.Sprintf("%v%v", last, v[0:len(S2)-1])
		if S1 == h {
			res = "Yes"
			break
		}
		v = h
	}

	fmt.Println(res)
}
