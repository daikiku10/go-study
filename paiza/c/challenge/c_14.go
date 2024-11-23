package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_14() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	iArr := strings.Split(i, " ")
	n, _ := strconv.Atoi(iArr[0]) // 箱の個数
	r, _ := strconv.Atoi(iArr[1]) // ボールの半径
	tyo := r * 2                  // 直径

	res := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		i2 := sc.Text()
		i2Arr := strings.Split(i2, " ")
		h, _ := strconv.Atoi(i2Arr[0]) // 高さ
		w, _ := strconv.Atoi(i2Arr[1]) // 幅
		d, _ := strconv.Atoi(i2Arr[2]) // 奥行き

		if h >= tyo && w >= tyo && d >= tyo {
			res = append(res, i+1)
		}
	}

	for _, v := range res {
		fmt.Println(v)
	}

}
