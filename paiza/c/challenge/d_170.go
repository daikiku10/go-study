package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeD_170() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputN := sc.Text()
	N, _ := strconv.Atoi(inputN)
	sc.Scan()
	inputM := sc.Text()
	M, _ := strconv.Atoi(inputM)

	fmt.Println(N * M)
}
