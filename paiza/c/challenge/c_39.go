package challenge

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ChallengeC_39() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	iArr := strings.Split(i, "+")

	pattern1 := `(/+)`
	pattern2 := `(<+)`

	count := 0

	for _, v := range iArr {
		r, _ := regexp.Compile(pattern1)
		matches := r.FindAllString(v, -1)
		r2, _ := regexp.Compile(pattern2)
		matches2 := r2.FindAllString(v, -1)

		it := 0
		zu := 0
		if len(matches) == 1 {
			it = len(matches[0])
		}
		if len(matches2) == 1 {
			zu = len(matches2[0]) * 10
		}
		count = count + zu + it
	}
	fmt.Println(count)
}
