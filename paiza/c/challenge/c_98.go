package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_98() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	N, _ := strconv.Atoi(i) // 人数

	type User struct {
		idx   int
		count int
	}

	pList := make([]User, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()
		s, _ := strconv.Atoi(i2)
		pList[i] = User{
			idx:   i,
			count: s, // ボールの個数
		}
	}

	sc.Scan()
	i3 := sc.Text()
	M, _ := strconv.Atoi(i3)

	for i := 0; i < M; i++ {
		sc.Scan()
		i4 := sc.Text()
		i4Arr := strings.Split(i4, " ")
		a, _ := strconv.Atoi(i4Arr[0]) // 対象者が
		b, _ := strconv.Atoi(i4Arr[1]) // この人に
		c, _ := strconv.Atoi(i4Arr[2]) // ボールを渡す個数

		if pList[a-1].count-c < 0 {
			pList[b-1].count = pList[b-1].count + pList[a-1].count // もらう側
			pList[a-1].count = 0                                   // 対象者
		} else {
			pList[a-1].count = pList[a-1].count - c // 対象者
			pList[b-1].count = pList[b-1].count + c // もらう側
		}
	}

	for _, v := range pList {
		fmt.Println(v.count)
	}
}
