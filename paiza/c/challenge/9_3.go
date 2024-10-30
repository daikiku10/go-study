package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Challenge9_3() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	inputInt, _ := strconv.Atoi(input)

	for i := 0; i < inputInt; i++ {
		sc.Scan()
		t := sc.Text()
		tInt, _ := strconv.Atoi(t)
		fmt.Printf("%3d\n", tInt)
	}
}
