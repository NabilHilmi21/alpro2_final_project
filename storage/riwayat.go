package storage

func Riwayat_funds(idx int, user string, donasi int, tempat string, tanggal string) {
	Arr_riwayat[Riwayat_count].Id = Arr_users[idx].Id
	Arr_riwayat[Riwayat_count].Nama = user
	Arr_riwayat[Riwayat_count].Nominal = donasi
	Arr_riwayat[Riwayat_count].Tempat = tempat
	Arr_riwayat[Riwayat_count].Tanggal = tanggal

	Riwayat_count++
}
