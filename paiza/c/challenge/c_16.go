package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChallengeC_16() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()

	r := strings.NewReplacer(
		"A", "4",
		"E", "3",
		"G", "6",
		"I", "1",
		"O", "0",
		"S", "5",
		"Z", "2",
	)
	res := r.Replace(i)
	fmt.Println(res)
}
