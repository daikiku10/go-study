package challenge

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 安息の地を求めて
func ChallengeC_10() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(iArr[0]) // x座標
	b, _ := strconv.Atoi(iArr[1]) // y座標
	R, _ := strconv.Atoi(iArr[2]) // 騒音の大きさ

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // 木陰の数

	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		xi, _ := strconv.Atoi(i2Arr[0])
		yi, _ := strconv.Atoi(i2Arr[1])

		x := math.Pow(float64((xi - a)), 2)
		y := math.Pow(float64((yi - b)), 2)
		rb := math.Pow(float64(R), 2)

		if x+y >= rb {
			fmt.Println("silent")
			continue
		}
		fmt.Println("noisy")
	}

}
