package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Challenge8_3() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	// 浮動小数点数に変換
	floatInput, _ := strconv.ParseFloat(input, 64)
	fmt.Printf("%.3f\n", floatInput)
}
