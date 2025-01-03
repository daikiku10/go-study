package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 【コドモンコラボ問題】 虹が見える日
func ChallengeC_152() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	// D, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	wList := strings.Split(sc.Text(), " ")

	resL := make([]string, 0, len(wList))
	for i, v := range wList {
		if i == 0 {
			continue
		}
		// 前日が雨且つ今日が晴れ
		if wList[i-1] == "2" && v == "0" {
			resL = append(resL, strconv.Itoa(i))
		}
	}

	if len(resL) == 0 {
		fmt.Println(-1)
		return
	}

	fmt.Println(strings.Join(resL, " "))
}
