# Rupiah

**Rupiah** adalah pustaka Golang yang membantu Anda menampilkan penulisan mata uang rupiah dengan benar sesuai EBI (Ejaan Bahasa Indonesia).

Aturan Penulisan
-
Berikut beberapa aturan penulisan mata uang Rupiah sesuai [PUEBI](http://badanbahasa.kemdikbud.go.id/lamanbahasa/sites/default/files/PUEBI.pdf) (Pedoman Umum Ejaan Bahasa Indonesia):

* Lambang mata uang **tidak diikuti tanda titik**.

    Misalnya:
    
    Rp
    
    $
    
* Angka yang menunjukkan bilangan besar dapat ditulis sebagian dengan huruf supaya lebih mudah dibaca.

    Misalnya:  
   
    Proyek pemberdayaan ekonomi rakyat itu memerlukan biaya 10 triliun rupiah.
* Penulisan bilangan dengan angka dan huruf sekaligus dilakukan dalam peraturan perundang-undangan, akta, dan kuitansi.
    
    Misalnya: 
    
    Setiap orang yang menyebarkan atau mengedarkan rupiah tiruan, sebagaimana dimaksud dalam Pasal 23 ayat (2), dipidana dengan pidana kurungan paling lama 1 (satu) tahun dan pidana denda paling banyak Rp200.000.000,00 (dua ratus juta rupiah).
* Penulisan bilangan yang dilambangkan dengan angka dan diikuti huruf dilakukan seperti berikut.

    Misalnya:
    
    Saya lampirkan tanda terima uang sebesar Rp900.500,50 (sembilan ratus ribu lima ratus rupiah lima puluh sen).
* Tanda koma dipakai sebelum angka desimal atau di antara rupiah dan sen yang dinyatakan dengan angka.

    Misalnya:
    
    Rp500,50
    
    Rp750,00

Pasang
-

``` bash
$ go get github.com/paimanbandi/rupiah
```

Fitur
-
Mengubah bilangan bertipe _float64_ dan _int64_ ke dalam bentuk _string_ penulisan mata uang rupiah sesuai Ejaan Bahasa Indonesia.

Peringatan
-
* Masukan bilangan bertipe _float64_ nilainya akan dibuat presisi mengikuti ketentuan penulisan sen, yaitu 2 angka di belakang dan tidak dibulatkan.
* Masukan bilangan bertipe _int64_ nilainya adalah dalam rupiah saja dan nilai sen adalah nol.

Penggunaan
-

```go
import "github.com/paimanbandi/rupiah"

func main() {
  fmt.Println(rupiah.FormatFloat64ToRp(5232154989000.50934))
  // Rp5.232.154.989.000,50

  fmt.Println(rupiah.FormatFloat64ToRp(-4154989000.00234))
  // -Rp4.154.989.000,00

  fmt.Println(rupiah.FormatFloat64ToRp(962000.23))
  // Rp962.000,23

  fmt.Println(rupiah.FormatFloat64ToRp(-4981000.00))
  // -Rp4.981.000,00

  fmt.Println(rupiah.FormatInt64ToRp(689025876))
  // Rp689.025.876,00

  fmt.Println(rupiah.FormatInt64ToRp(-73475689025))
  // -Rp73.475.689.025,00

  fmt.Println(rupiah.FormatFloat64ToWords(58765653454.67574))
  // lima puluh delapan miliar tujuh ratus enam puluh lima juta enam ratus lima puluh tiga ribu empat ratus lima puluh empat rupiah enam puluh tujuh sen

  fmt.Println(rupiah.FormatFloat64ToWords(-653454.564))
  // minus enam ratus lima puluh tiga ribu empat ratus lima puluh empat rupiah lima puluh enam sen

  fmt.Println(rupiah.FormatFloat64ToWords(76558765653454.67))
  // tujuh puluh enam triliun lima ratus lima puluh delapan miliar tujuh ratus enam puluh lima juta enam ratus lima puluh tiga ribu empat ratus lima puluh empat rupiah enam puluh tujuh sen

  fmt.Println(rupiah.FormatFloat64ToWords(-8756787657.04))
  // minus delapan miliar tujuh ratus lima puluh enam juta tujuh ratus delapan puluh tujuh ribu enam ratus lima puluh tujuh rupiah empat sen

  fmt.Println(rupiah.FormatInt64ToWords(653454))
  // enam ratus lima puluh tiga ribu empat ratus lima puluh empat rupiah

  fmt.Println(rupiah.FormatInt64ToWords(-4324398700))
  // minus empat miliar tiga ratus dua puluh empat juta tiga ratus sembilan puluh delapan ribu tujuh ratus rupiah

}
```

Lisensi
-
Lisensi MIT. Silakan lihat File LICENSE untuk informasi lebih lanjut.
