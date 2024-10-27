package challenge

import (
	"bufio"
	"fmt"
	"os"
)

func Challenge3_4() {
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		sc.Scan()
		if i < 9 {
			fmt.Print(sc.Text() + " ")
		} else {
			fmt.Println(sc.Text())
		}
	}
}
