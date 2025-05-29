package handlers

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"

	"github.com/NabilHilmi21/alpro2_final_project/tampilan"

	"github.com/NabilHilmi21/alpro2_final_project/tools"
)

func Funds_handler(choice int) {
	switch choice {
	case 1:
		tools.Donate_funds(storage.Current_user)
	case 2:
		var pilih int
		for {
			fmt.Println("1. Urutkan dari ID terkecil")
			fmt.Println("2. Urutkan dari ID terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihanmu: ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				tools.UrutMenaikFundsID(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tools.Tampil_funds()
			case 2:
				tools.UrutMenurunFundsID(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tools.Tampil_funds()
			case 0:
				return
			default:
				fmt.Println("Input tidak valid, coba lagi.")
			}
		}

	case 3:
		var pilih int
		for {
			fmt.Println("1. Urutkan dari total donasi terkecil")
			fmt.Println("2. Urutkan dari total donasi terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihanmu: ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				tools.UrutMenaikFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tools.Tampil_funds()
			case 2:
				tools.UrutMenurunFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tools.Tampil_funds()
			case 0:
				return
			default:
				fmt.Println("Input tidak valid, coba lagi.")
			}
		}
	case 0:
		break
	default:
		fmt.Println("Input tidak valid, coba lagi.")
	}
}

func User_handler(choice int) {
	switch choice {
	case 1:
		fmt.Println("Viewing transactions")
	case 2:
		var show_gui bool = true
		var funds_choice int

		if storage.Funds_count == 0 {
			fmt.Println("")
			fmt.Println("Tidak ada dana tersedia.")
			fmt.Println("")

			show_gui = false
		} else {
			for show_gui {
				tools.Tampil_funds()
				funds_choice = tampilan.Funds_option()

				if funds_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					show_gui = false
				} else {
					Funds_handler(funds_choice)
				}
			}
		}
	default:
		fmt.Println("Input tidak valid.")
	}
}
