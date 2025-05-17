package main

import "fmt"

type Proyek struct {
	Nama     string
	Kategori string
	Target   float64
	Dana     float64
	Donatur  int
}

var daftarProyek []Proyek

func main() {
	// Inisialisasi data contoh
	daftarProyek = []Proyek{
		{"Robot Edukasi", "Pendidikan", 5000000, 3500000, 45},
		{"Aplikasi Kesehatan", "Kesehatan", 10000000, 7500000, 32},
		{"Buku Anak", "Pendidikan", 2000000, 2500000, 28},
	}

	var input string

	for {
		tampilkanMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			tampilkanSemuaProyek()
		case "2":
			tambahProyek()
		case "3":
			editProyek()
		case "4":
			hapusProyek()
		case "5":
			cariProyek()
		case "6":
			urutkanProyek()
		case "7":
			tampilkanProyekBerhasil()
		case "8":
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tampilkanMenu() {
	fmt.Println("\n=== APLIKASI CROWDFUNDING ===")
	fmt.Println("1. Tampilkan Semua Proyek")
	fmt.Println("2. Tambah Proyek")
	fmt.Println("3. Edit Proyek")
	fmt.Println("4. Hapus Proyek")
	fmt.Println("5. Cari Proyek")
	fmt.Println("6. Urutkan Proyek")
	fmt.Println("7. Proyek Berhasil")
	fmt.Println("8. Keluar")
}

func tampilkanSemuaProyek() {
	fmt.Println("\nDaftar Semua Proyek:")
	for i, p := range daftarProyek {
		fmt.Printf("%d. %s (%s)\n", i+1, p.Nama, p.Kategori)
		fmt.Printf("   Target: Rp%.0f\n", p.Target)
		fmt.Printf("   Dana terkumpul: Rp%.0f\n", p.Dana)
		fmt.Printf("   Jumlah Donatur: %d donatur\n", p.Donatur)
		fmt.Printf("   Progress: %.0f%%\n", (p.Dana/p.Target)*100)
	}
}

func tambahProyek() {
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

	daftarProyek = append(daftarProyek, p)
	fmt.Println("Proyek berhasil ditambahkan!")
}

func editProyek() {
	tampilkanSemuaProyek()
	var index int
	fmt.Print("Pilih nomor proyek yang akan diedit: ")
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= len(daftarProyek) {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	fmt.Printf("Edit Proyek: %s\n", daftarProyek[index].Nama)
	fmt.Print("Nama Baru (kosongkan jika tidak diubah): ")
	var nama string
	fmt.Scanln(&nama)
	if nama != "" {
		daftarProyek[index].Nama = nama
	}

	// Implementasi edit field lainnya bisa ditambahkan di sini
	fmt.Println("Proyek berhasil diupdate!")
}

func hapusProyek() {
	tampilkanSemuaProyek()
	var index int
	fmt.Print("Pilih nomor proyek yang akan dihapus: ")
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= len(daftarProyek) {
		fmt.Println("Nomor proyek tidak valid!")
		return
	}

	daftarProyek = append(daftarProyek[:index], daftarProyek[index+1:]...)
	fmt.Println("Proyek berhasil dihapus!")
}

func cariProyek() {
	fmt.Println("\nPilihan Pencarian:")
	fmt.Println("1. Sequential Search (Nama)")
	fmt.Println("2. Binary Search (Nama) - Harus urut")
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
    
    // Menggunakan for loop biasa dengan index
    for i := 0; i < len(daftarProyek); i++ {
        p := daftarProyek[i]
        
        // Pencarian exact match sederhana
        if p.Nama == keyword {
            tampilkanDetailProyek(p)
            found = true
        }
    }
    
    if !found {
        fmt.Println("Proyek tidak ditemukan!")
    }
}

func binarySearch(keyword string) {
	// Untuk binary search, data harus diurutkan terlebih dahulu
	fmt.Println("\nHasil Pencarian Binary:")
	// Implementasi binary search bisa ditambahkan di sini
	fmt.Println("Fitur ini membutuhkan data terurut terlebih dahulu")
}

func urutkanProyek() {
	fmt.Println("\nPilihan Pengurutan:")
	fmt.Println("1. Selection Sort (Dana)")
	fmt.Println("2. Insertion Sort (Donatur)")
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

func tampilkanProyekBerhasil() {
	fmt.Println("\nProyek yang Berhasil Mencapai Target:")
	found := false
	for _, p := range daftarProyek {
		if p.Dana >= p.Target {
			tampilkanDetailProyek(p)
			found = true
		}
	}
	if !found {
		fmt.Println("Belum ada proyek yang mencapai target")
	}
}

func tampilkanDetailProyek(p Proyek) {
	fmt.Printf("\n%s (%s)\n", p.Nama, p.Kategori)
	fmt.Printf("Target: Rp%.2f\n", p.Target)
	fmt.Printf("Dana terkumpul: Rp%.2f (%d donatur)\n", p.Dana, p.Donatur)
	fmt.Printf("Progress: %.1f%%\n", (p.Dana/p.Target)*100)
}