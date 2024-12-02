package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_75() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0]) // 金額
	M, _ := strconv.Atoi(iArr[1]) // 乗車回数

	point := 0
	for i := 0; i < M; i++ {
		sc.Scan()
		fi, _ := strconv.Atoi(sc.Text()) // 運賃

		if fi > point {
			N = N - fi
			point = point + (fi / 10)
			f := fmt.Sprintf("%v %v", N, point)
			fmt.Println(f)
			continue
		}

		point = point - fi
		f := fmt.Sprintf("%v %v", N, point)
		fmt.Println(f)
	}

}
