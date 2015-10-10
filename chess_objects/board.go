package chess_objects



type Board struct{

	Board [12][12] Square
	 
}

func (b *Board) New(){
    for i := 0; i <= 11; i++{
    	for j := 0; j <= 11; j++{
	        for i > 3 && i < 8 { i++ }
	    	
	    	//Initialize out-of-bounds region 
	    	if i == 0 || i == 1 || i == 10 || i == 11 || j == 0 || j == 1 || j == 10 || j == 11 {
	            b.Board[i][j] = Square{ OutOfBounds : true }
	            
            }else if i == 3{ //White Pawns
	        	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : 1 }, X : i, Y : j }
	        	
            }else if i == 8{ //Black pawns
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : -1 }, X : i, Y : j }
            	
            }else if (i == 2 && j == 2) || (i == 2 && j == 9) { //White Castle
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : 4 }, X : i, Y : j }
            	
            }else if (i == 9 && j == 2) || (i == 9 && j == 9){ //Black Castle
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : -4 }, X : i, Y : j }
            	
    	    }else if (i == 2 && j == 3) || (i == 2 && j == 8){ //White Horse
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type :  2 }, X : i, Y : j }
            	
            }else if (i == 9 && j == 3) || (i == 9 && j == 8){ //Black Horse
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : -2 }, X : i, Y : j }
            	
            }else if (i == 2 && j == 4) || (i == 2 && j == 7){ //White Bishop
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : 3 }, X : i, Y : j }
            	
            }else if (i == 9 && j == 4) || (i == 9 && j == 7){ //Black Bishop
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : -3 }, X : i, Y : j }
            
            }else if (i == 2 && j == 5){ //White Queen
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : 5 }, X : i, Y : j }
            
            }else if (i == 5 && j == 5){ //Black Queen
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : -5 }, X : i, Y : j }
            
            }else if (i == 2 && j == 6){ //White King
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : 6 }, X : i, Y : j }
            
            }else if (i == 9 && j == 6){ //Black King
            	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, Piece_ : Piece{ Type : -6 }, X : i, Y : j }
            }
        }
    }
}