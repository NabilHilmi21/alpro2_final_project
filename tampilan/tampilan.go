package tampilan

import "fmt"


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

func admin_gui()int{
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

func users_option_admin()int{
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

func funds_option_admin()int{
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

func funds_option()int{
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