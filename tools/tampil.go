package tools

import (
	"fmt"

	"github.com/NabilHilmi21/alpro2_final_project/storage"
)

func Tampil_funds() {
	fmt.Println("")
	fmt.Printf("%-8s | %-20s | %-20s\n", "ID Funds", "Nama", "Total Uang")
	fmt.Printf("-----------------------------------------------------------\n")

	for i := 0; i < storage.Funds_count; i++ {
		fmt.Printf("%-8d | %-20s | %-20d\n", storage.Arr_funds[i].Id, storage.Arr_funds[i].Nama, storage.Arr_funds[i].Total)
	}

	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Println("")
}

func Tampil_users() {
	fmt.Println("")
	fmt.Printf("%-8s | %-20s | %-20s\n", "User ID", "Nama", "Total Donasi")
	fmt.Printf("----------------------------------------------\n")
	for i := 0; i < storage.Users_count; i++ {
		fmt.Printf("%-8d | %-20s | %-20d\n", storage.Arr_users[i].Id, storage.Arr_users[i].Nama, storage.Arr_users[i].Total)
	}
	fmt.Printf("----------------------------------------------\n")
	fmt.Println("")
}

func Tampil_riwayat() {
	if storage.Riwayat_count <= 0 {
		fmt.Println("Tidak ada riwayat donasi")
		fmt.Println("")
	} else {
		fmt.Println("")
		fmt.Printf("%-8s | %-20s | %-20s | %-20s | %-20s\n", "User ID", "Nama", "Donasi", "Tempat Donasi", "Tanggal")
		fmt.Printf("-----------------------------------------------------------\n")

		for i := 0; i < storage.Riwayat_count; i++ {
			fmt.Printf("%-8d | %-20s | %-20d | %-20s | %-20s\n", storage.Arr_riwayat[i].Id, storage.Arr_riwayat[i].Nama, storage.Arr_riwayat[i].Nominal, storage.Arr_riwayat[i].Tempat, storage.Arr_riwayat[i].Tanggal)
		}

		fmt.Printf("-----------------------------------------------------------\n")
		fmt.Println("")
	}
}
