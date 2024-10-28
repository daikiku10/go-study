package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Challenge8_4() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	inputArr := strings.Split(input, " ")

	// 浮動小数点数に変換
	N, _ := strconv.ParseFloat(inputArr[0], 64)
	M, _ := strconv.Atoi(inputArr[1])
	fmt.Printf("%.*f\n", M, N)
}
