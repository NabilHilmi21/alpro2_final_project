package main

import (
	"fmt"

	// dari folder "storage"
	"github.com/NabilHilmi21/alpro2_final_project/storage"

	// dari folder "structs"

	//dari folder "tampilan"
	"github.com/NabilHilmi21/alpro2_final_project/tampilan"

	// dari folder "tools"
	"github.com/NabilHilmi21/alpro2_final_project/search"
	// dari folder "handlers"
	// dari folder "auth"
)

// SEMUA FUNCTION SORTIN

// semua sorting users menggunakan INSERTION SORT

func riwayat_funds(idx int, user string, donasi int, tempat string, tanggal string) {
	storage.Arr_riwayat[storage.Riwayat_count].Id = storage.Arr_users[idx].Id
	storage.Arr_riwayat[storage.Riwayat_count].Nama = user
	storage.Arr_riwayat[storage.Riwayat_count].Nominal = donasi
	storage.Arr_riwayat[storage.Riwayat_count].Tempat = tempat
	storage.Arr_riwayat[storage.Riwayat_count].Tanggal = tanggal

	storage.Riwayat_count++
}

func donate_funds(user string) {
	var choice, tanggal string
	var donasi, index_fund, index_user int

	for {
		fmt.Print("Pilih tempat danamu ('-' untuk kembali): ")
		fmt.Scan(&choice)

		index_fund = search.Check_funds(storage.Arr_funds, storage.Funds_count, choice)
		index_user = search.Check_user(storage.Arr_users, storage.Users_count, user)

		if index_fund != -1 {
			fmt.Print("Berapa yang anda ingin donate: ")
			fmt.Scan(&donasi)

			fmt.Print("Tanggal berapa anda donate (dd-mm-yyyy): ")
			fmt.Scan(&tanggal)

			storage.Arr_funds[index_fund].Total += donasi
			storage.Arr_users[index_user].Total += donasi

			riwayat_funds(index_user, user, donasi, choice, tanggal)

		} else if choice == "-" {
			fmt.Println("Kembali...")
			return
		} else {
			fmt.Println("Tidak ada tempat dana yang bernama", choice, ", coba lagi")
		}
	}

}

func main() {
	var choice int

	for {
		choice = tampilan.Open_gui()
		admin_handler.Main_handler(choice)
		if choice == 0 {
			break
		}
	}

}
