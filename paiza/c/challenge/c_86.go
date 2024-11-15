package challenge

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ChallengeC_86() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()

	// r := strings.NewReplacer("a", "", "i", "", "u", "", "e", "", "o", "", "A", "", "I", "", "U", "", "E", "", "O", "")
	// res := r.Replace(i)

	// 結果終了後、調べた結果のやり方
	r := regexp.MustCompile(`[aiueo]`)
	res := r.ReplaceAllString(i, "")
	fmt.Println(res)
}
