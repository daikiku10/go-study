package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Event struct {
	ID int
	time.Time
}
type Event2 struct {
	ID   int
	Time time.Time
}

func main() {
	event := Event{
		ID:   1234,
		Time: time.Now(),
	}
	event2 := Event2{
		ID:   5678,
		Time: time.Now(),
	}
	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}
	b2, err := json.Marshal(event2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	fmt.Println(string(b2))

	// s := "hello"
	// fmt.Println(len(s))
	// k := "漢"
	// fmt.Println(len(k))
	// b := string([]byte{0xE6, 0xBC, 0xA2})
	// fmt.Printf("%s\n", b)

	// s := "héllo"
	// // パターン1
	// for i, v := range s {
	// 	fmt.Printf("position %d: %c\n", i, v)
	// }
	// fmt.Printf("len=%d\n", len(s))
	// fmt.Println(utf8.RuneCountInString(s))

	// // パターン2
	// runes := []rune(s)
	// for i, v := range runes {
	// 	fmt.Printf("position %d: %c\n", i, v)
	// }

	// 暗号化
	// encrypted, _ := utils.Aes256Encrypt("daiki", "erfd", "3280m18xuzTbGwC7")
	// encode := base64.StdEncoding.EncodeToString(encrypted)
	// more := base64.StdEncoding.EncodeToString([]byte(encode))
	// more = strings.ReplaceAll(more, "=", "_")
	// more = strings.ReplaceAll(more, "/", "-")
	// fmt.Println(concat([]string{"d", "a", "i"}))

}

// NO.39
func concat(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

// NO.40
func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader) // bは[]byte
	if err != nil {
		return nil, err
	}
	// sanitizeを呼び出す
	return sanitize(b), nil
}

func sanitize(b []byte) []byte {
	return bytes.TrimSpace(b)
}

func getFileData(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner)
}

func sum(s []int64) int64 {
	var total int64
	length := len(s)
	for i := 0; i < length; i++ {
		total += s[i]
	}
	return total
}
