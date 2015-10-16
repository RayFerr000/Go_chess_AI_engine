package chess_objects

type Piece struct{
    
    Type int
    Location *Square
    Captured bool

}
/**
* Move the current piece to a different square.
* Args:
*	s: The destination square.
*/
func (p *Piece) Move_piece(s *Square){
	if s.HasPiece || s.OutOfBounds {
		return
	}
	oldSquare := p.Location
	s.Piece = p
	s.HasPiece = true
	p.Location = s
	oldSquare.HasPiece = false
}