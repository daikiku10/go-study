package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_117() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")
	N, _ := strconv.Atoi(inputArr[0]) // 店舗数
	M, _ := strconv.Atoi(inputArr[1]) // 営業する月数

	sc.Scan()
	input2 := sc.Text()
	input2Arr := strings.Split(input2, " ")
	A, _ := strconv.Atoi(input2Arr[0]) // 建築費用
	B, _ := strconv.Atoi(input2Arr[1]) // 人件費
	C, _ := strconv.Atoi(input2Arr[2]) // 利益

	count := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		input3 := sc.Text()
		H, _ := strconv.Atoi(input3) // 杯数

		tes := C*H - A - B*M

		if tes < 0 {
			count++
		}
	}

	fmt.Println(count)

}
