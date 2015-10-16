package main

import (
		"fmt"
		objects "Go_chess_AI_engine/chess_objects"	
		)

func main(){

	b := objects.Board{}
	b.New()
    piece := b.Board[3][2].Piece

    piece.Move_piece(b.Board[4][2], &b)
    printBoard(&b)	
}

func printBoard(b * objects.Board){

    for i := 9; i >= 2; i--{
        for j := 9; j >= 2; j--{
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