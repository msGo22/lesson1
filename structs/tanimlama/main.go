package tanimlama

import "fmt"

type Ogrenci struct {
	Adi        string `json:"adi,omitempty"` // gelirse gösteririm gelmezse saklarım
	Soyadi     string `json:"soyadi"` // Gelirse gösteririm, gelmezse boş gösteririm
	Numarasi   string `json:"okul_numarasi"`
	Yas        int    `json:"yas"`
	Sinif      int
	VizeNotu   float64
	FinalNotu  float64
	Memleketi  string
	Hobileri   []string
	KanGrubu   string
	MedeniHali string
}

// Tanımlama Fonksiyonu
func NewOgrenci(adi, soyadi string, yas int) Ogrenci {
	return Ogrenci{
		Adi:    adi,
		Soyadi: soyadi,
		Yas:    yas,
		Sinif:  20,
	}
}

// Vize Not
func (o *Ogrenci) SetVizeNotu(vizeNotu float64) {
	o.vizeNotu = vizeNotu
	fmt.Println("İç", o.vizeNotu)
}

func (o *Ogrenci) GetVizeNotu() float64 {
	return o.vizeNotu
}

// Final Not
func (o *Ogrenci) SetFinalNotu(finalNotu float64) {
	o.finalNotu = finalNotu
	fmt.Println("İç", o.finalNotu)
}

func (o *Ogrenci) GetFinalNotu() float64 {
	return o.finalNotu
}

// Ek Fonksiyonlar
func (o *Ogrenci) OrtalamaHesapla(isKanaatEklenecek bool) float64 {
	if isKanaatEklenecek {
		o.kanaatPuaniEkle()
	}
	return (o.vizeNotu * 0.4) + (o.finalNotu * 0.6)
}

// İç Fonksiyonlar
func (o *Ogrenci) kanaatPuaniEkle() {
	o.finalNotu *= 1.2
	if o.finalNotu > 100 {
		o.finalNotu = 100
	}
}
