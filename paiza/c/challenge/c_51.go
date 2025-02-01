package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// カード並べ 97点
func ChallengeC_51() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	A, _ := strconv.Atoi(iArr[0])
	B, _ := strconv.Atoi(iArr[1])
	C, _ := strconv.Atoi(iArr[2])
	D, _ := strconv.Atoi(iArr[3])

	resMAX := 0
	resMAX = ka(A, B, C, D)
	ABDC := ka(A, B, D, C)
	if resMAX < ABDC {
		resMAX = ABDC
	}
	ACBD := ka(A, C, B, D)
	if resMAX < ACBD {
		resMAX = ACBD
	}
	ACDB := ka(A, C, D, B)
	if resMAX < ACDB {
		resMAX = ACDB
	}
	ADBC := ka(A, D, B, C)
	if resMAX < ADBC {
		resMAX = ADBC
	}
	ADCB := ka(A, D, C, B)
	if resMAX < ADCB {
		resMAX = ADCB
	}
	CABD := ka(C, A, B, D)
	if resMAX < CABD {
		resMAX = CABD
	}
	CDBA := ka(C, D, B, A)
	if resMAX < CDBA {
		resMAX = CDBA
	}
	DBCA := ka(D, B, C, A)
	if resMAX < DBCA {
		resMAX = DBCA
	}
	DABC := ka(D, A, B, C)
	if resMAX < DABC {
		resMAX = DABC
	}
	DACB := ka(D, A, C, B)
	if resMAX < DACB {
		resMAX = DACB
	}
	DCBA := ka(D, C, B, A)
	if resMAX < DCBA {
		resMAX = DCBA
	}
	fmt.Println(resMAX)
}

func ka(a, b, c, d int) int {
	return ((a * 10) + b) + ((c * 10) + d)
}

// func maxKa(A, B, C, D int) int {
// 	permutations := [][]int{
// 		{A, B, C, D}, {A, B, D, C}, {A, C, B, D}, {A, C, D, B},
// 		{A, D, B, C}, {A, D, C, B}, {C, A, B, D}, {C, D, B, A},
// 		{D, B, C, A}, {D, A, B, C}, {D, A, C, B}, {D, C, B, A},
// 	}

// 	resMAX := 0
// 	for _, p := range permutations {
// 		res := ka(p[0], p[1], p[2], p[3])
// 		if res > resMAX {
// 			resMAX = res
// 		}
// 	}

// 	return resMAX
// }

// func main() {
// 	A, B, C, D := 1, 2, 3, 4 // 例として値を設定
// 	fmt.Println(maxKa(A, B, C, D))
// }
