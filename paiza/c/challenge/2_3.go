package challenge

import "fmt"

func Challenge2_3() {
	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Println(i + 1)
		} else {
			fmt.Print(i + 1)
			fmt.Print(" ")
		}
	}
}
