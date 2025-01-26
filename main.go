package main

import "fmt"

func ReverseString(input string) string {
	n := len(input)
	balik := make([]byte, n)
	for i := 0; i < n; i++ {
		balik[i] = input[n-1-i]
	}
	return string(balik)
}

func main() {
	var input string
	fmt.Print("Masukan Kaalimat : ")
	fmt.Scanln(&input)
	balik := ReverseString(input)
	fmt.Println("Hasil Pembalikan : ", balik)
}
