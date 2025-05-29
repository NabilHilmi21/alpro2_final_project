package tampilan

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"
)

// Menampilkan menu utama
func Open_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%10sSelamat Datang di TelkomFunds!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Masuk sebagai: ")
	fmt.Println("1. Admin")
	fmt.Println("2. Pengguna")
	fmt.Println("0. Keluar")
	fmt.Println("==============================================")
	fmt.Print("--- Masukkan Pilihan Anda: ")

	var choice int

	fmt.Scan(&choice)
	fmt.Println("")

	return choice
}

// Menampilkan menu Admin
func Admin_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%8sSelamat Datang, Admin TelkomFunds!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Menu: ")
	fmt.Println("1. Lihat semua pengguna")
	fmt.Println("2. Lihat semua dana")
	fmt.Println("3. Tambah pengguna baru")
	fmt.Println("4. Tambah dana baru")
	fmt.Println("5. Lihat semua riwayat transaksi")
	fmt.Println("0. Kembali ke menu utama")
	fmt.Println("==============================================")
	fmt.Print("--- Masukkan Pilihan Anda: ")

	var choice int

	fmt.Scan(&choice)

	return choice
}

// Menampilkan menu Pengguna
func User_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%9sSelamat Datang di TelkomFunds %s!    \n", "", storage.Current_user)
	fmt.Println("==============================================")
	fmt.Println("--- Menu: ")
	fmt.Println("1. Lihat riwayat transaksi saya")
	fmt.Println("2. Lihat dana yang tersedia")
	fmt.Println("0. Kembali ke menu utama")
	fmt.Println("==============================================")
	fmt.Print("--- Masukkan Pilihan Anda: ")

	var choice int
	fmt.Scan(&choice)

	return choice
}

// Menampilkan menu login pengguna
func Users_login() int {
	fmt.Println("==============================================")
	fmt.Printf("%10sSelamat Datang di Halaman Masuk    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Masuk atau Daftar: ")
	fmt.Println("1. Masuk")
	fmt.Println("2. Daftar")
	fmt.Println("0. Kembali ke menu utama")
	fmt.Println("==============================================")
	fmt.Print("-- Masukkan Pilihan Anda: ")

	var choice int
	fmt.Scan(&choice)

	return choice
}

// Menampilkan opsi pengguna untuk Admin
func Users_option_admin() int {
	var choice int

	fmt.Println("--- MENU PENGGUNA ADMIN")
	fmt.Println("1. Hapus pengguna")
	fmt.Println("2. Urutkan berdasarkan ID")
	fmt.Println("3. Urutkan berdasarkan total donasi")
	fmt.Println("0. Kembali")
	fmt.Print("--- Masukkan Pilihan Anda: ")

	fmt.Scan(&choice)

	return choice
}

// Menampilkan opsi dana untuk Admin
func Funds_option_admin() int {
	var choice int

	fmt.Println("--- MENU DANA ADMIN")
	fmt.Println("1. Hapus tempat donasi")
	fmt.Println("2. Urutkan berdasarkan ID")
	fmt.Println("3. Urutkan berdasarkan total donasi")
	fmt.Println("0. Kembali")
	fmt.Print("--- Masukkan Pilihan Anda: ")

	fmt.Scan(&choice)

	return choice
}

// Menampilkan opsi dana untuk Pengguna
func Funds_option() int {
	var choice int

	fmt.Println("--- MENU DANA PENGGUNA")
	fmt.Println("1. Donasi ke tempat dana pilihan Anda")
	fmt.Println("2. Urutkan berdasarkan ID")
	fmt.Println("3. Urutkan berdasarkan total donasi")
	fmt.Println("0. Kembali")
	fmt.Print("--- Masukkan Pilihan Anda: ")

	fmt.Scan(&choice)

	return choice
}
