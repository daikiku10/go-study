package challenge

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_150() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0]) // 家の数
	D, _ := strconv.Atoi(iArr[1]) // 距離

	sc.Scan()
	i2Arr := strings.Split(sc.Text(), " ")
	X, _ := strconv.Atoi(i2Arr[0]) // 座標x
	Y, _ := strconv.Atoi(i2Arr[1]) // 座標y

	count := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		i3Arr := strings.Split(sc.Text(), " ")
		xi, _ := strconv.Atoi(i3Arr[0]) // 相手座標x
		yi, _ := strconv.Atoi(i3Arr[1]) // 相手座標y

		ans := math.Abs(float64(X-xi)) + math.Abs(float64(Y-yi))

		if ans <= float64(D) {
			count++
		}
	}

	fmt.Println(count)

}
