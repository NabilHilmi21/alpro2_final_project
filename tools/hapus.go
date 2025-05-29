package hapus

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/search"
	"github.com/NabilHilmi21/alpro2_final_project/storage"
)

func Hapus_users(x string) {
	var index int

	index = search.Check_user(storage.Arr_users, storage.Users_count, x)

	if index == -1 {
		fmt.Println("Tidak ada user yang bernama", x)
		return
	}

	for i := index; i < storage.Users_count-1; i++ {
		storage.Arr_users[i] = storage.Arr_users[i+1]
	}

	storage.Users_count--

	fmt.Println("User bernama", x, "berhasil dihapus.")

}

func Hapus_funds(x string) {
	var index int

	index = search.Check_funds(storage.Arr_funds, storage.Funds_count, x)

	if index == -1 {
		fmt.Println("Tidak ada fund yang bernama", x)
		return
	}

	for i := index; i < storage.Funds_count-1; i++ {
		storage.Arr_funds[i] = storage.Arr_funds[i+1]
	}

	storage.Funds_count--

	fmt.Println("Fund", x, "berhasil dihapus.")

}
