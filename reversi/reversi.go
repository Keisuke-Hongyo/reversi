package Reversi

import (
	"fmt"
	"math"
)

const PlayerA = 1
const PlayerB = 2

const xMax = 8
const yMax = 8

// Reversi 構造体・レシーバ
// ------------------------------------------------------------
type Reversi struct {
	matrix         [yMax][xMax]int8 // 盤面データ
	nowPlayer      int8             // 現在のプレイヤー
	nextPlayer     int8             // 現在のプレイヤー
	player_A_point uint8            // プレイヤーAの得点
	player_B_point uint8            // プレイヤーBの得点
}

// 外部用レシーバー
// --------------------------------------------------------------

// New 初期化
func New() Reversi {
	var b Reversi
	b.initialize()
	return b
}

// Play 操作用ルーチン
func (b *Reversi) Play() {
	var x, y, swp int8
	for {
		// 盤の表示
		b.printMatrix()

		// コンソールの表示
		if b.nowPlayer == PlayerA {
			fmt.Print("Player_A x,y = ")
		} else {
			fmt.Print("Player_B x,y = ")
		}

		// 座標の入力
		_, _ = fmt.Scanf("%d,%d", &x, &y)

		// 石が置けるか確認
		if !b.setStone(x, y, b.nowPlayer) {
			// 置けない時の表示
			fmt.Printf(" x=%d y=%dは置けません\n", x, y)
		} else {
			// 置けるならループを抜ける
			break
		}
	}

	//次の相手に変更
	swp = b.nowPlayer
	b.nowPlayer = b.nextPlayer
	b.nextPlayer = swp
}

// 内部使用レシーバー
// --------------------------------------------------------------

// initialize 初期化ルーチン
func (b *Reversi) initialize() {
	// 盤面の初期化
	b.matrix = [xMax][xMax]int8{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 2, 0, 0, 0},
		{0, 0, 0, 2, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	// 先手の設定
	b.nowPlayer = PlayerA
	b.nextPlayer = PlayerB
}

// 置き石の処理
func (b *Reversi) setStone(x int8, y int8, st int8) bool {
	// 置き場所に石があるか？
	if b.matrix[y][x] != 0 {
		// もう置かれているから置けない
		return false
	}

	// 処理ルーチン
	chk := false
	for ang := 0.0; ang < 360; ang += 45.0 {
		// 検索方向の設定
		dx, dy := distnation(ang)
		// 石がおけるか確認
		c, cnt := b.stCheck(x, y, dx, dy)
		// 石が置けるなら
		if c {
			// 石のひっくり返し処理
			b.reverseStone(x, y, dx, dy, cnt)
			chk = true
		}
	}
	return chk
}

// ポイント計算処理
func (b *Reversi) pointUpdate() {
	// ポイント初期化
	b.player_A_point = 0
	b.player_B_point = 0

	// ポイントカウント処理
	for y := 0; y < len(b.matrix); y++ {
		for x := 0; x < len(b.matrix[y]); x++ {
			if PlayerA == b.matrix[y][x] {
				b.player_A_point++
			} else if PlayerB == b.matrix[y][x] {
				b.player_B_point++
			}
		}
	}
}

// ボードの表示
func (b *Reversi) printMatrix() {
	fmt.Println("   01234567   ")
	for y := 0; y < len(b.matrix); y++ {
		fmt.Print(" ")
		fmt.Print(y)
		fmt.Print(" ")
		for x := 0; x < len(b.matrix[y]); x++ {
			fmt.Print(b.matrix[y][x])
		}
		fmt.Print(" ")
		fmt.Print(y)
		fmt.Println(" ")
	}
	fmt.Println("   01234567   ")

	// ポイント表示
	b.pointUpdate()
	fmt.Printf("Player_A = %d Player_B = %d\n", b.player_A_point, b.player_B_point)
}

// 石の反転処理
func (b *Reversi) reverseStone(x int8, y int8, dx int8, dy int8, cnt int8) {
	var i int8
	//ひっくり返す処理
	b.matrix[y][x] = b.nowPlayer
	for i = 0; i < cnt; i++ {
		y = y + dy
		x = x + dx
		b.matrix[y][x] = b.nowPlayer
	}
}

// 置ける場所か確認処理
func (b *Reversi) stCheck(x int8, y int8, dx int8, dy int8) (bool, int8) {
	var cnt int8
	cnt = 0
	for {
		// 盤上の端をチェック
		if (x+dx) < 0 || (x+dx) > (xMax-1) || (y+dy) < 0 || (y+dy) > (yMax-1) {
			break
		} else {
			// 確認マスの移動
			x = x + dx
			y = y + dy
		}

		//
		if b.matrix[y][x] != b.nextPlayer {
			break
		}
		cnt = cnt + 1
	}

	if b.matrix[y][x] == b.nowPlayer && cnt != 0 {
		return true, cnt
	} else {
		return false, 0
	}
}

// --------------------------------------------------------------

// 内部関数
// -------------------------------------------------------------------------
// 検索方向処理
func distnation(ang float64) (int8, int8) {
	dx := int8(math.Round(math.Cos(ang / 180.0 * math.Pi)))
	dy := (-1.0) * int8(math.Round(math.Sin(ang/180.0*math.Pi)))
	return dx, dy
}

// -------------------------------------------------------------------------
