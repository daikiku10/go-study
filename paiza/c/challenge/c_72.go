package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// モンスターの進化
func ChallengeC_72() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	ATK, _ := strconv.Atoi(iArr[0])
	DEF, _ := strconv.Atoi(iArr[1])
	AGI, _ := strconv.Atoi(iArr[2])

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text()) // 進化先の数

	isSin := false
	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		si := i2Arr[0]
		MinATK, _ := strconv.Atoi(i2Arr[1])
		MaxATK, _ := strconv.Atoi(i2Arr[2])
		MinDEF, _ := strconv.Atoi(i2Arr[3])
		MaxDEF, _ := strconv.Atoi(i2Arr[4])
		MinAGI, _ := strconv.Atoi(i2Arr[5])
		MaxAGI, _ := strconv.Atoi(i2Arr[6])

		// 攻撃
		if !(MinATK <= ATK && ATK <= MaxATK) {
			continue
		}

		// 防御
		if !(MinDEF <= DEF && DEF <= MaxDEF) {
			continue
		}

		// 素早さ
		if !(MinAGI <= AGI && AGI <= MaxAGI) {
			continue
		}

		isSin = true
		fmt.Println(si)
	}
	if !isSin {
		fmt.Println("no evolution")
	}
}
