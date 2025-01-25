package challenge

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 降水量の計測
func ChallengeC_160() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")

	rain := 0
	for _, v := range iArr {
		vInt, _ := strconv.Atoi(v)
		rain = rain + vInt
	}

	ave := float64(rain) / float64(N)
	// 切り上げ
	res := math.Ceil(ave)
	fmt.Println(res)
}
