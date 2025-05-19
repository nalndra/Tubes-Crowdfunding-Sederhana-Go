package donation

import (
	"fmt"

	project "github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/dataProyek"
)

func Donasi() {
	var nomorProyek, jumlahDonasi int

	fmt.Print("Pilih nomor proyek yang ingin didonasi: ")
	fmt.Scanln(&nomorProyek)

	if nomorProyek < 1 || nomorProyek > len(project.DaftarProyek) {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	fmt.Print("Masukkan jumlah donasi: Rp")
	fmt.Scanln(&jumlahDonasi)

	if jumlahDonasi <= 0 {
		fmt.Println("Jumlah donasi harus lebih dari 0")
		return
	}

	proyek := &project.DaftarProyek[nomorProyek-1]
	proyek.Dana += float64(jumlahDonasi)
	proyek.Donatur++

	fmt.Printf("\nTerima kasih telah mendonasikan Rp%d untuk proyek %s\n", jumlahDonasi, proyek.Nama)

	progress := (proyek.Dana / proyek.Target) * 100
	fmt.Printf("Progress terkini: %.0f%%\n", progress)

	if proyek.Dana >= proyek.Target {
		fmt.Println("Selamat! Proyek ini telah mencapai target pendanaan!")
	}

	fmt.Println("-------------------------------------------------------------------------------")
}
