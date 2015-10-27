package chess_objects
import (
        "math"
        "fmt"
       )

type Board struct{

	Board [12][12] *Square
    Pieces  [] *Piece
	 
}

func (b *Board) New(){
    for i := 0; i <= 11; i++{
    	for j := 0; j <= 11; j++{
	        if i > 3 && i < 8 { 
                square := Square{ OutOfBounds: false, HasPiece : false, X : i, Y : j }
                b.Board[i][j] = &square
            }
	    	
	    	//Initialize out-of-bounds region 
	    	if i == 0 || i == 1 || i == 10 || i == 11 || j == 0 || j == 1 || j == 10 || j == 11 {
	            square := Square{ OutOfBounds : true, X : i, Y : j }
                b.Board[i][j] = &square
	            
            }else if i == 3{ //White Pawns
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
	        	b.Board[i][j] = &square
                piece := Piece{ Type : 1, Location : b.Board[i][j], IsWhite : true, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)   	
            
            }else if i == 8{ //Black pawns
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -1, Location : b.Board[i][j], IsWhite : false,  Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 2 && j == 2) || (i == 2 && j == 9) { //White Castle
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 4, Location : b.Board[i][j], IsWhite : true, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 9 && j == 2) || (i == 9 && j == 9){ //Black Castle
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -4, Location : b.Board[i][j], IsWhite : false, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
    	    
            }else if (i == 2 && j == 3) || (i == 2 && j == 8){ //White Horse
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 2, Location : b.Board[i][j], IsWhite : true, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 9 && j == 3) || (i == 9 && j == 8){ //Black Horse
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -2, Location : b.Board[i][j], IsWhite : false, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 2 && j == 4) || (i == 2 && j == 7){ //White Bishop
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 3, Location : b.Board[i][j], IsWhite : true, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 9 && j == 4) || (i == 9 && j == 7){ //Black Bishop
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -3, Location : b.Board[i][j], IsWhite : false, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 2 && j == 5){ //White Queen
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 5, Location : b.Board[i][j], IsWhite : true, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 9 && j == 5){ //Black Queen
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -5, Location : b.Board[i][j], IsWhite : false, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 2 && j == 6){ //White King
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 6, Location : b.Board[i][j], IsWhite : true, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 9 && j == 6){ //Black King
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -6, Location : b.Board[i][j], IsWhite : false, Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      }
        }
    }
}
//get_column returns a specific column from the board.
func (b *Board) get_column(column int) [8]*Square{
    var tmp [8]*Square
    pos := 0
    for i:= 2; i <= 9; i++{
        tmp[pos] = b.Board[i][column]
        pos++
    }
    return tmp
}
//get_row returns a specific from the board.
func (b *Board) get_row(row int) [8]*Square{
    var tmp [8]*Square
    pos := 0
    for i:= 2; i <= 9; i++{
        tmp[pos] = b.Board[row][i]
        pos++
    }
    return tmp    
}
// valid_diagonal_path will determine is a diagonal path is a valid move. Used for bishops and queens.
func (b *Board) valid_diagonal_path(currentSquare *Square, newSquare *Square) bool{
    if ( math.Abs(float64(newSquare.X - currentSquare.X)) != math.Abs(float64(newSquare.Y - currentSquare.Y)) ){  // bishop paths must be of slope = +1 or -1
        return false
    } 
    xDirection := 1
    yDirection := 1
    if newSquare.X - currentSquare.X < 0 { xDirection = -1 }  // determine which direction this piece is moving
    if newSquare.Y - currentSquare.Y < 0 { yDirection = -1 }

    i := currentSquare.X
    j := currentSquare.Y
    
    square := b.Board[i + xDirection][j + yDirection]
    for square != newSquare {
        if ( square.HasPiece == false ){
            i = square.X
            j = square.Y
            square = b.Board[i + xDirection][j + yDirection]
        } else {
            fmt.Println("Not a valid bishop move")
            return false
        }
    }
    return true
}


