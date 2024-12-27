package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 【50万人記念問題】観光の計画
func ChallengeC_108() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // 観光名所数
	kl := make(map[int]struct {
		idx int
		tt  int
		il  []string
	}, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		xi, _ := strconv.Atoi(sc.Text()) // i番目の観光名所滞在時間
		kl[i] = struct {
			idx int
			tt  int
			il  []string
		}{
			idx: i,
			tt:  xi,
		}
	}
	for i := 0; i < N; i++ {
		sc.Scan()
		iArr := strings.Split(sc.Text(), " ")
		kl[i] = struct {
			idx int
			tt  int
			il  []string
		}{
			idx: kl[i].idx,
			tt:  kl[i].tt,
			il:  iArr,
		}
	}
	sc.Scan()
	K, _ := strconv.Atoi(sc.Text())

	res := 0
	pre := 0
	for i := 0; i < K; i++ {
		sc.Scan()
		yi, _ := strconv.Atoi(sc.Text()) // 訪れたい観光名所

		if i == 0 {
			res = kl[yi-1].tt
			pre = yi
			continue
		}

		res = res + kl[yi-1].tt // 滞在時間
		iStr := kl[pre-1].il[yi-1]
		iInt, _ := strconv.Atoi(iStr)
		res = res + iInt // 移動時間
		pre = yi
	}
	fmt.Println(res)
}
