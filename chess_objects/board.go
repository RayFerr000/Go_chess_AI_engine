package chess_objects



type Board struct{

	Board [12][12] Square
	 
}

func (b *Board) New(){
    for i := 0; i <= 11; i++{
    	for j := 0; j <= 11; j++{
	    
	    	//Initialize out-of-bounds region 
	    	if i == 0 || i == 1 || i == 10 || i == 11 || j == 0 || j == 1 || j == 10 || j == 11 {
	            b.Board[i][j] = Square{ OutOfBounds : true }
            } else{
	        	b.Board[i][j] = Square{ OutOfBounds : false, HasPiece : true, piece : Piece{ Type : 1 }, X : j, Y : i }
            }
    	}
    }
}