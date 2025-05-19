package sort

import (
	"fmt"

	project "github.com/nalndra/Tubes-Crowdfunding-Sederhana-Go/feature/dataProyek"
)

func UrutkanProyek() {
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
	n := len(project.DaftarProyek)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if project.DaftarProyek[j].Dana < project.DaftarProyek[minIdx].Dana {
				minIdx = j
			}
		}
		project.DaftarProyek[i], project.DaftarProyek[minIdx] = project.DaftarProyek[minIdx], project.DaftarProyek[i]
	}
}

func insertionSort() {
	n := len(project.DaftarProyek)
	for i := 1; i < n; i++ {
		key := project.DaftarProyek[i]
		j := i - 1
		for j >= 0 && project.DaftarProyek[j].Donatur > key.Donatur {
			project.DaftarProyek[j+1] = project.DaftarProyek[j]
			j = j - 1
		}
		project.DaftarProyek[j+1] = key
	}
}
