package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeD_162() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputInt, _ := strconv.Atoi(input)

	fmt.Println(inputInt * 150)
}