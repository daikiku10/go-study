package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_112() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()

	N, _ := strconv.Atoi(i) // 日数

	min := 0
	max := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()
		i2Arr := strings.Split(i2, " ")

		s, _ := strconv.Atoi(i2Arr[0]) // 出発時間
		f, _ := strconv.Atoi(i2Arr[1]) // フライト時間
		t, _ := strconv.Atoi(i2Arr[2]) // 到着時間

		time := s + f + (24 - t)

		if i == 0 {
			min = time
			max = time
			continue
		}

		if time < min {
			min = time
		}

		if time > max {
			max = time
		}
	}

	fmt.Println(min)
	fmt.Println(max)
}
