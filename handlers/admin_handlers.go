package admin_handlers

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"
	"github.com/NabilHilmi21/alpro2_final_project/tampilan"

	"github.com/NabilHilmi21/alpro2_final_project/user_handlers"

	"github.com/NabilHilmi21/alpro2_final_project/hapus"
)

func Main_handler(choice int) {
	var admin_log bool = false
	var user_log bool = false
	var admin_choice, user_choice int

	switch choice {
	case 1:
		fmt.Println("Logging in as admin...")
		fmt.Println("")

		if auth.Admin_auth() {
			admin_log = true

			for admin_log {
				admin_choice = tampilan.Admin_gui()

				if admin_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					admin_log = false
				} else {
					Admin_handler(admin_choice)
				}
			}
		} else {
			fmt.Println("Login gagal. Kembali ke menu utama.")
		}
	case 2:
		fmt.Println("Login user...")
		if admin.User_auth() {
			user_log = true

			for user_log {
				user_choice = tampilan.User_gui()

				if user_choice == 0 {
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					user_log = false
				} else {
					user_handlers.User_handler(user_choice)
				}
			}
		} else {
			fmt.Println("Login gagal. Kembali ke menu utama.")
		}
	case 0:
		fmt.Println("Exiting menu...")
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
	}
}

func Admin_handler(choice int) {
	switch choice {
	case 1:
		var show_gui bool = true
		var users_choice int

		if storage.Users_count == 0 {
			fmt.Println("")
			fmt.Println("Tidak ada user.")
			fmt.Println("")

			show_gui = false
		} else {
			for show_gui {
				tempil.Tampil_users()
				users_choice = tampilan.Users_option_admin()

				if users_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")

					show_gui = false
				} else {
					admin_handlers.Users_handler_admin(users_choice)
				}
			}
		}

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
				tampil.Tampil_funds()
				funds_choice = tampilan.Funds_option_admin()

				if funds_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					show_gui = false
				} else {
					admin_handlers.Funds_handler_admin(funds_choice)
				}
			}
		}
	case 3:
		fmt.Println("")
		baca.Baca_users()
	case 4:
		fmt.Println("")
		baca.Baca_funds()
	case 5:
		fmt.Println("")
		tampil.Tampil_riwayat()
	default:
		fmt.Println("Input tidak valid")
	}
}

func Funds_handler_admin(choice int) {
	switch choice {
	case 1:
		var nama_dana string
		fmt.Print("Ketik nama dana yang ingin dihapus: ")
		fmt.Scan(&nama_dana)
		hapus.Hapus_funds(nama_dana)
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
				sorting.UrutMenaikFundsID(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tampil.Tampil_funds()
			case 2:
				sorting.UrutMenurunFundsID(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tampil.Tampil_funds()
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
				sorting.UrutMenaikFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil.Tampil_funds()
			case 2:
				sorting.UrutMenurunFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil.Tampil_funds()
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

func Users_handler_admin(choice int) {
	switch choice {
	case 1:
		var nama_user string
		fmt.Print("Ketik nama user yang ingin dihapus: ")
		fmt.Scan(&nama_user)
		hapus.Hapus_users(nama_user)
	case 2:

		/* BUAT SORTING BARU BUAT NGE SORT ID SAMA TOTAL DONASI PAKE CARA SORTING LAIN */

		var pilih int
		for {
			fmt.Println("1. Urutkan dari ID terkecil")
			fmt.Println("2. Urutkan dari ID terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihanmu: ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				sorting.UrutMenaikFundsID(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tampil.Tampil_funds()
			case 2:
				sorting.UrutMenurunFundsID(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tampil.Tampil_funds()
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
				sorting.UrutMenaikFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil.Tampil_funds()
			case 2:
				sorting.UrutMenurunFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil.Tampil_funds()
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
