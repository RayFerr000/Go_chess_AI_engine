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

        case math.Abs(float64(pieceType)) == 3:  //bishop
        	return validate_bishop_move(board.valid_diagonal_path(currentSquare, newSquare))

        case math.Abs(float64(pieceType)) == 4: //castle
        	if currentSquare.X == newSquare.X{  //horizontal castle move
        		return valid_castle_move(currentSquare, newSquare, board.get_row(currentSquare.X))
        	
        	} else if currentSquare.Y == newSquare.Y{  //vertical castle move
        		return valid_castle_move(currentSquare, newSquare, board.get_column(currentSquare.Y))
            
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
func valid_castle_move(currentSquare *Square, newSquare *Square, path [8]*Square) bool{
	i, j, k := 0,7,1  //Variables to control our loop.
	startSquare, endSquare := currentSquare, newSquare
	if currentSquare.Y > newSquare.Y || currentSquare.X > newSquare.X{  //Determining direction in which we are going to traverse the row/column.
		i, j, k = 7,0,-1		
	}
	for i != j {
		square := path[i]
		if square == startSquare{  //Found the first square in our path so now check to see if path is clear or not.
			for square != endSquare{	
				if (square != currentSquare) && (square.HasPiece == true){
					return false
				}
				i += k
				square = path[i]	
			}
			if square.HasPiece == false { return true }
		} else {
			i += k
			square = path[i]
		}
	}
	return false
}

func validate_bishop_move(validPath bool) bool{
	return validPath
}










