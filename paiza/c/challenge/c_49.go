package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeC_49() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	N, _ := strconv.Atoi(i) // ログ数

	res := 0
	kai := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		f := sc.Text()
		fInt, _ := strconv.Atoi(f)

		if i == 0 {
			res = fInt - 1
			kai = fInt
			continue
		}

		if fInt-kai >= 0 {
			res = res + fInt - kai
			kai = fInt
			continue
		}

		res = res + kai - fInt
		kai = fInt
	}
	fmt.Println(res)
}
