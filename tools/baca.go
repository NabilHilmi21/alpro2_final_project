package baca

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/search"
	"github.com/NabilHilmi21/alpro2_final_project/storage"
)

func Baca_funds() {
	var funds_nama string
	var jumlah int = 0

	for {
		fmt.Print("Nama tempat danamu ('-' untuk berhenti): ")
		fmt.Scan(&funds_nama)

		if funds_nama == "-" {
			break
		}

		if search.Check_funds(storage.Arr_funds, storage.Funds_count, funds_nama) != -1 {
			fmt.Println("Tempat dana dengan nama tersebut sudah ada. Coba lagi yang lain")
			continue
		}

		storage.Arr_funds[storage.Funds_count].Id, storage.Arr_funds[storage.Funds_count].Nama = storage.Funds_count, funds_nama

		storage.Funds_count++
		jumlah++
	}

	fmt.Println("")
	fmt.Println(jumlah, "fund berhasil ditambahkan")
	fmt.Println("")
}

func Baca_users() {
	var user_nama string
	var jumlah int = 0

	for {
		fmt.Print("Enter username ('-' untuk berhenti): ")
		fmt.Scan(&user_nama)

		if user_nama == "-" {
			break
		}

		if search.Check_user(storage.Arr_users, storage.Users_count, user_nama) != -1 || user_nama == "-" {
			fmt.Println("User dengan username tersebut sudah ada. Coba lagi yang lain")
			continue
		}

		storage.Arr_users[storage.Users_count].Id, storage.Arr_users[storage.Users_count].Nama = storage.Users_count, user_nama

		storage.Users_count++
		jumlah++
	}

	fmt.Println("")
	fmt.Println(jumlah, "user berhasil ditambahkan")
	fmt.Println("")

}
