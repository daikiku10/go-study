package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_90() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	SArr := strings.Split(sc.Text(), "")

	telOf := map[int]int{
		0: 12,
		1: 3,
		2: 4,
		3: 5,
		4: 6,
		5: 7,
		6: 8,
		7: 9,
		8: 10,
		9: 11,
	}

	res := 0
	for i := 0; i < len(SArr); i++ {
		if SArr[i] == "-" {
			continue
		}

		sInt, _ := strconv.Atoi(SArr[i])

		res = res + (telOf[sInt] * 2)
	}
	fmt.Println(res)
}
