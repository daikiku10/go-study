package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeD_246() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputInt, _ := strconv.Atoi(input)

	fmt.Println(inputInt + 10)
}
