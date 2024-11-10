package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_155() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	S := sc.Text()
	sc.Scan()
	input2 := sc.Text()
	N, _ := strconv.Atoi(input2)

	for i := 0; i < N; i++ {
		sc.Scan()
		title := sc.Text()

		if strings.Contains(title, S) {
			fmt.Println("Yes")
			continue
		}
		fmt.Println("No")
	}
}
