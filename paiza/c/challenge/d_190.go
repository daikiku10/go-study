package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeD_190() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputN := sc.Text()
	N, _ := strconv.Atoi(inputN)
	fmt.Println(N * 10)
}
