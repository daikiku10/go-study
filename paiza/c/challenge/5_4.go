package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Challenge5_4() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	input := sc.Text()
	inputArr := strings.Split(input, " ")

	fmt.Println(strings.Join(inputArr, ","))
}
