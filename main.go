package main

import (
	"fmt"

	project "github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/dataProyek"
	donation "github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/donasi"
	"github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/search"
	"github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/sort"
)

var currentUser string

func main() {
	// Inisialisasi data contoh
	project.DaftarProyek = []project.Proyek{
		{Nama: "Robot_Edukasi", Kategori: "Pendidikan", Target: 5000000, Dana: 3500000, Donatur: 45, Pemilik: "user1"},
		{Nama: "Aplikasi_Kesehatan", Kategori: "Kesehatan", Target: 10000000, Dana: 7500000, Donatur: 32, Pemilik: "user2"},
		{Nama: "Buku_Anak", Kategori: "Pendidikan", Target: 2000000, Dana: 2500000, Donatur: 28, Pemilik: "user1"},
	}

	// Simulasi login
	login()
	fmt.Scanln(&currentUser)

	var input string
	for {
		tampilkanMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			project.TampilkanSemuaProyek(donation.Donasi)
		case "2":
			project.TambahProyekBaru(currentUser)
		case "3":
			search.CariProyek()
		case "4":
			sort.UrutkanProyek()
		case "5":
			project.KelolaProyekSaya(currentUser, project.EditProyekSaya, project.HapusProyekSaya)
		case "6":
			project.ProyekBerhasilPendanaan()
		case "7":
			fmt.Println("Sampai Jumpa Lagi...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func login() {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                    SELAMAT DATANG DI APLIKASI CROWDFUNDING                    ")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Print("Masukkan Username: ")
}

func tampilkanMenu() {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                                   MENU UTAMA                                  ")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Tampilkan Semua Proyek")
	fmt.Println("2. Tambah Proyek Baru")
	fmt.Println("3. Cari Proyek")
	fmt.Println("4. Urutkan Proyek")
	fmt.Println("5. Kelola Proyek Saya")
	fmt.Println("6. Proyek Berhasil didanai")
	fmt.Println("7. Keluar dari Aplikasi")
	fmt.Println("-------------------------------------------------------------------------------")
}
