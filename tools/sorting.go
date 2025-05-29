// semua sorting funds menggunakan SELECTION SORT
package sorting

import (
	"github.com/NabilHilmi21/alpro2_final_project/storage"
	"github.com/NabilHilmi21/alpro2_final_project/structs"
)

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
