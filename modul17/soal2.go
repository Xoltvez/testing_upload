package main

import "fmt"

func inputData() (string, int) {
	var x string
	var n int
	fmt.Scan(&x)
	fmt.Scan(&n)
	return x, n
}

func bacaDataString(n int) []string {
	var data = make([]string, n)
	var i int

	fmt.Println("Masukkan", n, "string:")
	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
	return data
}

func cariString(data []string, x string) (bool, int, int) {
	var ditemukan bool = false
	var jumlah, posisi, i int
	posisi = -1

	for i = 0; i < len(data); i++ {
		if data[i] == x {
			jumlah++
			if !ditemukan {
				ditemukan = true
				posisi = i + 1
			}
		}
	}
	return ditemukan, posisi, jumlah
}

func tampilkanHasil(x string, ditemukan bool, posisi int, jumlah int) {
	if ditemukan {
		fmt.Println("a. String ditemukan.")
		fmt.Println("b. Ditemukan pada posisi ke:", posisi)
	} else {
		fmt.Println("a. String tidak ditemukan.")
		fmt.Println("b. Tidak ditemukan, jadi tidak ada posisi.")
	}
	fmt.Println("c. Jumlah kemunculan string:", jumlah)
	if jumlah >= 2 {
		fmt.Println("d. Ada sedikitnya dua string", x)
	} else {
		fmt.Println("d. Tidak ada sedikitnya dua string", x)
	}
}

func main() {
	var x string
	var n, posisi, jumlah int
	var data []string
	var ditemukan bool

	x, n = inputData()
	data = bacaDataString(n)
	ditemukan, posisi, jumlah = cariString(data, x)
	tampilkanHasil(x, ditemukan, posisi, jumlah)
}
