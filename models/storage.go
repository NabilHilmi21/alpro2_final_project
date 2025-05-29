package models

type Tab_funds [NMAX]funds
type Tab_users [NMAX]users
type Tab_riwayat [NMAX]riwayat

var arr_funds Tab_funds
var arr_users Tab_users
var arr_riwayat Tab_riwayat

var funds_count int = 0
var users_count int = 0
var riwayat_count int = 0

var current_user string
