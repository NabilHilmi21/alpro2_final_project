package auth

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"
)

func User_auth() bool {
	var logged_in bool = false
	var username string
	var i int = 1

	for i <= 5 && logged_in == false {
		fmt.Print("Enter username: ")
		fmt.Scan(&username)

		if search.Check_user(storage.Arr_users, storage.Users_count, username) != -1 {
			fmt.Println("Kamu telah login sebagai", username)
			fmt.Println("")

			storage.Current_user = username

			logged_in = true
			return logged_in
		}

		fmt.Println("User not found, you have", 5-i, "tries left")
		fmt.Println("")
		i++
	}

	return logged_in
}

func Admin_auth() bool {
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
