package search

/*
check apakah user ada apa engga menggunakan SEQUENTIAL SEARCH
dan mengembalikan INDEX user tersebut
*/

func check_user(T tab_users, n int, x string) int {
	var ketemu int = -1
	var k int = 0

	for ketemu == -1 && k < n {
		if T[k].nama == x {
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

func check_funds(T tab_funds, n int, x string) int {
	var ketemu int = -1
	var k int = 0

	for ketemu == -1 && k < n {
		if T[k].nama == x {
			ketemu = k
		}
		k++
	}
	return ketemu
}
