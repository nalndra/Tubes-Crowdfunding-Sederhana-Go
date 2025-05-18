package main

import "fmt"

type Proyek struct {
    Nama     string
    Kategori string
    Target   float64
    Dana     float64
    Donatur  int
    Pemilik  string // Tambahan field untuk menyimpan pemilik proyek
}

var daftarProyek []Proyek
var currentUser string // Variabel untuk menyimpan user yang sedang login

func main() {
    // Inisialisasi data contoh
    daftarProyek = []Proyek{
        {"Robot_Edukasi", "Pendidikan", 5000000, 3500000, 45, "user1"},
        {"Aplikasi_Kesehatan", "Kesehatan", 10000000, 7500000, 32, "user2"},
        {"Buku_Anak", "Pendidikan", 2000000, 2500000, 28, "user1"},
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
            tampilkanSemuaProyek()
        case "2":
            tambahProyekBaru()
        case "3":
            cariProyek()
        case "4":
            urutkanProyek()
        case "5":
            kelolaProyekSaya()
        case "6":
            proyekBerhasilPendanaan()
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

func tampilkanSemuaProyek() {
	fmt.Println("\n-------------------------------------------------------------------------------")
    fmt.Println("                              DAFTAR SEMUA PROYEK                              ")
    fmt.Println("-------------------------------------------------------------------------------")
    var i int
    var pilihan string
	for i = 0; i < len(daftarProyek); i++ {
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
    fmt.Scanln(&pilihan)

    switch pilihan {
    case "1":
        donasi()
    case "2":
        return
    default:
        fmt.Println("Pilihan Tidak Valid!")
    }
}

func donasi() {
    var nomorProyek, jumlahDonasi int
    
    fmt.Print("Pilih nomor proyek yang ingin didonasi: ")
    fmt.Scanln(&nomorProyek)
    
    if nomorProyek < 1 || nomorProyek > len(daftarProyek) {
        fmt.Println("Nomor proyek tidak valid!")
        return
    }
    
    fmt.Print("Masukkan jumlah donasi: Rp")
    fmt.Scanln(&jumlahDonasi)
    
    if jumlahDonasi <= 0 {
        fmt.Println("Jumlah donasi harus lebih dari 0")
        return
    }
    
    // Update data proyek
    proyek := &daftarProyek[nomorProyek-1]
    proyek.Dana += float64(jumlahDonasi)
    proyek.Donatur++
    
    fmt.Printf("\nTerima kasih telah mendonasikan Rp%d untuk proyek %s\n", jumlahDonasi, proyek.Nama)
    
    // Tampilkan progress terbaru
    progress := (proyek.Dana / proyek.Target) * 100
    fmt.Printf("Progress terkini: %.0f%%\n", progress)
    
    // Cek jika proyek sudah mencapai target
    if proyek.Dana >= proyek.Target {
        fmt.Println("Selamat! Proyek ini telah mencapai target pendanaan!")
    }
    
    fmt.Println("-------------------------------------------------------------------------------")
}

func kelolaProyekSaya() {
    fmt.Println("\n-------------------------------------------------------------------------------")
    fmt.Println("                              KELOLA PROYEK SAYA                               ")
    fmt.Println("-------------------------------------------------------------------------------")
    
    // Tampilkan hanya proyek milik currentUser
    var proyekSaya []Proyek
    for i := 0; i < len(daftarProyek); i++ {
        if daftarProyek[i].Pemilik == currentUser {
            proyekSaya = append(proyekSaya, daftarProyek[i])
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
            editProyekSaya(proyekSaya)
            break
        case "2":
            hapusProyekSaya(proyekSaya)
            break
        case "3":
            return
        default:
            fmt.Println("Pilihan tidak valid!")
            continue
        }
        
        // Keluar dari loop setelah selesai edit/hapus
        break
    }
}

func editProyekSaya(proyekSaya []Proyek) {
    var nomor int
    fmt.Print("Pilih nomor proyek yang akan diedit: ")
    fmt.Scanln(&nomor)
    
    if nomor < 1 || nomor > len(proyekSaya) {
        fmt.Println("Nomor proyek tidak valid!")
        return
    }

    var indexDiDaftarUtama int = -1
    for i := 0; i < len(daftarProyek); i++ {
        if daftarProyek[i].Nama == proyekSaya[nomor-1].Nama && daftarProyek[i].Pemilik == currentUser {
            indexDiDaftarUtama = i
            break
        }
    }

    if indexDiDaftarUtama == -1 {
        fmt.Println("Proyek tidak ditemukan!")
        return
    }

    // Proses edit
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

func hapusProyekSaya(proyekSaya []Proyek) {
    var nomor int
    fmt.Print("Pilih nomor proyek yang akan dihapus: ")
    fmt.Scanln(&nomor)
    
    if nomor < 1 || nomor > len(proyekSaya) {
        fmt.Println("Nomor proyek tidak valid!")
        return
    }

    // Cari proyek yang dipilih di daftar utama
    proyekDipilih := proyekSaya[nomor-1]
    var indexDiDaftarUtama int = -1
    for i := 0; i < len(daftarProyek); i++ {
        if daftarProyek[i].Nama == proyekDipilih.Nama && daftarProyek[i].Pemilik == currentUser {
            indexDiDaftarUtama = i
            break
        }
    }

    if indexDiDaftarUtama == -1 {
        fmt.Println("Proyek tidak ditemukan!")
        return
    }

    // Konfirmasi penghapusan
    var konfirmasi string
    fmt.Printf("Anda yakin ingin menghapus proyek %s? (y/n): ", daftarProyek[indexDiDaftarUtama].Nama)
    fmt.Scanln(&konfirmasi)
    
    if konfirmasi == "y" || konfirmasi == "Y" {
        daftarProyek = append(daftarProyek[:indexDiDaftarUtama], daftarProyek[indexDiDaftarUtama+1:]...)
        fmt.Println("Proyek berhasil dihapus!")
    } else {
        fmt.Println("Penghapusan dibatalkan.")
    }
}

// Modifikasi fungsi tambahProyekBaru untuk menyimpan informasi pemilik
func tambahProyekBaru() {
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
    
    p.Pemilik = currentUser // Set pemilik proyek sebagai user yang sedang login
    
    daftarProyek = append(daftarProyek, p)
    fmt.Println("Proyek berhasil ditambahkan!")
}

func cariProyek() {
	fmt.Println("\n-------------------------------------------------------------------------------")
    fmt.Println("                                  CARI PROYEK                                  ")
    fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Sequential Search (Nama)")
	fmt.Println("2. Binary Search (Nama)")
    fmt.Println("-------------------------------------------------------------------------------")
	var pilihan string
	fmt.Print("Pilih metode: ")
	fmt.Scanln(&pilihan)

	var keyword string
	fmt.Print("Masukkan nama proyek: ")
	fmt.Scanln(&keyword)

	switch pilihan {
	case "1":
		sequentialSearch(keyword)
	case "2":
		binarySearch(keyword)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func sequentialSearch(keyword string) {
    fmt.Println("\nHasil Pencarian Sequential:")
    found := false
    
    for i := 0; i < len(daftarProyek); i++ {
        if daftarProyek[i].Nama == keyword {
            tampilkanDetailProyek(daftarProyek[i])
            found = true
        }
    }
    
    if !found {
        fmt.Println("Proyek tidak ditemukan!")
    }
}

func binarySearch(keyword string) {
    fmt.Println("\nHasil Pencarian Binary:")
    
    low := 0
    high := len(daftarProyek) - 1
    found := false
    
    for low <= high {
        mid := low + (high-low)/2
        
        if daftarProyek[mid].Nama == keyword {
            tampilkanDetailProyek(daftarProyek[mid])
            found = true
            break
        } else if daftarProyek[mid].Nama < keyword {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    if !found {
        fmt.Println("Proyek tidak ditemukan!")
    }
}

func urutkanProyek() {
	fmt.Println("\n-------------------------------------------------------------------------------")
    fmt.Println("                                 URUTKAN PROYEK                                ")
    fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("1. Selection Sort (Dana)")
	fmt.Println("2. Insertion Sort (Donatur)")
    fmt.Println("-------------------------------------------------------------------------------")
	var pilihan string
	fmt.Print("Pilih metode: ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		selectionSort()
		fmt.Println("Proyek berhasil diurutkan berdasarkan Dana!")
	case "2":
		insertionSort()
		fmt.Println("Proyek berhasil diurutkan berdasarkan Donatur!")
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func selectionSort() {
	n := len(daftarProyek)
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

func insertionSort() {
	n := len(daftarProyek)
	for i := 1; i < n; i++ {
		key := daftarProyek[i]
		j := i - 1
		for j >= 0 && daftarProyek[j].Donatur > key.Donatur {
			daftarProyek[j+1] = daftarProyek[j]
			j = j - 1
		}
		daftarProyek[j+1] = key
	}
}

func proyekBerhasilPendanaan() {
    fmt.Println("\n-------------------------------------------------------------------------------")
    fmt.Println("                            PROYEK BERHASIL DIDANAI                            ")
    fmt.Println("-------------------------------------------------------------------------------")
    found := false
    for i := 0; i < len(daftarProyek); i++ {
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
    var i int
	fmt.Printf("\n%d. %s (%s)\n", i+1, p.Nama, p.Kategori)
	fmt.Printf("   Target: Rp%.0f\n", p.Target)
	fmt.Printf("   Dana Terkumpul: Rp%.0f\n", p.Dana)
    fmt.Printf("   Jumlah Donatur: %d donatur\n", p.Donatur)
	fmt.Printf("   Progress: %.0f%%\n", (p.Dana/p.Target)*100)
}