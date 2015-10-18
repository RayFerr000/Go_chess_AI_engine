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

    switch{
        
        case math.Abs(float64(pieceType)) == 1:  //pawn
        	return valid_pawn_move(currentSquare.X, currentSquare.Y, newSquare.X, newSquare.Y)
        
        case math.Abs(float64(pieceType)) == 2:  //knight
        	return valid_knight_move(currentSquare.X, currentSquare.Y, newSquare.X, newSquare.Y)

        case math.Abs(float64(pieceType)) == 4: //castle
        	if currentSquare.X == newSquare.X{  //horizontal castle move
        		return valid_castle_move(currentSquare.X, currentSquare.Y, newSquare.X, newSquare.Y, board.get_horizontal_path(currentSquare, newSquare))
        	
        	} else if currentSquare.Y == newSquare.Y {  //vertical castle move
        		return valid_castle_move(currentSquare.X, currentSquare.Y, newSquare.X, newSquare.Y, board.get_vertical_path(currentSquare, newSquare))
            
            } else { return false } 
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

//valid_knight_move determines if this is a valid knight move.
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

//valid_castle_move determines if this is a valid castle move. pathToNewSquare is the row/column that the castle is moving along.
func valid_castle_move(oldX int, oldY int, newX int, newY int, pathToNewSquare []*Square) bool{
	for _,element := range pathToNewSquare{
		fmt.Println(*element)

	}
	/**for _,square := range pathToNewSquare{
		if square.X == newX && square.Y == newY{  //arrived at destination square. Check if a piece is currently there
			if square.HasPiece == { return false }
			return true
		}
	}*/
	return false

}