package challenge

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// カードのスコア 不正解
func ChallengeC_105() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")

	// 文字列->数値 ソートする
	iIntArr := make([]int, 0, N)
	for _, v := range iArr {
		vInt, _ := strconv.Atoi(v)
		iIntArr = append(iIntArr, vInt)
	}
	sort.SliceStable(iIntArr, func(i, j int) bool { return iIntArr[i] < iIntArr[j] })

	resl := make([]int, 0, N)
	for i, v := range iIntArr {
		if i+1 == N {
			resl = append(resl, v)
			continue
		}

		if v+1 == iIntArr[i+1] {
			continue
		}
		resl = append(resl, v)
	}

	resCnt := 0
	for _, v := range resl {
		resCnt = resCnt + v
	}
	fmt.Println(resCnt)
}
