package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 【キャンペーン問題】 下桁ルール
func ChallengeC_93() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")

	xCnt := 0
	xl := strings.Split(iArr[0], "")
	yCnt := 0
	yl := strings.Split(iArr[1], "")

	for _, v := range xl {
		vInt, _ := strconv.Atoi(v)
		xCnt = xCnt + vInt
	}

	for _, v := range yl {
		vInt, _ := strconv.Atoi(v)
		yCnt = yCnt + vInt
	}

	xCl := strings.Split(strconv.Itoa(xCnt), "")
	yCl := strings.Split(strconv.Itoa(yCnt), "")

	xRes, _ := strconv.Atoi(xCl[len(xCl)-1])
	yRes, _ := strconv.Atoi(yCl[len(yCl)-1])

	if xRes > yRes {
		fmt.Println("Bob")
		return
	}

	if yRes > xRes {
		fmt.Println("Alice")
		return
	}

	if xRes == yRes {
		fmt.Println("Draw")
	}
}
