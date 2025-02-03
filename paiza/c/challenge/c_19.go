package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 完全数とほぼ完全数
func ChallengeC_19() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	Q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < Q; i++ {
		sc.Scan()
		n, _ := strconv.Atoi(sc.Text())

		cnt := 0
		for j := 1; j <= n; j++ {
			if n == j {
				continue
			}
			if n%j == 0 {
				cnt = cnt + j
			}
		}

		if cnt == n {
			fmt.Println("perfect")
			continue
		}
		if cnt+1 == n || cnt-1 == n {
			fmt.Println("nearly")
			continue
		}
		fmt.Println("neither")
	}
}
