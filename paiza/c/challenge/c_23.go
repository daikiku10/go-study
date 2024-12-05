package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_23() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		count := 0
		for _, v := range iArr {
			for _, j := range i2Arr {
				if v == j {
					count++
				}
			}
		}
		fmt.Println(count)
	}
}
