package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Challenge4_1() { // 4_2, 4_3, 4_4
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	input, _ := strconv.Atoi(sc.Text()) // 数値であること

	for i := 0; i < input; i++ {
		fmt.Println(i + 1)
	}

	// paiza側のgoバージョンがおそらく 1.22とか？
	// for v := range input {
	// 	fmt.Println(v + 1)
	// }
}
