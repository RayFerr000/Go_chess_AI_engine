package chess_objects

import "fmt"

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
                piece := Piece{ Type : 1, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)   	
            
            }else if i == 8{ //Black pawns
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -1, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 2 && j == 2) || (i == 2 && j == 9) { //White Castle
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 4, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 9 && j == 2) || (i == 9 && j == 9){ //Black Castle
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -4, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
    	    
            }else if (i == 2 && j == 3) || (i == 2 && j == 8){ //White Horse
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 2, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 9 && j == 3) || (i == 9 && j == 8){ //Black Horse
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -2, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 2 && j == 4) || (i == 2 && j == 7){ //White Bishop
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 3, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      	
            
            }else if (i == 9 && j == 4) || (i == 9 && j == 7){ //Black Bishop
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -3, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 2 && j == 5){ //White Queen
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 5, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 9 && j == 5){ //Black Queen
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -5, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 2 && j == 6){ //White King
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : 6, Location : b.Board[i][j], Captured : false }
                square.Piece = &piece
                b.Pieces = append(b.Pieces,&piece)      
            
            }else if (i == 9 && j == 6){ //Black King
                square := Square{ OutOfBounds : false, HasPiece : true, X : i, Y : j }
                b.Board[i][j] = &square
                piece := Piece{ Type : -6, Location : b.Board[i][j], Captured : false }
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
//get_horizontal_path goes to the row the s1 and s2 are on and finds all squares that are between them.
func (b*Board) get_horizontal_path(s1 *Square , s2 *Square) []*Square{
    var low, hi int = s1.Y, s2.Y  //variables to grab slice.
    
    if s1.Y > s2.Y{ 
        low, hi = s2.Y, s1.Y
    }
    
    row := b.get_row(s1.X)
    var tmp []*Square = row[low:hi]
    return tmp
}
//get_vertical_path goes to the column the s1 and s2 are on and finds all squares that are between them.
func (b*Board) get_vertical_path(s1 *Square , s2 *Square) []*Square{
    var i, j, k = 0,0,1  //Variables to control our loop since we don't know which way we are traversing the column. Default is bottom -> top
    
    if s1.X > s2.X {  //Check for top -> bottom traversal
        i, k = 7,-1
    }
    tmp := []*Square
    row := b.get_column(s1.Y)
    for true{
        if (row[i].X == s1.X && row[i].Y == s1.Y) || (row[i].X == s2.X && row[i].Y == s2.Y){
            
        }
    }
    return tmp
}

