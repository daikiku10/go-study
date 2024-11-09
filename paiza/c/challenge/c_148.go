package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_148() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")
	N, _ := strconv.Atoi(inputArr[0])
	L, _ := strconv.Atoi(inputArr[1])

	level := L
	for i := 0; i < N; i++ {
		sc.Scan()
		input2 := sc.Text()
		tekiL, _ := strconv.Atoi(input2)

		if level > tekiL { // 勝ち
			point := tekiL / 2
			iP := int(point) // Go言語の場合、整数同士の割り算で小数点以下が切り捨てられるため、ここは不要
			level = level + iP
		}
		if tekiL > level { // 負け
			point := level / 2
			iP := int(point) // Go言語の場合、整数同士の割り算で小数点以下が切り捨てられるため、ここは不要
			level = iP
		}
	}
	fmt.Println(level)
}
