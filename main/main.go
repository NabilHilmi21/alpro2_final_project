package main

import (
	"github.com/NabilHilmi21/alpro2_final_project/tampilan"

	"github.com/NabilHilmi21/alpro2_final_project/handlers"
)

// SEMUA FUNCTION SORTIN

// semua sorting users menggunakan INSERTION SORT

/* TAMBAHAN:
- REGISTER USER
- KONSISTEN BAHASA!!!!!!!
- RAPIHKAN GUI INPUT
*/

func main() {
	var choice int

	for {
		choice = tampilan.Open_gui()
		handlers.Main_handler(choice)
		if choice == 0 {
			break
		}
	}
}
