package main

import (
	"reversi/reversi"
)

// ---------------------------------------------------------------

func main() {
	board := Reversi.New()
	var mode uint8
	mode = 0
	for {
		// 対戦モード
		switch mode {
		// ２人対戦
		case 0:
			board.Play()
			break
		}
	}
}
