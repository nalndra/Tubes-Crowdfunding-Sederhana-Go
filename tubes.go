package main

import "fmt"

type Proyek struct {
	Nama     string
	Kategori string
	Target   float64
	Dana     float64
	Donatur  int
	Pemilik  string
}

const NMAX int = 100

type DaftarProyek [NMAX]Proyek

func main() {
	var daftarProyek DaftarProyek
	var n int = 3

	daftarProyek[0] = Proyek{"Robot_Edukasi", "Pendidikan", 5000000, 3500000, 45, "user1"}
	daftarProyek[1] = Proyek{"Aplikasi_Kesehatan", "Kesehatan", 10000000, 7500000, 32, "user2"}
	daftarProyek[2] = Proyek{"Buku_Anak", "Pendidikan", 2000000, 2500000, 28, "user1"}

	var currentUser string
	login(&currentUser)

	var input string
	for {
		tampilkanMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			tampilkanSemuaProyek(&daftarProyek, n, func() {
				donasi(&daftarProyek, n)
			})
		case "2":
			tambahProyekBaru(&daftarProyek, &n, currentUser, "", "", 0, 0, 0)
		case "3":
			cariProyek(daftarProyek, n)
		case "4":
			urutkanProyek(&daftarProyek, n)
		case "5":
			kelolaProyekSaya(&daftarProyek, &n, currentUser)
		case "6":
			proyekBerhasilPendanaan(daftarProyek, n)
		case "7":
			fmt.Println("Sampai Jumpa Lagi...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func login(currentUser *string) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                    SELAMAT DATANG DI APLIKASI CROWDFUNDING                    ")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(currentUser)
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

func tampilkanSemuaProyek(daftarProyek *DaftarProyek, n int, donasiFunc func()) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                              DAFTAR SEMUA PROYEK                              ")
	fmt.Println("-------------------------------------------------------------------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s (%s)\n", i+1, daftarProyek[i].Nama, daftarProyek[i].Kategori)
		fmt.Printf("   Target: Rp%.0f\n", daftarProyek[i].Target)
		fmt.Printf("   Dana terkumpul: Rp%.0f\n", daftarProyek[i].Dana)
		fmt.Printf("   Jumlah Donatur: %d donatur\n", daftarProyek[i].Donatur)
		fmt.Printf("   Progress: %.0f%%\n", (daftarProyek[i].Dana/daftarProyek[i].Target)*100)
		fmt.Println()
	}
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Donasi")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Print("Pilih menu: ")

	var pilihan string
	fmt.Scanln(&pilihan)

	if pilihan == "1" {
		donasi(daftarProyek, n)
	}
}

func donasi(daftarProyek *DaftarProyek, n int) {
	var nomorProyek, jumlahDonasi int

	fmt.Print("Pilih nomor proyek yang ingin didonasi: ")
	fmt.Scanln(&nomorProyek)

	if nomorProyek < 1 || nomorProyek > n {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	fmt.Print("Masukkan jumlah donasi: Rp")
	fmt.Scanln(&jumlahDonasi)

	if jumlahDonasi <= 0 {
		fmt.Println("Jumlah donasi harus lebih dari 0")
		return
	}

	proyek := &daftarProyek[nomorProyek-1]
	proyek.Dana += float64(jumlahDonasi)
	proyek.Donatur++

	fmt.Printf("\nTerima kasih telah mendonasikan Rp%d untuk proyek %s\n", jumlahDonasi, proyek.Nama)

	progress := (proyek.Dana / proyek.Target) * 100
	fmt.Printf("Progress terkini: %.0f%%\n", progress)

	if proyek.Dana >= proyek.Target && proyek.Target > 0 {
		fmt.Println("Selamat! Proyek ini telah mencapai target pendanaan!")
	}
	fmt.Println("-------------------------------------------------------------------------------")
}

func kelolaProyekSaya(daftarProyek *DaftarProyek, n *int, currentUser string) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                              KELOLA PROYEK SAYA                               ")
	fmt.Println("-------------------------------------------------------------------------------")

	var proyekSaya DaftarProyek
	var jumlahProyekSaya int = 0

	for i := 0; i < *n; i++ {
		if daftarProyek[i].Pemilik == currentUser {
			proyekSaya[jumlahProyekSaya] = daftarProyek[i]
			jumlahProyekSaya++
		}
	}

	if jumlahProyekSaya == 0 {
		fmt.Println("Anda belum memiliki proyek.")
		return
	}

	for i := 0; i < jumlahProyekSaya; i++ {
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

		if pilihan == "1" {
			editProyekSaya(daftarProyek, n, proyekSaya, jumlahProyekSaya, currentUser)
		} else if pilihan == "2" {
			hapusProyekSaya(daftarProyek, n, proyekSaya, jumlahProyekSaya, currentUser)
		} else if pilihan == "3" {
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func editProyekSaya(daftarProyek *DaftarProyek, n *int, proyekSaya DaftarProyek, jumlahProyekSaya int, currentUser string) {
	var nomor int
	fmt.Print("Pilih nomor proyek yang akan diedit: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > jumlahProyekSaya {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	indexDiDaftarUtama := -1
	for i := 0; i < *n; i++ {
		if daftarProyek[i].Nama == proyekSaya[nomor-1].Nama && daftarProyek[i].Pemilik == currentUser {
			indexDiDaftarUtama = i
		}
	}

	if indexDiDaftarUtama == -1 {
		fmt.Println("Proyek tidak ditemukan!")
		return
	}

	fmt.Printf("\nEdit Proyek: %s\n", daftarProyek[indexDiDaftarUtama].Nama)

	var namaBaru, kategoriBaru string
	var targetBaru, danaBaru float64
	var donaturBaru int

	fmt.Print("Nama Baru (kosongkan jika tidak diubah): ")
	fmt.Scanln(&namaBaru)
	if namaBaru != "" {
		daftarProyek[indexDiDaftarUtama].Nama = namaBaru
	}

	fmt.Print("Kategori Baru (kosongkan jika tidak diubah): ")
	fmt.Scanln(&kategoriBaru)
	if kategoriBaru != "" {
		daftarProyek[indexDiDaftarUtama].Kategori = kategoriBaru
	}

	fmt.Print("Target Dana Baru (0 jika tidak diubah): ")
	fmt.Scanln(&targetBaru)
	if targetBaru != 0 {
		daftarProyek[indexDiDaftarUtama].Target = targetBaru
	}

	fmt.Print("Dana Terkumpul Baru (0 jika tidak diubah): ")
	fmt.Scanln(&danaBaru)
	if danaBaru != 0 {
		daftarProyek[indexDiDaftarUtama].Dana = danaBaru
	}

	fmt.Print("Jumlah Donatur Baru (0 jika tidak diubah): ")
	fmt.Scanln(&donaturBaru)
	if donaturBaru != 0 {
		daftarProyek[indexDiDaftarUtama].Donatur = donaturBaru
	}

	fmt.Println("Proyek berhasil diupdate!")
}

func hapusProyekSaya(daftarProyek *DaftarProyek, n *int, proyekSaya DaftarProyek, jumlahProyekSaya int, currentUser string) {
	var nomor int
	fmt.Print("Pilih nomor proyek yang akan dihapus: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > jumlahProyekSaya {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	proyekDipilih := proyekSaya[nomor-1]
	indexDiDaftarUtama := -1
	for i := 0; i < *n; i++ {
		if daftarProyek[i].Nama == proyekDipilih.Nama && daftarProyek[i].Pemilik == currentUser {
			indexDiDaftarUtama = i
		}
	}

	if indexDiDaftarUtama == -1 {
		fmt.Println("Proyek tidak ditemukan!")
		return
	}

	var konfirmasi string
	fmt.Printf("Anda yakin ingin menghapus proyek %s? (y/n): ", daftarProyek[indexDiDaftarUtama].Nama)
	fmt.Scanln(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		for i := indexDiDaftarUtama; i < *n-1; i++ {
			daftarProyek[i] = daftarProyek[i+1]
		}
		*n--
		fmt.Println("Proyek berhasil dihapus!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}
}

func tambahProyekBaru(daftarProyek *DaftarProyek, n *int, currentUser string, nama string, kategori string, target float64, dana float64, donatur int) {
	if *n >= NMAX {
		fmt.Println("Kapasitas proyek sudah penuh!")
		return
	}

	if nama == "" {
		fmt.Println("\nTambah Proyek Baru")
		fmt.Print("Nama Proyek: ")
		fmt.Scanln(&nama)
		fmt.Print("Kategori: ")
		fmt.Scanln(&kategori)
		fmt.Print("Target Dana: ")
		fmt.Scanln(&target)
		fmt.Print("Dana Awal: ")
		fmt.Scanln(&dana)
		fmt.Print("Jumlah Donatur Awal: ")
		fmt.Scanln(&donatur)
	}

	var p Proyek
	p.Nama = nama
	p.Kategori = kategori
	p.Target = target
	p.Dana = dana
	p.Donatur = donatur
	p.Pemilik = currentUser

	daftarProyek[*n] = p
	*n++
	fmt.Println("Proyek berhasil ditambahkan!")
}

func cariProyek(daftarProyek DaftarProyek, n int) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                                  CARI PROYEK                                  ")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Nama (Sequential Search)")
	fmt.Println("2. Kategori (Binary Search)")
	fmt.Println("-------------------------------------------------------------------------------")
	var pilihan string
	fmt.Print("Pilih metode: ")
	fmt.Scanln(&pilihan)

	var keyword string
	switch pilihan {
		case "1":
			fmt.Print("Masukkan nama proyek: ")
			fmt.Scanln(&keyword)
			sequentialSearch(daftarProyek, n, keyword)
		case "2":
			fmt.Print("Masukkan Kategori proyek: ")
			fmt.Scanln(&keyword)
			proyekTerurut := daftarProyek
        	sortByCategory(&proyekTerurut, n)
			binarySearch(proyekTerurut, n, keyword)
		default:
			fmt.Println("Pilihan tidak valid!")
	}
}

func sequentialSearch(daftarProyek DaftarProyek, n int, keyword string) {
	fmt.Println("\nHasil Pencarian Berdasarkan Nama:")
	found := false

	for i := 0; i < n; i++ {
		if daftarProyek[i].Nama == keyword {
			tampilkanDetailProyek(daftarProyek[i])
			found = true
		}
	}

	if !found {
		fmt.Println("Proyek tidak ditemukan!")
	}
}

func sortByCategory(daftarProyek *DaftarProyek, n int) {
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if daftarProyek[j].Kategori > daftarProyek[j+1].Kategori {
                daftarProyek[j], daftarProyek[j+1] = daftarProyek[j+1], daftarProyek[j]
            }
        }
    }
}

func binarySearch(daftarProyek DaftarProyek, n int, keyword string) {
	fmt.Println("\nHasil Pencarian Berdasarkan Kategori:")
	left := 0
	right := n - 1
	found := false
	var indexDitemukan int = -1

	for left <= right {
		mid := (left + right) / 2
		if daftarProyek[mid].Kategori == keyword {
			indexDitemukan = mid
			found = true
			break
		} else if daftarProyek[mid].Kategori < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found {
        i := indexDitemukan
        for i >= 0 && daftarProyek[i].Kategori == keyword {
            tampilkanDetailProyek(daftarProyek[i])
            i--
        }

        i = indexDitemukan + 1
        for i < n && daftarProyek[i].Kategori == keyword {
            tampilkanDetailProyek(daftarProyek[i])
            i++
        }
    } else {
        fmt.Println("Proyek dengan kategori", keyword, "tidak ditemukan!")
    }

}

func urutkanProyek(daftarProyek *DaftarProyek, n int) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                                 URUTKAN PROYEK                                ")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Total Dana Terkumpul (Selection Sort)")
	fmt.Println("2. Jumlah Donatur (Insertion Sort)")
	fmt.Println("-------------------------------------------------------------------------------")
	var pilihan string
	fmt.Print("Pilih metode: ")
	fmt.Scanln(&pilihan)

	if pilihan == "1" {
		selectionSort(daftarProyek, n)
		fmt.Println("Proyek berhasil diurutkan berdasarkan Total Dana Terkumpul!")
	} else if pilihan == "2" {
		insertionSort(daftarProyek, n)
		fmt.Println("Proyek berhasil diurutkan berdasarkan Jumlah Donatur!")
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

func selectionSort(daftarProyek *DaftarProyek, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarProyek[j].Dana < daftarProyek[minIdx].Dana {
				minIdx = j
			}
		}
		daftarProyek[i], daftarProyek[minIdx] = daftarProyek[minIdx], daftarProyek[i]
	}
}

func insertionSort(daftarProyek *DaftarProyek, n int) {
	for i := 1; i < n; i++ {
		key := daftarProyek[i]
		j := i - 1
		for j >= 0 && daftarProyek[j].Donatur > key.Donatur {
			daftarProyek[j+1] = daftarProyek[j]
			j--
		}
		daftarProyek[j+1] = key
	}
}

func proyekBerhasilPendanaan(daftarProyek DaftarProyek, n int) {
	fmt.Println("\n-------------------------------------------------------------------------------")
	fmt.Println("                            PROYEK BERHASIL DIDANAI                            ")
	fmt.Println("-------------------------------------------------------------------------------")
	found := false
	for i := 0; i < n; i++ {
		if daftarProyek[i].Dana >= daftarProyek[i].Target {
			tampilkanDetailProyek(daftarProyek[i])
			found = true
		}
	}
	if !found {
		fmt.Println("Belum ada proyek yang mencapai target")
	}
}

func tampilkanDetailProyek(p Proyek) {
	fmt.Printf("\n%s (%s)\n", p.Nama, p.Kategori)
	fmt.Printf("   Target: Rp%.0f\n", p.Target)
	fmt.Printf("   Dana Terkumpul: Rp%.0f\n", p.Dana)
	fmt.Printf("   Jumlah Donatur: %d donatur\n", p.Donatur)
	progress := 0.0
	if p.Target > 0 {
		progress = (p.Dana / p.Target) * 100
	}
	fmt.Printf("   Progress: %.0f%%\n", progress)
}