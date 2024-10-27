package challenge

import (
	"bufio"
	"fmt"
	"os"
)

func Challenge3_3() {
	sc := bufio.NewScanner(os.Stdin)

	for i := 0; i < 2; i++ {
		sc.Scan()
		fmt.Println(sc.Text())
	}
}
