package main

import "fmt"

func main()  {
	sayi := 22
	sayininAdresi := &sayi
	sayininAdresininAdresi := &sayininAdresi
	fmt.Println("Değer:", sayi)
	fmt.Println("Adres: ", sayininAdresi)
	fmt.Println("Adresin Adresi:", sayininAdresininAdresi)
	// fonksiyon
	SayiArttir(&sayi)
	fmt.Println(sayi)
	KareHesapla(&sayi)
	fmt.Println(sayi)
}


// şevket -> numarası değiştirdiğğimiz
// 3 x 128 tanimlama

// -kimya
// alameddin (8 byte)
// şevket  (8 byte)
// - Matematik
// Yavuz (8 byte)
// şevket (8 byte)


func SayiArttir(n *int)  {
	*n = *n + 5
	fmt.Println("İç Değer", *n)
}

func KareHesapla(n *int)  {
	*n = *n * *n
}


// haftalık antrenman: İki değişkenin adreslerini swap yapan fonksiyon