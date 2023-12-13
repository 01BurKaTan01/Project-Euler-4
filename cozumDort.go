package main

import "fmt"

func main() {
	sayac := 100
	enbuyuk := 0

	for i := 100; i < 1000; i++ {

		for j := 999; j > sayac; j-- {
			if palindromdur((i * j)) {
				enbuyuk = buyuguAl(enbuyuk, i*j)
			}
		}

		sayac++

	}

	fmt.Println(enbuyuk)
}

func buyuguAl(bir int, iki int) int {
	if bir > iki {
		return bir
	}

	return iki
}

func palindromdur(sayi int) bool {
	kayit := sayi
	ters := 0
	rakam := sayi % 10

	for sayi > 0 {
		ters = ters*10 + rakam
		sayi /= 10
		rakam = sayi % 10
	}

	return kayit == ters
}
