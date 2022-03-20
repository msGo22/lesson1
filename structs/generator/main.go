package main

import (
	"fmt"
	"github.com/msGo22/lesson1/structs/tanimlama"
	"reflect"
)

func main()  {
	/**
	Yalın tanımlama (Önerilmez)
	sevket := tanimlama.Ogrenci{
		Adi:       "Sevket",
		Soyadi:    "Yılmaz",
		Numarasi:  "2",
	}
	 */
	// Generator ile tanımalama
	sevket := tanimlama.NewOgrenci("Şevket", "Yılmaz", 24)
	alameddin := tanimlama.NewOgrenci("Alameddin", "Çelik", 24)
	fmt.Println(reflect.TypeOf(sevket))
	sevket.SetVizeNotu(87)
	sevket.SetFinalNotu(66)

	fmt.Println("Vize Dış: ",sevket.GetVizeNotu())
	fmt.Println("Final Dış: ",sevket.GetFinalNotu())
	fmt.Println("Ortalama Dış: ",sevket.OrtalamaHesapla(false))


	ogrenciler := map[string]tanimlama.Ogrenci{
		sevket.Adi : sevket,
	}

	for k,v := range ogrenciler{
		fmt.Println("ogrenci Adı:", k)
		fmt.Println("ogrenci Vize Notunu Göster", v.GetVizeNotu())
	}

	ogrenciDurumları := map[*tanimlama.Ogrenci]bool{
		&sevket: true,
		&alameddin: false,
	}

	for k,v := range ogrenciDurumları{
		if v{
			fmt.Println(k.Adi)
		}
	}




}
