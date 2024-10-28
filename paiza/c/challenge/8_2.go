package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Challenge8_2() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fmt.Println(strings.TrimRight(sc.Text(), "0"))
}
