package auth

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"

	"github.com/NabilHilmi21/alpro2_final_project/tools"
)

func User_reg() {
	var username, password string
	var show_reg bool = true

	for show_reg {
		fmt.Println("")
		fmt.Print("--- Masukkan nama pengguna: ")
		fmt.Scan(&username)

		if tools.Check_user(storage.Arr_users, storage.Users_count, username) != -1 {
			fmt.Println("Nama pengguna sudah ada, silakan coba lagi.")
			fmt.Println("")
			continue
		}

		fmt.Print("--- Masukkan kata sandi: ")
		fmt.Scan(&password)
		fmt.Println("")

		if len(password) < 3 {
			fmt.Println("Kata sandi harus minimal 3 karakter, silakan coba lagi.")
			fmt.Println("")
			continue
		}

		storage.Arr_users[storage.Users_count].Id = storage.Users_count
		storage.Arr_users[storage.Users_count].Nama = username
		storage.Arr_users[storage.Users_count].Password = password

		storage.Users_count++

		fmt.Println("Registrasi berhasil, Anda sekarang bisa masuk.")
		fmt.Println("")

		show_reg = false
	}
}

func User_auth() bool {
	var logged_in bool = false
	var username, password string
	var index_user int
	var i int = 1

	for i <= 5 && !logged_in {
		fmt.Println("")
		fmt.Print("Masukkan nama pengguna: ")
		fmt.Scan(&username)

		index_user = tools.Check_user(storage.Arr_users, storage.Users_count, username)

		fmt.Print("Masukkan kata sandi: ")
		fmt.Scan(&password)
		fmt.Println("")

		if index_user != -1 && storage.Arr_users[index_user].Password == password {
			fmt.Println("Anda telah berhasil masuk sebagai", username, ".")
			fmt.Println("")

			storage.Current_user = username

			logged_in = true
			return logged_in
		}

		fmt.Println("Nama pengguna atau kata sandi salah. Sisa percobaan:", 5-i, "kali.")
		fmt.Println("")
		i++
	}

	return logged_in
}

func Admin_auth() bool {
	var logged_in bool = false
	var username, password string
	var i int = 1

	for i <= 5 && !logged_in {
		fmt.Print("Masukkan nama pengguna: ")
		fmt.Scan(&username)

		fmt.Print("Masukkan kata sandi: ")
		fmt.Scan(&password)

		fmt.Println("")

		if username == "admin123" && password == "admin123" {
			fmt.Println("Admin berhasil masuk.")
			fmt.Println("")
			logged_in = true
			return logged_in
		}

		fmt.Println("Kredensial salah. Anda memiliki", 5-i, "percobaan tersisa.")
		fmt.Println("")
		i++
	}

	return logged_in
}
