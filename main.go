package main

import (
		"fmt"
		objects "Go_chess_AI_engine/chess_objects"
		)

func main(){
	board := objects.Board{}
	board.New()
	fmt.Println(board)
}