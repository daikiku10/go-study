package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_142() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	S := sc.Text()

	sc.Scan()
	input2 := sc.Text()
	N, _ := strconv.Atoi(input2)

	sc.Scan()
	input3 := sc.Text()
	menus := strings.Split(input3, " ")

	ans := "No"
	for i := 0; i < N; i++ {
		if strings.Contains(menus[i], S) {
			ans = "Yes"
			break
		}
	}
	fmt.Println(ans)

}
