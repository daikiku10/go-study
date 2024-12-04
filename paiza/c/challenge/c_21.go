package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_21() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	iArr := strings.Split(sc.Text(), " ")
	xc, _ := strconv.Atoi(iArr[0])
	yc, _ := strconv.Atoi(iArr[1])
	r1, _ := strconv.Atoi(iArr[2])
	r2, _ := strconv.Atoi(iArr[3])

	// r1 2 <= (x - xc)2 + (y - yc)2 <= r2 2

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	for i := 0; i < N; i++ {
		sc.Scan()
		i2Arr := strings.Split(sc.Text(), " ")
		x, _ := strconv.Atoi(i2Arr[0])
		y, _ := strconv.Atoi(i2Arr[1])

		if r1*r1 <= ((x-xc)*(x-xc))+((y-yc)*(y-yc)) && ((x-xc)*(x-xc))+((y-yc)*(y-yc)) <= r2*r2 {
			fmt.Println("yes")
			continue
		}
		fmt.Println("no")
	}

}
