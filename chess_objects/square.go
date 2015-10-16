package chess_objects

type Square struct{
	OutOfBounds bool
	HasPiece bool
	Piece *Piece
	X int
	Y int
}