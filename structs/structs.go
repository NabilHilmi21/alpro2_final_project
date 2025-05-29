package structs

const NMAX int = 99

type Users struct {
	Id    int
	Total int
	Nama  string
}

type Funds struct {
	Id      int
	Nama    string
	Nominal int
	Total   int
}

type Riwayat struct {
	Id      int
	Nama    string
	Nominal int
	Tanggal string
	Tempat  string
}
