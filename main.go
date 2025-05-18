package main

import "fmt"

// Struktur untuk data langganan
type dataLangganan struct {
	nama    string
	metode  string
	biaya   float64
	tanggal int
	status  string
}

// Konstanta maksimal langganan
const NMAX int = 10

// Array langganan
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
			
		case 3:
			
		case 4:
			TampilkanLangganan(A, jumlah)
		case 5:

		case 6:
            SelectionSortBiaya(&A, jumlah)
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

// Fungsi untuk menambah langganan
func MenambahLangganan(A *TabLangganan, jumlah *int) {
    if *jumlah >= NMAX {
        fmt.Println("Data langganan penuh!")
        return
    }
	
	fmt.Println("=== Tambah Langganan ===")
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[*jumlah].nama)
	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&A[*jumlah].metode)
	fmt.Print("Biaya Bulanan: ")
	fmt.Scan(&A[*jumlah].biaya)
	fmt.Print("Tanggal Pembayaran (1-31): ")
	fmt.Scan(&A[*jumlah].tanggal)
	fmt.Print("Status (aktif/nonaktif): ")
	fmt.Scan(&A[*jumlah].status)
     *jumlah++
    fmt.Println("Langganan berhasil ditambahkan.")
}



func SelectionSortBiaya(A *TabLangganan, N int) {
    var i, idx, pass int
    var temp dataLangganan
    pass = 1

    for pass < N {
        idx = pass 
        i = pass-1

        for i < N {
          if A[i].biaya > A[idx].biaya {
               idx = i
            }
            i++
       }

        temp = A[pass-1]
        A[pass-1] = A[idx]
        A[idx] = temp
		pass++
    
}
}


// Fungsi untuk menampilkan semua data langganan
func TampilkanLangganan(A TabLangganan, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data langganan.")
		return
	}

	fmt.Println("=== Daftar Langganan ===")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. Nama: %s\n", i+1, A[i].nama)
		fmt.Printf("   Metode Pembayaran: %s\n", A[i].metode)
		fmt.Printf("   Biaya Bulanan: %.2f\n", A[i].biaya)
		fmt.Printf("   Tanggal Pembayaran: %d\n", A[i].tanggal)
		fmt.Printf("   Status: %s\n", A[i].status)
		fmt.Println("-------------------------")
	}
	fmt.Printf("Total langganan: %d\n", jumlah)
}