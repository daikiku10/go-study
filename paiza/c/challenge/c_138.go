package challenge

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func ChallengeC_138() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // 人数

	type User struct {
		idx  int
		cnt  int
		rank int
	}
	// 生徒リストの作成
	sList := make([]User, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text()) // 記録
		sList[i] = User{
			idx: i,
			cnt: a,
		}
	}

	// 記録順
	sort.SliceStable(sList, func(i, j int) bool { return sList[i].cnt >= sList[j].cnt })

	// 順位づけ
	j := 1
	for i := range sList {
		if i == 0 {
			sList[i].rank = j
			j++
			continue
		}

		if sList[i].cnt == sList[i-1].cnt {
			sList[i].rank = sList[i-1].rank
			j++
			continue
		}
		sList[i].rank = j
		j++
	}

	// idx順
	sort.SliceStable(sList, func(i, j int) bool { return sList[i].idx < sList[j].idx })

	for _, v := range sList {
		fmt.Println(v.rank)
	}
}
