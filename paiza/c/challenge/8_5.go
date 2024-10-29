package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Challenge8_5() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputInt, _ := strconv.Atoi(input)

	for i := 0; i < inputInt; i++ {
		sc.Scan()
		t := sc.Text()
		tArr := strings.Split(t, " ")

		// 浮動小数点数に変換
		N, _ := strconv.ParseFloat(tArr[0], 64)
		M, _ := strconv.Atoi(tArr[1])
		fmt.Printf("%.*f\n", M, N)
	}
}
