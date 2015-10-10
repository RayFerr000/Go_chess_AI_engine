package chess_objects

type Square struct{
	OutOfBounds bool
	HasPiece bool
	Piece_ Piece
	X int
	Y int
}