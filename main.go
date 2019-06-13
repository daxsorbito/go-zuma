package main

import (
	"fmt"
	"strings"
)

func main() {
	// Zuma
	/*
		1. 	Input: WWRRRBB, RBW
			Output: <empty>
			Process: WWRR[R]BB -> WWBB -> WWBB[B] -> WW -> [W]WW -> <empty>
		2. 	Input: QWERRDDSA, RDSSEAQA
			Output: QQWWEE
			Process: QWERR[R]DDSA -> QWEDDSA -> QWED[D]DSA -> QWESA -> QWES[S]A -> QWESSA -> QWESS[S]A -> QWEA -> QWE[E]A -> QWEEA -> QWEEA[A] -> QWEEAA -> Q[Q]WEEAA -> QQWEEAA -> QQWEEAA[A] -> QQWEE
		3. 	Input: ASDFFDSAFFFFRT, AADSDFSF
			Output: RT
			Process: [A]ASDFFDSAFFFFRT -> AASDFFDSAFFFFRT -> [A]AASDFFDSAFFFFRT -> SD[D]FFDSAFFFFRT -> SDDFFDSAFFFFRT -> [S]SDDFFDSAFFFFRT -> SSDDFFDSAFFFFRT -> SS[D]DDFFDSAFFFFRT
					-> SSFFDSAFFFFRT -> SS[F]FFDSAFFFFRT -> SSDSAFFFFRT -> [S]SSDSAFFFFRT -> DSAFFFFRT -> DSA[F]FFFFRT -> DSART

	*/
	fmt.Println("WWRRRBB, RBW >>", getRemainingCandies("WWRRRBB", "RBW"))
	fmt.Println("QWERRDDSA, RDSSEAQA >>", getRemainingCandies("QWERRDDSA", "RDSSEAQA"))
	fmt.Println("ASDFFDSAFFFFRT, AADSDFSF >>", getRemainingCandies("ASDFFDSAFFFFRT", "AADSDFSF"))
}

func getRemainingCandies(inputBoard, inputBalls string) string {
	/*
		1. Get the firstBall
		2. Get the first index of the firstBall in the board (NEXT: count the subsequesnt same firstBall in the board and insert the firstBall in that group instead)
		3. Insert the firstBall on the index found in step 2.
		4. If the location of the newly inserted firstBall's subsequent's same firstBall is greater than 3 then remove all of the subsequest's same firstBall with the newly inserted firstBall
		5. Recreate the board and loop to step 1 but get the next Ball
	*/
	board := strings.Split(inputBoard, "")
	balls := strings.Split(inputBalls, "")

	for _, currentBall := range balls {
		// fmt.Println("board>>>", board, currentBall)

		foundIndex := getFirstIndexOfFirstBallFromBoard(board, currentBall)
		// fmt.Println("foundIndex>>>", foundIndex)
		if foundIndex > -1 {
			lastIndex := getLastSubsequestBallIndex(board, foundIndex)
			// fmt.Println("foundIndex>>>", foundIndex, lastIndex, (lastIndex - foundIndex))
			if (lastIndex - foundIndex) >= 2 { // 3 pairs
				// fmt.Println(">>>", board, currentBall, foundIndex, lastIndex, len(board))
				board = append(board[:foundIndex], board[lastIndex:len(board)]...)
				continue
			}

			board = append(append(board[:foundIndex], currentBall), board[foundIndex:]...)

			// fmt.Println("after insert>>", board)
			continue
			// board = append(board[:foundIndex], append(board, append(board, board[foundIndex:]...)...)

		}
		board = append(board, currentBall)

	}
	// 1. Get the first Ball
	// currentBall := balls[0]
	// fmt.Println("board>>>", board)

	// 2. Get the first index of the firstBall in the board
	// foundIndex := getFirstIndexOfFirstBallFromBoard(board, currentBall)
	// lastIndex := getLastSubsequestBallIndex(board, foundIndex)
	// fmt.Println("foundIndex>>>>>>>>", foundIndex, board[foundIndex:], lastIndex, len(board), len(board), board[foundIndex:lastIndex], board)
	// board = append(board[:foundIndex], board[lastIndex:len(board)]...)
	// fmt.Println("newBoard>>>", board)

	return strings.Join(board, "")
}

func getFirstIndexOfFirstBallFromBoard(board []string, ball string) int {
	for k, v := range board {
		if v == ball {
			return k
		}
	}
	return -1
}

func getLastSubsequestBallIndex(board []string, startIndex int) int {
	currentBall := board[startIndex]
	for k, v := range board[startIndex:] {
		if v != currentBall {
			return k + startIndex
		}
	}
	return len(board)
}
