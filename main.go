package main

import (
		"fmt"
		objects "Go_chess_AI_engine/chess_objects"	
		)

func main(){

	board := objects.Board{}
	board.New()
    printBoard(&board)	
}

func printBoard(b * objects.Board){

    for i := 2; i <= 9; i++{
        for j := 2; j <= 9; j++{
            if b.Board[i][j].HasPiece == true{
                if b.Board[i][j].Piece.Type > 0{
                    fmt.Print(b.Board[i][j].Piece.Type)
                    fmt.Print("  ")
                }else{
                    fmt.Print(b.Board[i][j].Piece.Type)
                    fmt.Print(" ")
                }
            }else{
                fmt.Print("X  ")
            }

        }
        fmt.Println(" ")
    }
}