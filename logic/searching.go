package search

import "github.com/NabilHilmi21/alpro2_final_project/models"

/*
check apakah user ada apa engga menggunakan SEQUENTIAL SEARCH
dan mengembalikan INDEX user tersebut
*/

func check_user(T models.Tab_users, n int, x string) int {
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

func check_funds(T models.Tab_funds, n int, x string) int {
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
