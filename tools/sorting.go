// semua sorting funds menggunakan SELECTION SORT
package tools

import (
	"github.com/NabilHilmi21/alpro2_final_project/storage"
	"github.com/NabilHilmi21/alpro2_final_project/structs"
)

// SORTING UNTUK FUNDS!!!!!!
// SEMUA SORTING FUNDS MENGGUNAKAN SELECTION SORT

func UrutMenaikFundsID(A *storage.Tab_funds, n int) {
	var pass, idx, i int
	var temp structs.Funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].Id > A[i].Id {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func UrutMenurunFundsID(A *storage.Tab_funds, n int) {
	var pass, idx, i int
	var temp structs.Funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].Id < A[i].Id {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func UrutMenaikFundsTotal(A *storage.Tab_funds, n int) {
	var pass, idx, i int
	var temp structs.Funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].Total > A[i].Total {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func UrutMenurunFundsTotal(A *storage.Tab_funds, n int) {
	var pass, idx, i int
	var temp structs.Funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].Total < A[i].Total {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

// SORTING UNTUK USERS!!!!!!
// SEMUA SORTING USERS MENGGUNAKAN INSERTION SORT

func UrutMenaikUsersID(A *storage.Tab_users, n int) {
	var pass, i, temp int

	pass = 1
	for pass < n-1 {
		i = pass
		temp = A[pass].Id

		for i > 0 && temp < A[i].Id {
			A[i].Id = A[i-1].Id
			i--
		}
		A[i].Id = temp
		pass++
	}
}

func UrutMenurunUsersID(A *storage.Tab_users, n int) {
	var pass, i, temp int

	pass = 1
	for pass < n-1 {
		i = pass
		temp = A[pass].Id

		for i > 0 && temp > A[i].Id {
			A[i].Id = A[i-1].Id
			i--
		}
		A[i].Id = temp
		pass++
	}
}

func UrutMenaikUsersTotal(A *storage.Tab_users, n int) {
	var pass, i, temp int

	pass = 1
	for pass < n-1 {
		i = pass
		temp = A[pass].Total

		for i > 0 && temp < A[i].Total {
			A[i].Total = A[i-1].Total
			i--
		}
		A[i].Total = temp
		pass++
	}
}

func UrutMenurunUsersTotal(A *storage.Tab_users, n int) {
	var pass, i, temp int

	pass = 1
	for pass < n-1 {
		i = pass
		temp = A[pass].Total

		for i > 0 && temp > A[i].Total {
			A[i].Total = A[i-1].Total
			i--
		}
		A[i].Total = temp
		pass++
	}
}
