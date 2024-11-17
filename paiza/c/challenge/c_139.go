package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ChallengeC_139() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	N, _ := strconv.Atoi(i)

	nItems := make([]int, N)
	items := make([]int, N)
	misItems := []int{}
	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()
		x, _ := strconv.Atoi(i2)
		nItems[i] = i + 1
		items[i] = x
	}

	for _, product := range nItems {
		found := false
		for _, order := range items {
			if product == order {
				found = true
				break
			}
		}

		if !found {
			misItems = append(misItems, product)
		}
	}

	fmt.Println(len(misItems))
}
