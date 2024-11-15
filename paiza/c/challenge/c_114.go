package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeC_114() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	N, _ := strconv.Atoi(input) // 単語数

	isSuccess := false
	preItem := ""
	afterItem := ""
	for i := 0; i < N; i++ {
		sc.Scan()
		input2 := sc.Text()

		if i == 0 {
			preItem = input2
			continue
		}

		// 最後の文字
		last := preItem[len(preItem)-1:]
		// 最初の文字
		first := input2[0:1]

		if last != first {
			afterItem = input2
			isSuccess = false
			break
		}

		// しりとり成功
		preItem = input2
		isSuccess = true
	}

	if isSuccess {
		fmt.Println("Yes")
		return
	}

	fmt.Println(preItem[len(preItem)-1:] + " " + afterItem[0:1])
}
