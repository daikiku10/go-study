package challenge

import (
	"bufio"
	"fmt"
	"os"
)

func Challenge5_2() {
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < 3; i++ {
		sc.Scan()

		if i < 2 {
			fmt.Print(sc.Text() + "|")
		} else {
			fmt.Println(sc.Text())
		}
	}
}
