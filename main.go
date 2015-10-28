package main

import (
		"fmt"
		objects "Go_chess_AI_engine/chess_objects"	
		)

func main(){

	b := objects.Board{}
	b.New()
    piece1 := b.Board[3][6].Piece  //  White pawn
    piece2 := b.Board[8][6].Piece  //  black pawn
    piece3 := b.Board[2][6].Piece  //  White king
    piece4 := b.Board[9][6].Piece  //  black king

    /*
    piece1_2 := b.Board[3][5].Piece  // pawn
    piece1_3 := b.Board[3][7].Piece  // pawn
    piece2 := b.Board[2][3].Piece  // knight
    piece3 := b.Board[2][2].Piece  // castle
    piece4 := b.Board[2][4].Piece  //bishop
    */
    piece1.Move_piece(b.Board[5][6], &b)
    piece2.Move_piece(b.Board[6][6], &b)
    piece3.Move_piece(b.Board[3][6], &b)
    piece4.Move_piece(b.Board[8][6], &b)
    piece3.Move_piece(b.Board[4][6], &b)
    piece4.Move_piece(b.Board[7][6], &b)
    piece3.Move_piece(b.Board[3][6], &b)
    piece4.Move_piece(b.Board[8][6], &b)
    /*
    piece1.Move_piece(b.Board[6][2], &b)
    piece1_2.Move_piece(b.Board[4][5], &b)
    piece1_3.Move_piece(b.Board[4][7], &b)
    piece1_3.Move_piece(b.Board[5][7], &b)
    piece2.Move_piece(b.Board[4][4], &b)
    piece3.Move_piece(b.Board[4][2], &b)
    piece4.Move_piece(b.Board[6][8], &b)
    */
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
                fmt.Print("O  ")
            }

        }
        fmt.Println(" ")
    }
}