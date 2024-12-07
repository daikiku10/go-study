package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ストラックアウト
func ChallengeC_89() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(iArr[0]) // 高さ
	W, _ := strconv.Atoi(iArr[1]) // 幅

	out := make([]string, 0, H*W)
	for i := 0; i < H; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), "")
		out = append(out, i2Arr...)
	}

	mato := make([]int, 0, H*W)
	for i := 0; i < H; i++ {
		sc.Scan()
		i3Arr := strings.Split(sc.Text(), " ")
		for _, v := range i3Arr {
			i, _ := strconv.Atoi(v)
			mato = append(mato, i)
		}

	}
	cnt := 0
	for i, v := range out {
		if v == "o" {
			cnt = cnt + mato[i]
		}
	}
	fmt.Println(cnt)
}
