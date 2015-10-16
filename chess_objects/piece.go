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
		fmt.Println("Not a valid move.")
		return
	}
	currentSquare := piece.Location
	newSquare.Piece = piece
	newSquare.HasPiece = true
	piece.Location = newSquare
	currentSquare.HasPiece = false
}

//validate_move determines if moving any piece from currentSquare to newSquare is a valid move.
func validate_move(currentSquare *Square, newSquare *Square, board *Board, pieceType int) bool{
	oldX := currentSquare.X
	oldY := currentSquare.Y
	newX := newSquare.X
	newY := newSquare.Y
    switch{
        
        case math.Abs(float64(pieceType)) == 1:  //pawn
        	return valid_pawn_move(oldX, oldY, newX, newY)
        
        case math.Abs(float64(pieceType)) == 2:
        	return valid_knight_move(oldX, oldY, newX, newY)
        /**
        case math.Abs(float64(pieceType)) == 2:
        case math.Abs(float64(pieceType)) == 2:
        case math.Abs(float64(pieceType)) == 2:
        case math.Abs(float64(pieceType)) == 2:
*/
    }
    return false
}

//valid_pawn_move determines if this is a valid pawn move.
func valid_pawn_move(oldX int, oldY int, newX int, newY int) bool{
    
    if (newX - oldX == 1) && (oldY == newY){
		return true
	}else{
		fmt.Println("Not a valid pawn move")
		return false
	}
}

//valid_knight_move determines if this is a knight pawn move.
func valid_knight_move(oldX int, oldY int, newX int, newY int) bool{
	cond1 := (math.Abs(float64(newX)) - math.Abs(float64(oldX))) == 2.0 && (math.Abs(float64(newY)) - math.Abs(float64(oldY))) == 1.0 //Two horiontal one vertical
	cond2 := (math.Abs(float64(newX)) - math.Abs(float64(oldX))) == 1.0 && (math.Abs(float64(newY)) - math.Abs(float64(oldY))) == 2.0 //Two vertical one horizontal
    if (cond1) || (cond2){
		return true
	} else{
		fmt.Println("Not a valid knight move")
		return false
	}

}