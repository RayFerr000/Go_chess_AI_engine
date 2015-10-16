package chess_objects



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

