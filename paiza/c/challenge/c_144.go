package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeC_144() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	N, _ := strconv.Atoi(input)

	winCount := 0
	for i := 0; i < N; i++ {
		sc.Scan()
		input2 := sc.Text()
		input2Arr := strings.Split(input2, " ")
		aris := input2Arr[0]
		bob := input2Arr[1]
		if aris == "G" && bob == "C" {
			winCount++
		}
		if aris == "C" && bob == "P" {
			winCount++
		}
		if aris == "P" && bob == "G" {
			winCount++
		}
	}
	fmt.Println(winCount)
}
