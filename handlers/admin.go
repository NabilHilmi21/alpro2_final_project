package handlers

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/auth"
	"github.com/NabilHilmi21/alpro2_final_project/storage"
	"github.com/NabilHilmi21/alpro2_final_project/tampilan"
	"github.com/NabilHilmi21/alpro2_final_project/tools"
)

func Main_handler(choice int) {
	var admin_log bool = false
	var user_log bool = false
	var admin_choice, user_choice int

	switch choice {
	case 1:
		fmt.Println("Sedang masuk sebagai admin...")
		fmt.Println("")

		if auth.Admin_auth() {
			admin_log = true

			for admin_log {
				admin_choice = tampilan.Admin_gui()

				if admin_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama.")
					fmt.Println("")
					admin_log = false
				} else {
					Admin_handler(admin_choice)
				}
			}
		} else {
			fmt.Println("Masuk gagal. Kembali ke menu utama.")
			fmt.Println("")
		}
	case 2:
		user_choice = tampilan.Users_login()

		switch user_choice {
		case 1:
			if auth.User_auth() {
				user_log = true

				for user_log {
					var login_choice int = tampilan.User_gui()

					if login_choice == 0 {
						fmt.Println("Kembali ke menu utama.")
						fmt.Println("")
						user_log = false
					} else {
						User_handler(login_choice)
					}
				}
			} else {
				fmt.Println("Masuk gagal. Kembali ke menu utama.")
				fmt.Println("")
			}

		case 2:
			auth.User_reg()
		}

	case 0:
		fmt.Println("Keluar dari menu...")
		fmt.Println("")
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		fmt.Println("")
	}
}

func Admin_handler(choice int) {
	switch choice {
	case 1:
		var show_gui bool = true
		var users_choice int

		if storage.Users_count == 0 {
			fmt.Println("")
			fmt.Println("Tidak ada pengguna.")
			fmt.Println("")

			show_gui = false
		} else {
			for show_gui {
				tools.Tampil_users()
				users_choice = tampilan.Users_option_admin()

				if users_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama.")
					fmt.Println("")

					show_gui = false
				} else {
					Users_handler_admin(users_choice)
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
				tools.Tampil_funds()
				funds_choice = tampilan.Funds_option_admin()

				if funds_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama.")
					fmt.Println("")
					show_gui = false
				} else {
					Funds_handler_admin(funds_choice)
				}
			}
		}
	case 3:
		fmt.Println("")
		tools.Baca_users()
	case 4:
		fmt.Println("")
		tools.Baca_funds()
	case 5:
		fmt.Println("")
		tools.Tampil_riwayat()
	default:
		fmt.Println("Masukan tidak valid.")
		fmt.Println("")
	}
}

func Funds_handler_admin(choice int) {
	switch choice {
	case 1:
		var nama_dana string

		fmt.Println("")
		fmt.Print("Ketik nama dana yang ingin dihapus: ")
		fmt.Scan(&nama_dana)
		fmt.Println("")

		tools.Hapus_funds(nama_dana)
	case 2:
		var pilih int
		for {
			fmt.Println("1. Urutkan dari ID terkecil")
			fmt.Println("2. Urutkan dari ID terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihan Anda: ")
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
				fmt.Println("")
				return
			default:
				fmt.Println("Masukan tidak valid, silakan coba lagi.")
				fmt.Println("")
			}
		}

	case 3:
		var pilih int
		for {
			fmt.Println("1. Urutkan dari total donasi terkecil")
			fmt.Println("2. Urutkan dari total donasi terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				tools.UrutMenaikFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tools.Tampil_funds()
			case 2:
				tools.UrutMenurunFundsTotal(&storage.Arr_funds, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terbesar.")
				tools.Tampil_funds()
			case 0:
				fmt.Println("")
				return
			default:
				fmt.Println("Masukan tidak valid, silakan coba lagi.")
				fmt.Println("")
			}
		}
	case 0:
		break
	default:
		fmt.Println("Masukan tidak valid, silakan coba lagi.")
		fmt.Println("")
	}
}

func Users_handler_admin(choice int) {
	switch choice {
	case 1:
		var nama_user string
		fmt.Print("Ketik nama pengguna yang ingin dihapus: ")
		fmt.Scan(&nama_user)
		tools.Hapus_users(nama_user)
	case 2:

		/* BUAT tools BARU BUAT NGE SORT ID SAMA TOTAL DONASI PAKE CARA tools LAIN */

		var pilih int
		for {
			fmt.Println("")
			fmt.Println("--- MENU URUTKAN ID PENGGUNA")
			fmt.Println("1. Urutkan dari ID terkecil")
			fmt.Println("2. Urutkan dari ID terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("-- Masukkan Pilihan Anda: ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				tools.UrutMenaikUsersID(&storage.Arr_users, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tools.Tampil_users()
			case 2:
				tools.UrutMenurunUsersID(&storage.Arr_users, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tools.Tampil_users()
			case 0:
				fmt.Println("")
				return
			default:
				fmt.Println("Masukan tidak valid, silakan coba lagi.")
				fmt.Println("")
			}
		}

	case 3:
		var pilih int
		for {
			fmt.Println("")
			fmt.Println("--- MENU URUTKAN DONASI PENGGUNA")
			fmt.Println("1. Urutkan dari total donasi terkecil")
			fmt.Println("2. Urutkan dari total donasi terbesar")
			fmt.Println("0. Kembali")
			fmt.Print("--- Masukkan Pilihan Anda: ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				tools.UrutMenaikUsersTotal(&storage.Arr_users, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tools.Tampil_users()
			case 2:
				tools.UrutMenurunUsersTotal(&storage.Arr_users, storage.Funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terbesar.")
				tools.Tampil_users()
			case 0:
				fmt.Println("")
				return
			default:
				fmt.Println("Masukan tidak valid, silakan coba lagi.")
				fmt.Println("")
			}
		}
	case 0:
		break
	default:
		fmt.Println("Masukan tidak valid, silakan coba lagi.")
		fmt.Println("")
	}
}
