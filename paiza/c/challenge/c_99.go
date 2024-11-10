package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_99() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")
	N, _ := strconv.Atoi(inputArr[0]) // 折り紙の枚数
	D, _ := strconv.Atoi(inputArr[1]) // 一片の長さ

	yoko := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		input2 := sc.Text()
		d, _ := strconv.Atoi(input2)
		yoko = yoko + (D - d)
	}

	fmt.Println(D * yoko)
}
