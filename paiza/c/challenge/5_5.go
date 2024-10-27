package challenge

import (
	"bufio"
	"fmt"
	"os"
)

func Challenge5_5() {
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		sc.Scan()
		input := sc.Text()
		if i == 0 {
			fmt.Print(input + " " + "|")
		} else if i < 9 {
			fmt.Print(" " + input + " " + "|")
		} else {
			fmt.Println(" " + input)
		}
	}
}
