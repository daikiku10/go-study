package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 合言葉
func ChallengeC_147() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	sc.Scan()
	t := sc.Text()

	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")

	if len(sArr) != len(tArr) {
		fmt.Println("NO")
		return
	}

	if s == t {
		fmt.Println("NO")
		return
	}

	for _, sv := range sArr {
		for i, tv := range tArr {
			if sv == tv {
				tArr = tArr[:i+copy(tArr[i:], tArr[i+1:])]
				break
			}
		}
	}

	if len(tArr) > 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
