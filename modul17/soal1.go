package main

import "fmt"

func main() {
	var angka, ratarata float64
	var jumlah float64 = 0
	var banyak int = 0

	for {
		fmt.Scan(&angka)
		if angka == 9999 {
			break
		}
		jumlah += angka
		banyak++
	}

	if banyak > 0 {
		ratarata = jumlah / float64(banyak)
		fmt.Println(ratarata)
	} else {
		fmt.Println("Tidak ada data untuk dihitung.")
	}
}
