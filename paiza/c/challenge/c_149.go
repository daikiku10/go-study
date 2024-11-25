package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChallengeC_149() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	S := sc.Text() // 26文字
	SArr := strings.Split(S, "")
	sc.Scan()
	T := sc.Text()
	TArr := strings.Split(T, "")

	res := make([]string, len(TArr))
	for i, v := range TArr {
		for _, sv := range SArr {
			if v == strings.ToLower(sv) || v == strings.ToUpper(sv) {
				// svが小文字である
				if sv == strings.ToLower(v) {
					res[i] = strings.ToLower(v)
				}
				if sv == strings.ToUpper(v) {
					res[i] = strings.ToUpper(v)
				}
			}
		}
	}

	resStr := strings.Join(res, "")
	fmt.Println(resStr)
}
