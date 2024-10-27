package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Challenge7_5() {
	sc := bufio.NewScanner(os.Stdin)

	// var N int
	// var M int
	var NArr []int
	var MArr []int

	// 入力値の受け取り
	for i := 0; i < 3; i++ {
		sc.Scan()
		input := sc.Text()
		inputArr := strings.Split(input, " ")
		if i == 0 {
			// N, _ = strconv.Atoi(inputArr[0])
			// M, _ = strconv.Atoi(inputArr[1])
		}
		if i == 1 {
			for i := 0; i < len(inputArr); i++ {
				inputInt, _ := strconv.Atoi(inputArr[i])
				NArr = append(NArr, inputInt)
			}
		}
		if i == 2 {
			for i := 0; i < len(inputArr); i++ {
				inputInt, _ := strconv.Atoi(inputArr[i])
				MArr = append(MArr, inputInt)
			}
		}
	}

	// 出力
	count := 0
	for i := 0; i < len(MArr); i++ {
		for j := 0; j < MArr[i]; j++ {
			if j < MArr[i]-1 {
				fmt.Print(strconv.Itoa(NArr[count]) + " ")
				count++
			} else {
				fmt.Println(strconv.Itoa(NArr[count]))
				count++
			}
		}
	}

}
