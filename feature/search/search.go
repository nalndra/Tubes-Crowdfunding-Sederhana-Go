package search

import (
	"fmt"

	project "github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/dataProyek"
)

func CariProyek() {
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

	for i := 0; i < len(project.DaftarProyek); i++ {
		if project.DaftarProyek[i].Nama == keyword {
			project.TampilkanDetailProyek(project.DaftarProyek[i])
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
	high := len(project.DaftarProyek) - 1
	found := false

	for low <= high {
		mid := low + (high-low)/2

		if project.DaftarProyek[mid].Nama == keyword {
			project.TampilkanDetailProyek(project.DaftarProyek[mid])
			found = true
			break
		} else if project.DaftarProyek[mid].Nama < keyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Proyek tidak ditemukan!")
	}
}
