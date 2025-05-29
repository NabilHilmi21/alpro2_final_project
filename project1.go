package main

import "fmt"

/*
	NOTE:
	FIX:
	- Kalo udah milih admin/user, ga bisa balik ke menu utama lagi

	ADD:
	- Kalo milih user, bisa enter usernamenya terus di simpen ke array
	- Kalo milih admin, enter username: admin123, password: admin123 (for demonstration purposes only)
*/

// global constants
const NMAX int = 99

// global variables
type user struct {
	id int
	nama string
}

type tabUser [NMAX]user
var users tabUser

// buat tampilin menu utama
func open_gui() int {
	fmt.Println("==============================================")
	fmt.Printf("%10sWelcome to TelkomFunds!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Login as: ")
	fmt.Println("1. Admin")
	fmt.Println("2. User")
	fmt.Println("0. Exit")
	fmt.Print("--- Input Your Choice: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Println("==============================================")
	return choice
}

// buat tampilin menu admin dan user

func open_admin() int {
	fmt.Println("==============================================")
	fmt.Printf("%8sWelcome to TelkomFunds Admin!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Menu: ")
	fmt.Println("1. View all users")
	fmt.Println("2. View all funds")
	fmt.Println("3. Add new user")
	fmt.Println("4. Add new fund")
	fmt.Println("0. Kebali ke menu utama")
	fmt.Print("Your input: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Println("==============================================")
	return choice
}

func open_user() int {
	fmt.Println("==============================================")
	fmt.Printf("%9sWelcome to TelkomFunds User!    \n", "")
	fmt.Println("==============================================")
	fmt.Println("--- Menu: ")
	fmt.Println("1. View my transactions")
	fmt.Println("2. View available funds")
	fmt.Println("3. Make a fund")
	fmt.Println("5. Withdraw funds")
	fmt.Println("0. Kembali ke menu utama")
	fmt.Print("Your input: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Println("==============================================")
	return choice
}

func display_users() {
	if len(users) == 0 {
		fmt.Println("There are no users available.")
	} else {
		fmt.Println("List of users:", users)
	}
}

func user_handler(choice int) {
	switch choice {
	case 1:
		fmt.Println("Viewing my transactions...")
	case 2:
		fmt.Println("Viewing available funds...")
	case 3:
		fmt.Println("Making a fund...")
	case 4:
		fmt.Println("Withdrawing funds...")
	case 0:
		fmt.Println("Kembali ke menu utama...")
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
	}
}

func admin_handler(choice int) {
	switch choice {
	case 1:
		fmt.Println("Viewing all users...")
		display_users()
	case 2:
		fmt.Println("Viewing all funds...")
	case 3:
		fmt.Println("Adding new user...")

		var newUser user
		
		fmt.Print("Enter user ID: ")
		fmt.Scan(&newUser)
		users = append(users, newUser)
		fmt.Println("User added successfully!")
	case 4:
		fmt.Println("Adding new fund...")
	case 0:
		fmt.Println("Kembali ke menu utama...")
		open_gui()
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
	}
}

func main_handler(choice int) {
	var admin_choice, user_choice int

	switch choice {
	case 1:
		admin_choice = open_admin()
		fmt.Println("Pilihan admin:", admin_choice)

		admin_handler(admin_choice)
	case 2:
		user_choice = open_user()
		fmt.Println("Pilihan user:", user_choice)

		user_handler(user_choice)
	case 0:
		fmt.Println("Exiting menu...")
	default:
		fmt.Println("Pilihan tidak valid, silakan coba lagi.")
	}
}

func main() {
	var choice int
	var showMainGUI bool = true

	for {
		if showMainGUI {
			choice = open_gui()
			showMainGUI = false
		}

		if choice == 0 {
			fmt.Println("Terima kasih telah menggunakan TelkomFunds!")
			break
		}

		main_handler(choice)
	}
}


