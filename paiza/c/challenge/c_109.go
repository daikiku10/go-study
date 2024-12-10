package challenge

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// IDを登録順に並べよう
func ChallengeC_109() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	type User struct {
		idx  int
		Name string
	}

	ul := make([]User, 0, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		si := sc.Text()
		siArr := strings.Split(si, "")
		b := 0
		for j := 0; j < len(siArr); j++ {
			_, err := strconv.Atoi(siArr[j])
			if err == nil {
				b = j
				break
			}
		}
		siInd, _ := strconv.Atoi(si[b:])
		ul = append(ul, User{
			idx:  siInd,
			Name: si,
		})
	}

	sort.SliceStable(ul, func(i, j int) bool { return ul[i].idx < ul[j].idx })

	for _, v := range ul {
		fmt.Println(v.Name)
	}
}
