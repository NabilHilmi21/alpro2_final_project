package storage

import "github.com/NabilHilmi21/alpro2_final_project/structs"

type Tab_funds [structs.NMAX]structs.Funds
type Tab_users [structs.NMAX]structs.Users
type Tab_riwayat [structs.NMAX]structs.Riwayat

var Arr_funds Tab_funds
var Arr_users Tab_users
var Arr_riwayat Tab_riwayat

var Funds_count int = 0
var Users_count int = 0
var Riwayat_count int = 0

var Current_user string
