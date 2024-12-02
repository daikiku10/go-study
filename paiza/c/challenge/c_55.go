package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_55() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	G := sc.Text()

	count := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		si := sc.Text()

		if strings.Contains(si, G) {
			fmt.Println(si)
			count++
			continue
		}
	}

	if count == 0 {
		fmt.Println("None")
	}
}
