package models

const NMAX int = 99

type users struct {
	Id    int
	Total int
	Nama  string
}

type funds struct {
	Id      int
	Nama    string
	Nominal int
	Total   int
}

type riwayat struct {
	Id      int
	Nama    string
	Nominal int
	Tanggal string
	Tempat  string
}
