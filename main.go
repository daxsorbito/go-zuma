package main

import (
	"fmt"
	"strings"
)

func main() {
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
		foundIndex := getFirstIndexOfFirstBallFromBoard(board, currentBall)
		if foundIndex > -1 {
			lastIndex := getLastSubsequestBallIndex(board, foundIndex)
			if (lastIndex - foundIndex) >= 2 { // 3 pairs
				// fmt.Println(">>>", board, currentBall, foundIndex, lastIndex, len(board))
				board = append(board[:foundIndex], board[lastIndex:len(board)]...)
				continue
			}

			board = append(append(board[:foundIndex], currentBall), board[foundIndex:]...)
			continue
		}
		board = append(board, currentBall)

	}
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
