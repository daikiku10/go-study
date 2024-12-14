package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ドーナツ 不正解
func ChallengeB_138() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	H, _ := strconv.Atoi(iArr[0])
	W, _ := strconv.Atoi(iArr[1])

	siL := make([][]string, 0, H)
	for i := 0; i < H; i++ {
		sc.Scan()
		si := sc.Text() // # or .
		siArr := strings.Split(si, "")
		siL = append(siL, siArr)
	}

	cnt := 0
	for i, vi := range siL {
		for j, vj := range vi {
			if (j == 0 && vj == ".") || (j == W-1 && vj == ".") { // W-1をH-1にしていた
				continue
			}

			if vj == "#" {
				continue
			}

			// 上としたの場合も除外
			if i == 0 || i == H-1 {
				continue
			}

			// . のみがヒットする
			// 周りが#であるか確認
			// 左右
			if !(vi[j-1] == "#" && vi[j+1] == "#") {
				continue
			}
			// 上の列
			if !(siL[i-1][j-1] == "#" && siL[i-1][j] == "#" && siL[i-1][j+1] == "#") {
				continue
			}

			// 下の列
			if !(siL[i+1][j-1] == "#" && siL[i+1][j] == "#" && siL[i+1][j+1] == "#") {
				continue
			}

			cnt++
		}
	}
	fmt.Println(cnt)
}
