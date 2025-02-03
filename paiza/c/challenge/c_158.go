package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 定期購買の価格計算
func ChallengeC_158() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0]) // 種類
	M, _ := strconv.Atoi(iArr[1]) // 何週間後までか

	nList := make(map[int]struct {
		f int
		p int
	}, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		fi, _ := strconv.Atoi(i2Arr[0]) // 購入頻度
		pi, _ := strconv.Atoi(i2Arr[1]) // 価格
		nList[i] = struct {
			f int
			p int
		}{
			f: fi,
			p: pi,
		}
	}

	mList := make([][]int, 0, M)
	for i := 0; i < M; i++ {
		nl := make([]int, 0, N)
		for _, v := range nList {
			// 最初の週は全て購入する
			if i == 0 {
				nl = append(nl, v.p)
				continue
			}

			if i%v.f == 0 {
				nl = append(nl, v.p)
			}
		}
		mList = append(mList, nl)
	}

	res := 0
	for _, v := range mList {
		w := 0
		if len(v) >= 2 {
			w = 10
		}
		if len(v) >= 3 {
			w = 15
		}

		kei := 0
		for _, j := range v {
			kei = kei + j
		}

		if w == 0 {
			res = res + kei
			continue
		}

		if w == 10 {
			keiF := float64(kei)
			a := keiF * 0.9
			res = res + int(a)
		}

		if w == 15 {
			keiF := float64(kei)
			a := keiF * 0.85
			res = res + int(a)
		}
	}
	fmt.Println(res)
}
