package tools

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"
)

func Donate_funds(user string) {
	var choice, tanggal string
	var donasi, index_fund, index_user int

	for {
		fmt.Println("")
		fmt.Print("--- Pilih tempat danamu ('-' untuk kembali): ")
		fmt.Scan(&choice)

		index_fund = Check_funds(storage.Arr_funds, storage.Funds_count, choice)
		index_user = Check_user(storage.Arr_users, storage.Users_count, user)

		if index_fund != -1 {
			fmt.Print("--- Berapa yang anda ingin donate: ")
			fmt.Scan(&donasi)

			fmt.Print("--- Tanggal berapa anda donate (dd-mm-yyyy): ")
			fmt.Scan(&tanggal)
			fmt.Println("")

			storage.Arr_funds[index_fund].Total += donasi
			storage.Arr_users[index_user].Total += donasi

			storage.Riwayat_funds(index_user, user, donasi, choice, tanggal)

		} else if choice == "-" {
			fmt.Println("Kembali...")
			return
		} else {
			fmt.Println("Tidak ada tempat dana yang bernama", choice, ", coba lagi")
		}
	}

}
