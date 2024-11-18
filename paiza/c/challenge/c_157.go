package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChallengeC_157() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	sc.Scan()
	i2 := sc.Text()
	i2Arr := strings.Split(i2, " ")

	buke := []string{}
	for _, v := range i2Arr {
		fount := false
		for _, v2 := range buke {
			if v == v2 {
				fount = true
				break
			}
		}

		if !fount {
			buke = append(buke, v)
		}
	}
	fmt.Println(len(buke))
}
