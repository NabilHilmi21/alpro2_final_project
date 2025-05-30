✔️ **FINAL PROJECT ALGORITMA & PEMROGRAMAN 2: _TELKOM FUNDS_**

Proyek ini saya buat sebagai tugas akhir mata kuliah Algoritma & Pemrograman 2, menggunakan bahasa Go. Aplikasinya berbasis terminal (CLI) dan dirancang untuk membantu organisasi atau komunitas kecil mengelola data pengguna dan dana donasi secara sederhana tapi terstruktur.

Terdapat dua peran utama: Admin dan User. Admin bisa melihat, menghapus, menyortir data serta mengecek riwayat transaksi. User bisa login atau registrasi, berdonasi, dan melihat riwayat donasi mereka sendiri.

🗃️ **STRUKTUR FOLDER KODE**

<pre> ``` final_project/ 
  │ 
  ├── main/ # Entry point program 
  │ └── main.go 
  │ 
  ├── tampilan/ # Tampilan menu untuk user & admin 
  │ └── tampilan.go 
  │ 
  ├── auth/ # Login dan registrasi 
  │ └── auth.go 
  │ 
  ├── handlers/ # Alur dan logic pilihan user & admin 
  │ ├── admin.go 
  │ └── user.go 
  │ 
  ├── tools/ # Fitur pendukung seperti tampil, hapus, sorting, dll 
  │ ├── tampil.go 
  │ ├── hapus.go 
  │ ├── sorting.go 
  │ ├── baca.go 
  │ ├── searching.go 
  │ └── donate_funds.go 
  │ 
  ├── storage/ # Penyimpanan data (slicing dan riwayat) 
  │ ├── storage.go 
  │ └── riwayat.go 
  │ 
  ├── structs/ # Struct utama (User, Fund) 
  │ └── structs.go 
  │ 
  └── README.md # Dokumentasi proyek ``` </pre>

👥 **PERAN & AKSES**

🔐 **ADMIN**

1. Login menggunakan kredensial
2. Mengakses data semua user
2. Melihat menghapus dan menambahkan data user/fund
4. Mengurutkan dan mencari data
5. Melihat semua riwayat transaksi user

🙋‍♂️ **USER**

1. Registrasi / Login
2. Berdonasi ke dana yang tersedia
3. Melihat riwayat donasi pribadi

📌 **FITUR UTAMA**
1. ✅ Login & Register
2. 📄 Tampilkan daftar user & fund
3. 🔍 Searching & sorting data
4. 🗑️ Hapus data user/fund
5. 💸 Donasi ke fund
6. 🧾 Riwayat donasi

⚒️ **PERBAIKAN / TAMBAHAN FITUR**
1. Riwayat transaksi pribadi masih belum bisa
