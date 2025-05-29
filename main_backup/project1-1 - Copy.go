package main

import "fmt"

const NMAX int = 99

// CHECK LINE 619
// CHECK LINE 619
// CHECK LINE 619
// CHECK LINE 619
// CHECK LINE 619

type users struct {
	id    int
	total int
	nama  string
}

type funds struct {
	id      int
	nama    string
	nominal int
	total   int
}

type riwayat struct {
	id      int
	nama    string
	nominal int
	tanggal string
	tempat  string
}

type tab_funds [NMAX]funds
type tab_users [NMAX]users
type tab_riwayat [NMAX]riwayat

var arr_funds tab_funds
var arr_users tab_users
var arr_riwayat tab_riwayat

var funds_count int = 0
var users_count int = 0
var riwayat_count int = 0

var current_user string

// buat tampilin menu utama
func open_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%10sWelcome to TelkomFunds!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Login as: ")
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Println("0. Exit")
	fmt.Println("==============================================")
	fmt.Print("--- Input Your Choice: ")

	var choice int

	fmt.Scan(&choice)
	fmt.Println("")

	return choice
}

func admin_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%8sWelcome to TelkomFunds Admin!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Menu: ")
	fmt.Println("1. View all users")
	fmt.Println("2. View all funds")
	fmt.Println("3. Add new user")
	fmt.Println("4. Add new fund")
	fmt.Println("5. View all transaction history")
	fmt.Println("0. Kebali ke menu utama")
	fmt.Println("==============================================")
	fmt.Print("Your input: ")

	var choice int

	fmt.Scan(&choice)

	return choice
}

func user_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%9sWelcome to TelkomFunds User!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Menu: ")
	fmt.Println("1. View my transactions")
	fmt.Println("2. View available funds")
	fmt.Println("0. Kembali ke menu utama")
	fmt.Print("Your input: ")

	var choice int
	fmt.Scan(&choice)

	fmt.Println("==============================================")
	return choice
}

func users_option_admin() int {
	var choice int

	fmt.Println("--- MENU USERS ADMIN")
	fmt.Println("1. Hapus user")
	fmt.Println("2. Urutkan sesuai ID")
	fmt.Println("3. Urutkan sesuai total donasi")
	fmt.Println("0. Kembali")
	fmt.Print("--- Pilihanmu: ")

	fmt.Scan(&choice)

	return choice
}

func funds_option_admin() int {
	var choice int

	fmt.Println("--- MENU FUNDS ADMIN")
	fmt.Println("1. Hapus tempat donasi")
	fmt.Println("2. Urutkan sesuai ID")
	fmt.Println("3. Urutkan sesuai total donasi")
	fmt.Println("0. Kembali")
	fmt.Print("--- Pilihanmu: ")

	fmt.Scan(&choice)

	return choice
}

func funds_option() int {
	var choice int

	fmt.Println("--- MENU FUNDS USER")
	fmt.Println("1. Donasi ke tempat dana pilihamu")
	fmt.Println("2. Urutkan sesuai ID")
	fmt.Println("3. Urutkan sesuai total donasi")
	fmt.Println("0. Kembali")
	fmt.Print("--- Pilihanmu: ")

	fmt.Scan(&choice)

	return choice
}

// SEMUA FUNCTION SEARCHING

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

func hapus_users(x string) {
	var index int

	index = check_user(arr_users, users_count, x)

	if index == -1 {
		fmt.Println("Tidak ada user yang bernama", x)
		return
	}

	for i := index; i < users_count-1; i++ {
		arr_users[i] = arr_users[i+1]
	}

	users_count--

	fmt.Println("User bernama", x, "berhasil dihapus.")

}

func hapus_funds(x string) {
	var index int

	index = check_funds(arr_funds, funds_count, x)

	if index == -1 {
		fmt.Println("Tidak ada fund yang bernama", x)
		return
	}

	for i := index; i < funds_count-1; i++ {
		arr_funds[i] = arr_funds[i+1]
	}

	funds_count--

	fmt.Println("Fund", x, "berhasil dihapus.")

}

// SEMUA FUNCTION SORTING

// semua sorting funds menggunakan SELECTION SORT
func urutMenaikFundsID(A *tab_funds, n int) {
	var pass, idx, i int
	var temp funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].id > A[i].id {
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

func urutMenurunFundsID(A *tab_funds, n int) {
	var pass, idx, i int
	var temp funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].id < A[i].id {
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

func urutMenaikFundsTotal(A *tab_funds, n int) {
	var pass, idx, i int
	var temp funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].total > A[i].total {
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

func urutMenurunFundsTotal(A *tab_funds, n int) {
	var pass, idx, i int
	var temp funds

	pass = 1

	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].total < A[i].total {
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

// semua sorting users menggunakan INSERTION SORT

func user_auth() bool {
	var logged_in bool = false
	var username string
	var i int = 1

	for i <= 5 && logged_in == false {
		fmt.Print("Enter username: ")
		fmt.Scan(&username)

		if check_user(arr_users, users_count, username) != -1 {
			fmt.Println("Kamu telah login sebagai", username)
			fmt.Println("")

			current_user = username

			logged_in = true
			return logged_in
		}

		fmt.Println("User not found, you have", 5-i, "tries left")
		fmt.Println("")
		i++
	}

	return logged_in
}

func admin_auth() bool {
	var logged_in bool = false
	var username, password string
	var i int = 1

	for i <= 5 && logged_in == false {
		fmt.Print("Enter username: ")
		fmt.Scan(&username)

		fmt.Print("Enter password: ")
		fmt.Scan(&password)

		fmt.Println("")

		if username == "admin123" && password == "admin123" {
			fmt.Println("Admin login successful")
			fmt.Println("")
			logged_in = true
			return logged_in
		}

		fmt.Println("Wrong credentials, You have", 5-i, "tries left")
		fmt.Println("")
		i++
	}

	return logged_in
}

func riwayat_funds(idx int, user string, donasi int, tempat string, tanggal string) {
	arr_riwayat[riwayat_count].id = arr_users[idx].id
	arr_riwayat[riwayat_count].nama = user
	arr_riwayat[riwayat_count].nominal = donasi
	arr_riwayat[riwayat_count].tempat = tempat
	arr_riwayat[riwayat_count].tanggal = tanggal

	riwayat_count++
}

func tampil_riwayat() {
	if riwayat_count <= 0 {
		fmt.Println("Tidak ada riwayat donasi")
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Printf("%-8s | %-20s | %-20s | %-20s | %-20s\n", "User ID", "Nama", "Donasi", "Tempat Donasi", "Tanggal")
		fmt.Printf("-----------------------------------------------------------\n")

		for i := 0; i < riwayat_count; i++ {
			fmt.Printf("%-8d | %-20s | %-20d | %-20s | %-20s\n", arr_riwayat[i].id, arr_riwayat[i].nama, arr_riwayat[i].nominal, arr_riwayat[i].tempat, arr_riwayat[i].tanggal)
		}

		fmt.Printf("-----------------------------------------------------------\n")
		fmt.Println("")
	}
}

func donate_funds(user string) {
	var choice, tanggal string
	var donasi, index_fund, index_user int

	for {
		fmt.Print("Pilih tempat danamu ('-' untuk kembali): ")
		fmt.Scan(&choice)

		index_fund = check_funds(arr_funds, funds_count, choice)
		index_user = check_user(arr_users, users_count, user)

		if index_fund != -1 {
			fmt.Print("Berapa yang anda ingin donate: ")
			fmt.Scan(&donasi)

			fmt.Print("Tanggal berapa anda donate (dd-mm-yyyy): ")
			fmt.Scan(&tanggal)

			arr_funds[index_fund].total += donasi
			arr_users[index_user].total += donasi

			riwayat_funds(index_user, user, donasi, choice, tanggal)

		} else if choice == "-" {
			fmt.Println("Kembali...")
			return
		} else {
			fmt.Println("Tidak ada tempat dana yang bernama", choice, ", coba lagi")
		}
	}

}

func main_handler(choice int) {
	var admin_log bool = false
	var user_log bool = false
	var admin_choice, user_choice int

	switch choice {
	case 1:
		fmt.Println("Logging in as admin...")
		fmt.Println("")

		if admin_auth() {
			admin_log = true

			for admin_log {
				admin_choice = admin_gui()

				if admin_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					admin_log = false
				} else {
					admin_handler(admin_choice)
				}
			}
		} else {
			fmt.Println("Login gagal. Kembali ke menu utama.")
		}
	case 2:
		fmt.Println("Login user...")
		if user_auth() {
			user_log = true

			for user_log {
				user_choice = user_gui()

				if user_choice == 0 {
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					user_log = false
				} else {
					user_handler(user_choice)
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

func funds_handler(choice int) {
	switch choice {
	case 1:
		donate_funds(current_user)
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
				urutMenaikFundsID(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tampil_funds()
			case 2:
				urutMenurunFundsID(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tampil_funds()
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
				urutMenaikFundsTotal(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil_funds()
			case 2:
				urutMenurunFundsTotal(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil_funds()
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

func funds_handler_admin(choice int) {
	switch choice {
	case 1:
		var nama_dana string
		fmt.Print("Ketik nama dana yang ingin dihapus: ")
		fmt.Scan(&nama_dana)
		hapus_funds(nama_dana)
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
				urutMenaikFundsID(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tampil_funds()
			case 2:
				urutMenurunFundsID(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tampil_funds()
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
				urutMenaikFundsTotal(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil_funds()
			case 2:
				urutMenurunFundsTotal(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil_funds()
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

func users_handler_admin(choice int) {
	switch choice {
	case 1:
		var nama_user string
		fmt.Print("Ketik nama user yang ingin dihapus: ")
		fmt.Scan(&nama_user)
		hapus_users(nama_user)
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
				urutMenaikFundsID(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terkecil.")
				tampil_funds()
			case 2:
				urutMenurunFundsID(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan ID terbesar.")
				tampil_funds()
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
				urutMenaikFundsTotal(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil_funds()
			case 2:
				urutMenurunFundsTotal(&arr_funds, funds_count)
				fmt.Println("Dana berhasil diurutkan berdasarkan total donasi terkecil.")
				tampil_funds()
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

func user_handler(choice int) {
	switch choice {
	case 1:
		fmt.Println("Viewing transactions")
	case 2:
		var show_gui bool = true
		var funds_choice int

		if funds_count == 0 {
			fmt.Println("")
			fmt.Println("Tidak ada dana tersedia.")
			fmt.Println("")

			show_gui = false
		} else {
			for show_gui {
				tampil_funds()
				funds_choice = funds_option()

				if funds_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					show_gui = false
				} else {
					funds_handler(funds_choice)
				}
			}
		}
	default:
		fmt.Println("Input tidak valid.")
	}
}

func admin_handler(choice int) {
	switch choice {
	case 1:
		var show_gui bool = true
		var users_choice int

		if users_count == 0 {
			fmt.Println("")
			fmt.Println("Tidak ada user.")
			fmt.Println("")

			show_gui = false
		} else {
			for show_gui {
				tampil_users()
				users_choice = users_option_admin()

				if users_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")

					show_gui = false
				} else {
					users_handler_admin(users_choice)
				}
			}
		}

	case 2:
		var show_gui bool = true
		var funds_choice int

		if funds_count == 0 {
			fmt.Println("")
			fmt.Println("Tidak ada dana tersedia.")
			fmt.Println("")

			show_gui = false
		} else {
			for show_gui {
				tampil_funds()
				funds_choice = funds_option_admin()

				if funds_choice == 0 {
					fmt.Println("")
					fmt.Println("Kembali ke menu utama")
					fmt.Println("")
					show_gui = false
				} else {
					funds_handler_admin(funds_choice)
				}
			}
		}
	case 3:
		fmt.Println("")
		baca_users()
	case 4:
		fmt.Println("")
		baca_funds()
	case 5:
		fmt.Println("")
		tampil_riwayat()
	default:
		fmt.Println("Input tidak valid")
	}
}

func baca_funds() {
	var funds_nama string
	var jumlah int = 0

	for {
		fmt.Print("Nama tempat danamu ('-' untuk berhenti): ")
		fmt.Scan(&funds_nama)

		if funds_nama == "-" {
			break
		}

		if check_funds(arr_funds, funds_count, funds_nama) != -1 {
			fmt.Println("Tempat dana dengan nama tersebut sudah ada. Coba lagi yang lain")
			continue
		}

		arr_funds[funds_count].id, arr_funds[funds_count].nama = funds_count, funds_nama

		funds_count++
		jumlah++
	}

	fmt.Println("")
	fmt.Println(jumlah, "fund berhasil ditambahkan")
	fmt.Println("")
}

func baca_users() {
	var user_nama string
	var jumlah int = 0

	for {
		fmt.Print("Enter username ('-' untuk berhenti): ")
		fmt.Scan(&user_nama)

		if user_nama == "-" {
			break
		}

		if check_user(arr_users, users_count, user_nama) != -1 || user_nama == "-" {
			fmt.Println("User dengan username tersebut sudah ada. Coba lagi yang lain")
			continue
		}

		arr_users[users_count].id, arr_users[users_count].nama = users_count, user_nama

		users_count++
		jumlah++
	}

	fmt.Println("")
	fmt.Println(jumlah, "user berhasil ditambahkan")
	fmt.Println("")

}

func tampil_funds() {
	fmt.Println("")
	fmt.Printf("%-8s | %-20s | %-20s\n", "ID Funds", "Nama", "Total Uang")
	fmt.Printf("-----------------------------------------------------------\n")

	for i := 0; i < funds_count; i++ {
		fmt.Printf("%-8d | %-20s | %-20d\n", arr_funds[i].id, arr_funds[i].nama, arr_funds[i].total)
	}

	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Println("")
}

func tampil_users() {
	fmt.Println("")
	fmt.Printf("%-8s | %-20s | %-20s\n", "User ID", "Nama", "Total Donasi")
	fmt.Printf("----------------------------------------------\n")
	for i := 0; i < users_count; i++ {
		fmt.Printf("%-8d | %-20s | %-20d\n", arr_users[i].id, arr_users[i].nama, arr_users[i].total)
	}
	fmt.Printf("----------------------------------------------\n")
	fmt.Println("")
}

func main() {
	var choice int

	for {
		choice = open_gui()
		main_handler(choice)
		if choice == 0 {
			break
		}
	}

}
