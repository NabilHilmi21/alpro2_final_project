package tools

import "github.com/NabilHilmi21/alpro2_final_project/storage"

/*
check apakah user ada apa engga menggunakan SEQUENTIAL SEARCH
dan mengembalikan INDEX user tersebut
*/

func Check_user(T storage.Tab_users, n int, x string) int {
	var ketemu int = -1
	var k int = 0

	for ketemu == -1 && k < n {
		if T[k].Nama == x {
			ketemu = k
		}
		k++
	}
	return ketemu
}

/*
check apakah funds ada apa engga menggunakan SEQUENTIAL SEARCH
dan mengembalikan INDEX funds tersebut
*/

func Check_funds(T storage.Tab_funds, n int, x string) int {
	var ketemu int = -1
	var k int = 0

	for ketemu == -1 && k < n {
		if T[k].Nama == x {
			ketemu = k
		}
		k++
	}
	return ketemu
}
