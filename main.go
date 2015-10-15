package main

import (
		"fmt"
		objects "Go_chess_AI_engine/chess_objects"	
		)

func main(){

	board := objects.Board{}
	board.New()
    
	for _,element := range board.Pieces {
        if element.Location.X == 4{
            fmt.Println(element.Type)
        }
    }
}