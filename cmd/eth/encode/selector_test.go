package encode

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestGetSelector(*testing.T) {
	// s:=GetSelector("TraderJoeWithDraw(address,address,uint256,address)")
	// println(s)

	// f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	// if err != nil {
	//    return
	// }

	// defer func() {
	//    f.Close()
	// }()

	// // 组合一下即可，os.Stdout代表标准输出流
	// multiWriter := io.MultiWriter(os.Stdout, f)
	// log.SetOutput(multiWriter)
	// log.SetFlags( log.Ltime | log.Lshortfile)
	// log.Println("aaa")

	// s := "ethusdt"
	// log.Println(s[:2])

	// a := struct {
	// 	A decimal.Decimal
	// }{}
	// println(a.A.Equal(decimal.NewFromInt(0)))

	// net.Dial("1.15.132.255:11453",)
}

type Aaa struct {
	Aaaa string
	Bbbb int
}

// func TestBigMul(*testing.T) {
// 	var a Aaa
// 	fmt.Println(a)
// 	if a.Aaaa == "" {
// 		fmt.Println(111)
// 	} else {
// 		fmt.Println(222)
// 	}
// }

// func echo[T any](t T) string {
// 	return fmt.Sprintf("%v", t)
// }
// func TestAA(*testing.T) {
// 	echo(0)
// 	echo(int32(0))
// 	echo(uint32(0))
// 	echo(uint64(0))
// 	echo("hello")
// 	echo(struct{}{})
// 	echo(time.Now())
// }

// This program solves the (English) peg
// solitaire board game.
// http://en.wikipedia.org/wiki/Peg_solitaire

// package main

// import "fmt"

const N = 11 + 1 // length of a row (+1 for \n)

// The board must be surrounded by 2 illegal
// fields in each direction so that move()
// doesn't need to check the board boundaries.
// Periods represent illegal fields,
// ● are pegs, and ○ are holes.

var board = []rune(
	`...........
...........
....●●●....
....●●●....
..●●●●●●●..
..●●●○●●●..
..●●●●●●●..
....●●●....
....●●●....
...........
...........
`)

// center is the position of the center hole if
// there is a single one; otherwise it is -1.
var center int

func init() {
	n := 0
	for pos, field := range board {
		if field == '○' {
			center = pos
			n++
		}
	}
	if n != 1 {
		center = -1 // no single hole
	}
}

var moves int // number of times move is called

// move tests if there is a peg at position pos that
// can jump over another peg in direction dir. If the
// move is valid, it is executed and move returns true.
// Otherwise, move returns false.
func move(pos, dir int) bool {
	moves++
	if board[pos] == '●' && board[pos+dir] == '●' && board[pos+2*dir] == '○' {
		board[pos] = '○'
		board[pos+dir] = '○'
		board[pos+2*dir] = '●'
		return true
	}
	return false
}

// unmove reverts a previously executed valid move.
func unmove(pos, dir int) {
	board[pos] = '●'
	board[pos+dir] = '●'
	board[pos+2*dir] = '○'
}

// solve tries to find a sequence of moves such that
// there is only one peg left at the end; if center is
// >= 0, that last peg must be in the center position.
// If a solution is found, solve prints the board after
// each move in a backward fashion (i.e., the last
// board position is printed first, all the way back to
// the starting board position).
func solve() bool {
	var last, n int
	for pos, field := range board {
		// try each board position
		if field == '●' {
			// found a peg
			for _, dir := range []int{-1, -N, +1, +N} {
				// try each direction
				if move(pos, dir) {
					// a valid move was found and executed,
					// see if this new board has a solution
					if solve() {
						unmove(pos, dir)
						fmt.Println(string(board))
						return true
					}
					unmove(pos, dir)
				}
			}
			last = pos
			n++
		}
	}
	// tried each possible move
	if n == 1 && (center < 0 || last == center) {
		// there's only one peg left
		fmt.Println(string(board))
		return true
	}
	// no solution found for this board
	return false
}

func TestPeg(*testing.T) {

	if !solve() {
		fmt.Println("no solution found")
	}
	fmt.Println(moves, "moves tried")
}

func TestTopic(*testing.T) {
	eventSignature := []byte("AaveDeposit(address,uint256,address)")
	eventSignature2 := []byte("AaveBorrow(address,uint256,address,uint256)")
	_ = eventSignature2
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
}
