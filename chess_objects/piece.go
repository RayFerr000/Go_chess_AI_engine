package chess_objects

import (
		"math"
		"fmt"
       )
type Piece struct{
    
    Type int
    Location *Square
    Captured bool

}

//Move a piece from its current square to a new one if it is a valid move.
func (piece *Piece) Move_piece(newSquare *Square, board *Board){
	if newSquare.HasPiece || newSquare.OutOfBounds || !validate_move(piece.Location, newSquare, board, piece.Type){
		return
	}
	currentSquare := piece.Location
	newSquare.Piece = piece
	newSquare.HasPiece = true
	piece.Location = newSquare
	currentSquare.HasPiece = false
}

//Determine if moving a piece from currentSquare to newSquare is a valid move.
func validate_move(currentSquare *Square, newSquare *Square, board *Board, pieceType int) bool{
	oldX := currentSquare.X
	oldY := currentSquare.Y
	newX := newSquare.X
	newY := newSquare.Y
    switch{
        case math.Abs(float64(pieceType)) == 1:  //pawn
        	if (newX - oldX == 1) && (oldY == newY){
        		return true
        	}else{
        		fmt.Println("Not a valid pawn move")
        		return false
        	}
        /**
        case math.Abs(float64(pieceType)) == 2:
        	continue 
        case math.Abs(float64(pieceType)) == 2:
        case math.Abs(float64(pieceType)) == 2:
        case math.Abs(float64(pieceType)) == 2:
        case math.Abs(float64(pieceType)) == 2:
*/
    }
    return false
}