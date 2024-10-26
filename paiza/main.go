package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("test")
}

// 値取得・出力サンプル
func inputOutput() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	var n, _ = strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		var s = strings.Split(sc.Text(), " ")
		fmt.Println("hello = " + s[0] + " , world = " + s[1])
	}
}
