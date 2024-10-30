package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Challenge9_1() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputInt, _ := strconv.Atoi(input)

	fmt.Printf("%3d", inputInt)
}
