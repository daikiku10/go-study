package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Challenge3_5() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	input := sc.Text()

	inputArr := strings.Split(input, " ")

	for _, v := range inputArr {
		fmt.Println(v)
	}
}
