package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChallengeC_84() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	len := len(input)
	fmt.Println(strings.Repeat("+", len+2))
	fmt.Println("+" + input + "+")
	fmt.Println(strings.Repeat("+", len+2))
}
