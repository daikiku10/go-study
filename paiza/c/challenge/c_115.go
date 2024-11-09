package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_115() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")
	N, _ := strconv.Atoi(inputArr[0])
	M, _ := strconv.Atoi(inputArr[1]) // 渋滞の定義

	reK := 0
	for i := 1; i < N; i++ {
		sc.Scan()
		input2 := sc.Text()
		k, _ := strconv.Atoi(input2)
		if M >= k {
			reK = reK + k
		}
	}

	fmt.Println(reK)
}
