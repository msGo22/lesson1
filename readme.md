###Ders içeriği
- Pointers
  - Pointer nedir
  - Pointer gösterimi
  - Fonksiyon içerisinde kullanımları
  
- Veri yapıları
  - Struct Nedir?
  - Oluşturma Şekilleri
    - Standart
    - New
    - Generator
  - Struct Türleri
    - Defined Struct
    - Anoymus Struct
    - Nested Struct
    
  - Tag aktarımı
    - Json
    - Custom Tags
 ### Ödev: 

Bir kargo firması için tasarım çıkartıyoruz Kullanıcılar bir kargo talebi oluşturmak için tckimlik numarsı, adı soyadı adresleri ve telefon numaralarını girecekler ve teslim edecekleri kişininde bilgilerini girecekler. Bu bilgiler sipariş struct’ına kayıt edilmeli. Sipariş durumu için bir alan tutmalıyız ve siparişin tesliminde bu alanı teslim edildi şeklinde değiştirmeliyiz. Kullanıcının bilgileri inmemory bir alanda müşteriler olarak tutulmalı ve eğer bir müşterinin bilgileri değiştiyse aktif tüm siparişlerindeki adresi anlık güncellenmelidir. (Eski siparişlerinde de güncellenebilir*)

Sahip olması gereken özellikler
1.	Müşteri Sipariş oluşturma (İlk siparişte müşteri bilgileri ve alıcı bilgileri müşteriler değişkeninde tutulmalı)
2.	Kargo Durum sorgulama
3.	Müşteri Adres güncelleme
4.	Kargo firması kargoyu teslim edildi veya kargo geri gönderime çıktı olarak durumu güncelleyebilmelidir
5.

Not: InMemory kayıt olacağı için program kapandığında kayıtlarımızın silinmesini gözardı edebiliriz. 

