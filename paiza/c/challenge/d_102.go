package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeD_102() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputInt, _ := strconv.Atoi(input)

	fmt.Println(100 + (inputInt * 10))
}
