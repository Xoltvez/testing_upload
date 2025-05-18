package main

import (
	"fmt"
)

func bubbleSort(arr []int) int {
	swaps := 0
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
				fmt.Println("Swap:", arr)
			}
		}
	}
	return swaps
}

func main() {
	nilai := []int{75, 60, 90, 85, 70}
	fmt.Println("Nilai sebelum diurutkan:", nilai)
	jumlahSwap := bubbleSort(nilai)
	fmt.Println("Nilai yang terurut", nilai)
	fmt.Println("Jumlah Pertukaran", jumlahSwap)
}
