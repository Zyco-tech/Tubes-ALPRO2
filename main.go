package main

import "fmt"

type dataLangganan struct {
	nama    string
	metode  string
	biaya   float64
	tanggal int
	status  string
}

const NMAX int = 10

type TabLangganan [NMAX]dataLangganan

func main() {
	var A TabLangganan
	var jumlah int
	var pilihan int

	for {
		menu()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			MenambahLangganan(&A, &jumlah)
		case 2:
			ubahData(&A)
		case 3:
			hapusData(&A, &jumlah)
		case 4:
			TampilkanLangganan(A, jumlah)
		case 5:

		case 6:

		case 7:

		case 8:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("pilihan tidak valid,coba lagi")
		}
	}
}

func menu() {
	fmt.Println("=======MENU LANGGANAN==========")
	fmt.Println("1. tambah data langganan")
	fmt.Println("2. ubah data langganan")
	fmt.Println("3. hapus data langganan")
	fmt.Println("4.	tampilkan tabel data")
	fmt.Println("5.	search")
	fmt.Println("6.	sort data")
	fmt.Println("7.	summary pengeluaran perbulan")
	fmt.Println("8.Exit")
	fmt.Println("===============================")

}
