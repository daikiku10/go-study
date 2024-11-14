package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_130() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	N, _ := strconv.Atoi(input) // 問題数

	count := 0
	res := []int{}
	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()
		i2Arr := strings.Split(i2, " ")

		one := i2Arr[0]
		two := i2Arr[1]

		if !(one == "y" && two == "y") {
			count++
			res = append(res, i+1)
		}
	}

	fmt.Println(count)
	for _, v := range res {
		fmt.Println(v)
	}
}
