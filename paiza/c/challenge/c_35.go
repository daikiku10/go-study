package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 試験の合格判定
func ChallengeC_35() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	cnt := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		iArr := strings.Split(sc.Text(), " ")
		sel := iArr[0]                 // 理系 or 文系
		ei, _ := strconv.Atoi(iArr[1]) // 英語
		mi, _ := strconv.Atoi(iArr[2]) // 数学
		si, _ := strconv.Atoi(iArr[3]) // 理科
		ji, _ := strconv.Atoi(iArr[4]) // 国語
		gi, _ := strconv.Atoi(iArr[5]) // 地理

		isZClear := false
		isSClear := false
		if 350 <= ei+mi+si+ji+gi {
			isZClear = true
		}

		if !isZClear {
			continue
		}

		if sel == "l" {
			if 160 <= ji+gi {
				isSClear = true
			}
		} else if sel == "s" {
			if 160 <= mi+si {
				isSClear = true
			}
		}

		if isSClear {
			cnt++
		}
	}

	fmt.Println(cnt)
}
