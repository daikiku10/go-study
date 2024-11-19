package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_156() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	N, _ := strconv.Atoi(i) // 日数

	hour := 0
	second := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		e := sc.Text()
		eArr := strings.Split(e, " ")
		in := eArr[0]
		out := eArr[1]
		inArr := strings.Split(in, ":")
		outArr := strings.Split(out, ":")
		inH := 0
		inS := 0
		if inArr[0][0:1] == "0" {
			inH, _ = strconv.Atoi(inArr[0][1:]) // 後の文字を取得
		} else {
			inH, _ = strconv.Atoi(inArr[0])
		}
		if inArr[1][0:1] == "0" {
			inS, _ = strconv.Atoi(inArr[1][1:]) // 後の文字を取得
		} else {
			inS, _ = strconv.Atoi(inArr[1])
		}
		outH := 0
		outS := 0
		if outArr[0][0:1] == "0" {
			outH, _ = strconv.Atoi(outArr[0][1:]) // 後の文字を取得
		} else {
			outH, _ = strconv.Atoi(outArr[0])
		}
		if outArr[1][0:1] == "0" {
			outS, _ = strconv.Atoi(outArr[1][1:]) // 後の文字を取得
		} else {
			outS, _ = strconv.Atoi(outArr[1])
		}

		h := outH - inH
		s := outS - inS
		if 0 > s {
			h = h - 1
			s = 60 + s
		}
		hour = hour + h
		second = second + s
	}

	plus := 0
	if second > 60 {
		plus = second / 60
		second = second - (60 * plus)
	}
	hour = hour + plus
	fmt.Printf("%v %v", hour, second)
}
