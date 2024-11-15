package challenge

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ChallengeB_145() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i := sc.Text()
	iArr := strings.Split(i, " ")

	N, _ := strconv.Atoi(iArr[0]) // 縦横の大きさ
	K, _ := strconv.Atoi(iArr[1]) // 抽選回数

	cards := make([][]int, N) // ビンゴ表

	for i := 0; i < N; i++ {
		sc.Scan()
		i2 := sc.Text()

		i2Arr := strings.Split(i2, " ")
		i2ArrInt := make([]int, len(i2Arr))
		for j, v := range i2Arr {
			vInt, _ := strconv.Atoi(v)
			i2ArrInt[j] = vInt
		}
		cards[i] = i2ArrInt
	}

	sc.Scan()
	i3 := sc.Text()
	i3Arr := strings.Split(i3, " ")
	items := make([]int, K) // 抽選番号
	for i, v := range i3Arr {
		vInt, _ := strconv.Atoi(v)
		items[i] = vInt
	}

	// 結果のカード
	yokoResult := make([]int, N)
	tateResult := make([]int, N)
	naname1Result := 0
	naname2Result := 0

	// ビンゴ表を繰り返し
	for j := 0; j < N; j++ {
		yoko := cards[j]
		count := 0
		for k := 0; k < N; k++ {
			target := yoko[k]

			// 当たり
			if target == 0 {
				count++
				tateResult[k]++
				if j == k {
					naname1Result++
				}
				if j+k == N-1 {
					naname2Result++
				}
			}

			// 抽選番号で繰り返し
			for i := 0; i < len(items); i++ {
				bango := items[i]

				// 当たり
				if bango == target {
					count++
					tateResult[k]++
					if j == k {
						naname1Result++
					}
					if j+k == N-1 {
						naname2Result++
					}
				}
				// ハズレ
			}
		}
		yokoResult[j] = count
	}

	res := 0

	for _, v := range yokoResult {
		if v == N {
			res++
		}
	}
	for _, v := range tateResult {
		if v == N {
			res++
		}
	}
	if naname1Result == N {
		res++
	}
	if naname2Result == N {
		res++
	}
	fmt.Println(res)
}

// 終了後に調べた結果
// func main() {
// 	// ビンゴのカードサイズ
// 	const N = 5 // 例：5x5のビンゴカード
// 	bingoCard := make([][]int, N)
// 	marked := make([][]bool, N)

// 	// ビンゴカードとマーク配列の初期化
// 	for i := 0; i < N; i++ {
// 		bingoCard[i] = make([]int, N)
// 		marked[i] = make([]bool, N)
// 	}

// 	// ビンゴカードのデータと呼ばれた数字（例として、マスに入れる値を仮定）
// 	// 実際にはここで入力処理が必要です。
// 	numbers := []int{1, 2, 3, 4, 5} // 選ばれた数字のリスト（仮定）

// 	// 呼ばれた数字をビンゴカード上にマーク
// 	for _, num := range numbers {
// 		for i := 0; i < N; i++ {
// 			for j := 0; j < N; j++ {
// 				if bingoCard[i][j] == num {
// 					marked[i][j] = true
// 				}
// 			}
// 		}
// 	}

// 	// ビンゴのカウント
// 	bingoCount := 0

// 	// 横のビンゴ判定
// 	for i := 0; i < N; i++ {
// 		rowBingo := true
// 		for j := 0; j < N; j++ {
// 			if !marked[i][j] {
// 				rowBingo = false
// 				break
// 			}
// 		}
// 		if rowBingo {
// 			bingoCount++
// 		}
// 	}

// 	// 縦のビンゴ判定
// 	for j := 0; j < N; j++ {
// 		colBingo := true
// 		for i := 0; i < N; i++ {
// 			if !marked[i][j] {
// 				colBingo = false
// 				break
// 			}
// 		}
// 		if colBingo {
// 			bingoCount++
// 		}
// 	}

// 	// 対角線のビンゴ判定（右下がり）
// 	diagBingo1 := true
// 	for i := 0; i < N; i++ {
// 		if !marked[i][i] {
// 			diagBingo1 = false
// 			break
// 		}
// 	}
// 	if diagBingo1 {
// 		bingoCount++
// 	}

// 	// 対角線のビンゴ判定（左下がり）
// 	diagBingo2 := true
// 	for i := 0; i < N; i++ {
// 		if !marked[i][N-i-1] {
// 			diagBingo2 = false
// 			break
// 		}
// 	}
// 	if diagBingo2 {
// 		bingoCount++
// 	}

// 	fmt.Printf("ビンゴの数: %d\n", bingoCount)
// }
