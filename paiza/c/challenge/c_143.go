package challenge

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ChallengeC_143() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	inputArr := strings.Split(input, "")

	resArr := make([]string, len(inputArr))
	for i := 0; i < len(inputArr); i++ {
		if i > 0 {
			if inputArr[i] == "-" && inputArr[i-1] == "-" {
				continue
			}
			resArr = append(resArr, inputArr[i])
			continue
		}
		resArr = append(resArr, inputArr[i])
	}
	fmt.Println(strings.Join(resArr, ""))

	// 終了後に色々調べた結果
	// 正規表現で1つ以上連続する「-」を1つの「-」に置き換える
	re := regexp.MustCompile(`-+`)
	output := re.ReplaceAllString(input, "-")

	fmt.Println(output)
}
