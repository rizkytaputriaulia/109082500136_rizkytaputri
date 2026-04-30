# <h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center"> Rizkyta putri aulia - 109082500136 </p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila
diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y)
berdasarkan dua lingkaran tersebut.


#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func diDalam(l Lingkaran, t Titik) bool {
	dx := float64(t.x - l.pusat.x)
	dy := float64(t.y - l.pusat.y)
	jarak := math.Sqrt(dx*dx + dy*dy)
	return jarak <= float64(l.r)
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var x, y int

	
	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	l1 := Lingkaran{Titik{x1, y1}, r1}
	l2 := Lingkaran{Titik{x2, y2}, r2}
	t := Titik{x, y}

	dalam1 := diDalam(l1, t)
	dalam2 := diDalam(l2, t)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
### Output Unguided :

##### Output 
! (https://github.com/rizkytaputriaulia/ss_mod9.1.go/blob/main/Screenshot%202026-05-01%20001057.png)

Program ini dibuat untuk menentukan posisi sebuah titik terhadap dua lingkaran. Setiap lingkaran memiliki titik pusat dan jari-jari, lalu diberikan satu titik sembarang yang akan dicek posisinya. Cara kerjanya dengan menghitung jarak antara titik tersebut ke pusat masing-masing lingkaran, kemudian dibandingkan dengan panjang jari-jarinya. Dari situ bisa diketahui apakah titik berada di dalam atau di luar lingkaran. 




### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program
yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array
memiliki kapasitas penyimpanan data sejumlah elemen tertentu.
#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("Indeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("Indeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Println("Indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapus int
	fmt.Print("Masukkan indeks yang mau dihapus: ")
	fmt.Scan(&hapus)

	arr = append(arr[:hapus], arr[hapus+1:]...)

	fmt.Println("Array setelah dihapus:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	total := 0
	for _, v := range arr {
		total += v
	}
	rata := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", rata)

	var jumlah float64
	for _, v := range arr {
		selisih := float64(v) - rata
		jumlah += selisih * selisih
	}
	std := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi:", std)
}

```
### Output Unguided :

##### Output 
! (https://github.com/rizkytaputriaulia/mod9.2.go/blob/main/Screenshot%202026-05-01%20002731.png)
Program ini digunakan untuk mengolah sekumpulan bilangan yang disimpan dalam sebuah array. Pertama, pengguna diminta memasukkan jumlah data lalu mengisi elemen-elemennya. Setelah itu program bisa menampilkan isi array secara keseluruhan, lalu memisahkan elemen berdasarkan indeks ganjil dan genap. Selain itu, ada juga fitur untuk menampilkan elemen pada indeks kelipatan angka tertentu yang dimasukkan oleh pengguna. bisa menghapus salah satu elemen berdasarkan indeks yang dipilih, lalu menampilkan kembali isi array yang sudah diperbarui. Di bagian akhir, program menghitung nilai rata-rata dari data yang ada serta standar deviasinya untuk melihat sebaran nilainya.



### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang
memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang
digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
#### soal3.go

```go
package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	p := 1

	var hasil []string

	for {
		fmt.Print("Pertandingan ", p, ": ")
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		p++
	}

	fmt.Println("\nHasil:")
	for i, v := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, v)
	}
}

```
### Output Unguided :

##### Output 
! (https://github.com/rizkytaputriaulia/mod9.3.go/blob/main/Screenshot%202026-05-01%20003855.png)
Program ini digunakan untuk mencatat hasil pertandingan antara dua klub bola yang dimainkan beberapa kali. Pengguna diminta memasukkan nama kedua klub terlebih dahulu, lalu memasukkan skor dari setiap pertandingan secara berulang. Setiap hasil pertandingan akan dicek, jika salah satu klub memiliki skor lebih tinggi maka nama klub tersebut akan disimpan, sedangkan jika skornya sama maka akan dianggap sebagai hasil seri atau draw. input akan berhenti jika salah satu atau kedua skor yang dimasukkan bernilai negatif. Setelah itu, program akan menampilkan daftar hasil pertandingan yang sudah direkap sebelumnya, baik itu nama klub yang menang maupun hasil draw.



### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk
membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa
apakah membentuk palindrom.
#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	fmt.Scanf("%c", &input)
	for input != '.' && *n < NMAX {
		t[*n] = input
		*n++
		fmt.Scanf("%c", &input)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {

	var tAsli tabel = t
	balikanArray(&t, n)

	for i := 0; i < n; i++ {
		if tAsli[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks           : ")
	isiArray(&tab, &m)

	isPal := palindrom(tab, m)

	fmt.Print("Reverse teks   : ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Printf("Palindrom      ? %t\n", isPal)
}

```
### Output Unguided :

##### Output 
! (https://github.com/rizkytaputriaulia/mod9.4.go/blob/main/Screenshot%202026-05-01%20005341.png)
Program ini dirancang untuk mengelola sekumpulan karakter yang disimpan dalam sebuah array dengan kapasitas maksimal 127 elemen. Alur utamanya dimulai dari proses input karakter satu per satu oleh pengguna, di mana pengisian akan otomatis berhenti jika ditemukan karakter titik atau saat batas maksimum array sudah terpenuhi. Setelah data tersimpan, program memiliki kemampuan untuk memproses sekumpulan karakter. pembalikan, posisi setiap karakter akan ditukar sehingga urutannya menjadi terbalik dari posisi semula. Bersamaan dengan itu, program akan memverifikasi apakah urutan karakter tersebut tetap sama jika dibaca dari depan maupun dari belakang, yang menentukan apakah kumpulan karakter tersebut memenuhi syarat sebagai sebuah palindrom atau tidak.