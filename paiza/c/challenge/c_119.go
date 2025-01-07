package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// お菓子かいたずらか 不正解(順番が重要だがmapでfor rangeを使用したため)
func ChallengeC_119() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	N, _ := strconv.Atoi(iArr[0]) // 家の数
	M, _ := strconv.Atoi(iArr[1]) // 男の子の人数
	K, _ := strconv.Atoi(iArr[2]) // Aliceが帰ってしまう件数

	li := make(map[int]struct {
		idx   int
		isMan bool
	}, N)

	for i := 1; i <= N; i++ {
		li[i] = struct {
			idx   int
			isMan bool
		}{
			idx:   i,
			isMan: false,
		}
	}

	for i := 0; i < M; i++ {
		sc.Scan()
		bi, _ := strconv.Atoi(sc.Text())
		li[bi] = struct {
			idx   int
			isMan bool
		}{
			idx:   li[bi].idx,
			isMan: true,
		}
	}

	oCnt := 0
	bCnt := 0
	for i := 1; i <= len(li); i++ {
		if li[i].isMan {
			bCnt++
			if bCnt >= K {
				fmt.Println(oCnt)
				return
			}
			continue
		}
		bCnt = 0
		oCnt++
	}
	fmt.Println(oCnt)
}
