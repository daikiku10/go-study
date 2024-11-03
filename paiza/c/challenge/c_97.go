package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_97() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")
	N := inputArr[0]
	NInt, _ := strconv.Atoi(N)
	X := inputArr[1]
	XInt, _ := strconv.Atoi(X)
	Y := inputArr[2]
	YInt, _ := strconv.Atoi(Y)

	for i := 1; i <= NInt; i++ {
		if i%XInt == 0 && i%YInt == 0 {
			fmt.Println("AB")
			continue
		}
		if i%XInt == 0 {
			fmt.Println("A")
			continue
		}
		if i%YInt == 0 {
			fmt.Println("B")
			continue
		}
		fmt.Println("N")
	}
}
