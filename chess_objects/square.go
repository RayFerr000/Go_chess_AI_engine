package chess_objects

type Square struct{
	OutOfBounds bool
	HasPiece bool
	piece Piece
	X int
	Y int
}