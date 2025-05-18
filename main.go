package main

import "fmt"

type dataLangganan struct {
	nama           string
	metode         string
	biaya          float64
	tanggal, bulan int
	status         string
}

const NMAX int = 10

type tabLangganan [NMAX]dataLangganan

func main() {
	var A tabLangganan

}

func ubahData(A *tabLangganan) {
	var j int
	//prosedur print tabel
	fmt.Println("pilih nomor yg akan diubah")
	fmt.Scan(&j)
	fmt.Println("Masukkan Langganan Baru")
	fmt.Scan(&A[j].nama, &A[j].biaya, &A[j].tanggal, &A[j].bulan)
}

func hapusData(A *tabLangganan, n *int) {
	var j, i int
	fmt.Println("pilih nomor yg akan dihapus")
	fmt.Scan(&j)
	for i = j; i <= *n; i++ {
		A[j] = A[j+1]
	}
	*n--
}

func tampilData(A tabLangganan, n int) {
	var i int
	for i = 0; i <= n; i++ {
		fmt.Println(A[i].nama, A[i].metode, A[i].biaya, A[i].tanggal, A[i].bulan)
	}
}
