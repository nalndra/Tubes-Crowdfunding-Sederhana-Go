package project

import (
	"fmt"
)

type Proyek struct {
	Nama     string
	Kategori string
	Target   float64
	Dana     float64
	Donatur  int
	Pemilik  string // Tambahan field untuk menyimpan pemilik proyek
}

var DaftarProyek []Proyek

func TampilkanSemuaProyek(donasiFunc func()) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                              DAFTAR SEMUA PROYEK                              ")
	fmt.Println("-------------------------------------------------------------------------------")
	for i := 0; i < len(DaftarProyek); i++ {
		fmt.Printf("%d. %s (%s)\n", i+1, DaftarProyek[i].Nama, DaftarProyek[i].Kategori)
		fmt.Printf("   Target: Rp%.0f\n", DaftarProyek[i].Target)
		fmt.Printf("   Dana terkumpul: Rp%.0f\n", DaftarProyek[i].Dana)
		fmt.Printf("   Jumlah Donatur: %d donatur\n", DaftarProyek[i].Donatur)
		fmt.Printf("   Progress: %.0f%%\n", (DaftarProyek[i].Dana/DaftarProyek[i].Target)*100)
		fmt.Println()
	}
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Donasi")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Print("Pilih menu: ")
	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		donasiFunc()
	case "2":
		return
	default:
		fmt.Println("Pilihan Tidak Valid!")
	}
}

func KelolaProyekSaya(currentUser string, editFunc func(string, []Proyek), hapusFunc func(string, []Proyek)) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                              KELOLA PROYEK SAYA                               ")
	fmt.Println("-------------------------------------------------------------------------------")

	var proyekSaya []Proyek
	for i := 0; i < len(DaftarProyek); i++ {
		if DaftarProyek[i].Pemilik == currentUser {
			proyekSaya = append(proyekSaya, DaftarProyek[i])
		}
	}

	if len(proyekSaya) == 0 {
		fmt.Println("Anda belum memiliki proyek.")
		return
	}

	for i := 0; i < len(proyekSaya); i++ {
		fmt.Printf("%d. %s (%s)\n", i+1, proyekSaya[i].Nama, proyekSaya[i].Kategori)
		fmt.Printf("   Target: Rp%.0f\n", proyekSaya[i].Target)
		fmt.Printf("   Dana terkumpul: Rp%.0f\n", proyekSaya[i].Dana)
		fmt.Printf("   Jumlah Donatur: %d donatur\n", proyekSaya[i].Donatur)
		fmt.Println()
	}
	fmt.Println("-------------------------------------------------------------------------------")

	var pilihan string
	for {
		fmt.Println("\n1. Edit Proyek")
		fmt.Println("2. Hapus Proyek")
		fmt.Println("3. Kembali")
		fmt.Println("-------------------------------------------------------------------------------")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			editFunc(currentUser, proyekSaya)
			break
		case "2":
			HapusProyekSaya(currentUser, proyekSaya)
			break
		case "3":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		break
	}
}

func EditProyekSaya(currentUser string, proyekSaya []Proyek) {
	var nomor int
	fmt.Print("Pilih nomor proyek yang akan diedit: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(proyekSaya) {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	var indexDiDaftarUtama int = -1
	for i := 0; i < len(DaftarProyek); i++ {
		if DaftarProyek[i].Nama == proyekSaya[nomor-1].Nama && DaftarProyek[i].Pemilik == currentUser {
			indexDiDaftarUtama = i
			break
		}
	}

	if indexDiDaftarUtama == -1 {
		fmt.Println("Proyek tidak ditemukan!")
		return
	}

	fmt.Printf("\nEdit Proyek: %s\n", DaftarProyek[indexDiDaftarUtama].Nama)

	var namaBaru, kategoriBaru string
	var targetBaru, danaBaru float64
	var donaturBaru int

	fmt.Print("Nama Baru (kosongkan jika tidak diubah): ")
	fmt.Scanln(&namaBaru)
	if namaBaru != "" {
		DaftarProyek[indexDiDaftarUtama].Nama = namaBaru
	}

	fmt.Print("Kategori Baru (kosongkan jika tidak diubah): ")
	fmt.Scanln(&kategoriBaru)
	if kategoriBaru != "" {
		DaftarProyek[indexDiDaftarUtama].Kategori = kategoriBaru
	}

	fmt.Print("Target Dana Baru (0 jika tidak diubah): ")
	fmt.Scanln(&targetBaru)
	if targetBaru != 0 {
		DaftarProyek[indexDiDaftarUtama].Target = targetBaru
	}

	fmt.Print("Dana Terkumpul Baru (0 jika tidak diubah): ")
	fmt.Scanln(&danaBaru)
	if danaBaru != 0 {
		DaftarProyek[indexDiDaftarUtama].Dana = danaBaru
	}

	fmt.Print("Jumlah Donatur Baru (0 jika tidak diubah): ")
	fmt.Scanln(&donaturBaru)
	if donaturBaru != 0 {
		DaftarProyek[indexDiDaftarUtama].Donatur = donaturBaru
	}

	fmt.Println("Proyek berhasil diupdate!")
}

func HapusProyekSaya(currentUser string, proyekSaya []Proyek) {
	var nomor int
	fmt.Print("Pilih nomor proyek yang akan dihapus: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(proyekSaya) {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	proyekDipilih := proyekSaya[nomor-1]
	var indexDiDaftarUtama int = -1
	for i := 0; i < len(DaftarProyek); i++ {
		if DaftarProyek[i].Nama == proyekDipilih.Nama && DaftarProyek[i].Pemilik == currentUser {
			indexDiDaftarUtama = i
			break
		}
	}

	if indexDiDaftarUtama == -1 {
		fmt.Println("Proyek tidak ditemukan!")
		return
	}

	var konfirmasi string
	fmt.Printf("Anda yakin ingin menghapus proyek %s? (y/n): ", DaftarProyek[indexDiDaftarUtama].Nama)
	fmt.Scanln(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		DaftarProyek = append(DaftarProyek[:indexDiDaftarUtama], DaftarProyek[indexDiDaftarUtama+1:]...)
		fmt.Println("Proyek berhasil dihapus!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

func TambahProyekBaru(currentUser string) {
	var p Proyek
	fmt.Println("\nTambah Proyek Baru")
	fmt.Print("Nama Proyek: ")
	fmt.Scanln(&p.Nama)
	fmt.Print("Kategori: ")
	fmt.Scanln(&p.Kategori)
	fmt.Print("Target Dana: ")
	fmt.Scanln(&p.Target)
	fmt.Print("Dana Awal: ")
	fmt.Scanln(&p.Dana)
	fmt.Print("Jumlah Donatur Awal: ")
	fmt.Scanln(&p.Donatur)

	p.Pemilik = currentUser

	DaftarProyek = append(DaftarProyek, p)
	fmt.Println("Proyek berhasil ditambahkan!")
}

func ProyekBerhasilPendanaan() {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                            PROYEK BERHASIL DIDANAI                            ")
	fmt.Println("-------------------------------------------------------------------------------")
	found := false
	for i := 0; i < len(DaftarProyek); i++ {
		if DaftarProyek[i].Dana >= DaftarProyek[i].Target {
			TampilkanDetailProyek(DaftarProyek[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Belum ada proyek yang mencapai target")
	}
}

func TampilkanDetailProyek(p Proyek) {
	fmt.Printf("\n%s (%s)\n", p.Nama, p.Kategori)
	fmt.Printf("   Target: Rp%.0f\n", p.Target)
	fmt.Printf("   Dana Terkumpul: Rp%.0f\n", p.Dana)
	fmt.Printf("   Jumlah Donatur: %d donatur\n", p.Donatur)
	fmt.Printf("   Progress: %.0f%%\n", (p.Dana/p.Target)*100)
}
