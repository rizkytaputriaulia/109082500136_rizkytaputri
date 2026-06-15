# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Rizkyta Putri Aulia] - [NIM]<109082500136>

## Unguided 

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu.
Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort.
Masukan dimulai dengan sebuah integer 􁢾􁢾 (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi 􁢾􁢾 baris berikutnya selalu dimulai dengan sebuah integer 􁢾􁢾 (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.
#### 1.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	var A arrInt
	i := 0
	for i < n {
		fmt.Scan(&m)
		j := 0
		for j < m {
			fmt.Scan(&A[j])
			j++
		}
		selectionSort(&A, m)
		j = 0
		for j < m {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(A[j])
			j++
		}
		fmt.Println()
		i++
	}
}

```

##### Output 
![Screenshot Output 1.go] (https://github.com/rizkytaputriaulia/1.go.ss/blob/main/Screenshot%202026-06-15%20231451.png)
[penjelasan] : mengurutkan sekumpulan angka dari yang terkecil ke terbesar (ascending) sebanyak beberapa baris kelompok data sesuai keinginan pengguna. program pertama-tama akan meminta jumlah baris kelompok data dan jumlah angka di setiap barisnya , lalu setelah angka-angka tersebut diinput, program langsung memanggil fungsi selectionSort untuk mengurutkannya. lalu menukar posisinya ke urutan paling depan, dan proses ini terus diulang pada sisa angka yang belum terurut hingga semuanya rapi. Setelah proses pengurutan selesai, program akan langsung mencetak angka-angka yang sudah rapi tersebut ke samping dengan pemisah spasi, kemudian otomatis lanjut memproses kelompok baris data berikutnya sampai seluruh baris selesai diproses.


### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

#### 2.go

```go
package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j++
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i++
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int
	i = 1
	for i <= n-1 {
		idx_max = i - 1
		j = i
		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j++
		}
		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i++
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	i := 0
	for i < n {
		fmt.Scan(&m)

		var ganjil, genap arrInt
		ng := 0
		nge := 0

		j := 0
		for j < m {
			var x int
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil[ng] = x
				ng++
			} else {
				genap[nge] = x
				nge++
			}
			j++
		}

		selectionSortAsc(&ganjil, ng)
		selectionSortDesc(&genap, nge)

		j = 0
		for j < ng {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			j++
		}

		j = 0
		for j < nge {
			if ng > 0 || j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			j++
		}
		fmt.Println()
		i++
	}
}


```

##### Output 
![Screenshot Output 2.go] (https://github.com/rizkytaputriaulia/2.go.ss/blob/main/Screenshot%202026-06-15%20231518.png)
[penjelasan] : berfungsi untuk memisahkan angka ganjil dan genap dari sekumpulan data input, lalu mengurutkannya dengan aturan yang berbeda sebelum dicetak kembali dalam satu baris. untuk mengurutkan angka genap dari yang terbesar ke terkecil (descending). Di akhir proses untuk setiap barisnya, program langsung mencetak kelompok angka ganjil yang sudah rapi terlebih dahulu, disusul langsung oleh kelompok angka genap di sebelah kanannya dengan pemisah spasi, lalu otomatis berganti baris untuk memproses kelompok data berikutnya.


### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?


#### 3.go

```go
package main

import "fmt"

const NMAX = 1000001

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var A arrInt
	n := 0
	var x int

	for {
		fmt.Scan(&x)
		if x == -5313 {
			break
		}
		if x == 0 {
			insertionSort(&A, n)
			if n%2 == 1 {
				fmt.Println(A[n/2])
			} else {
				fmt.Println((A[n/2-1] + A[n/2]) / 2)
			}
		} else {
			A[n] = x
			n++
		}
	}
}

```

##### Output 
![Screenshot Output 3.go] 
9https://github.com/rizkytaputriaulia/3.go/blob/main/Screenshot%202026-06-15%20232130.png
[penjelasan] : berfungsi untuk menghitung nilai tengah (median) dari sekumpulan angka yang dimasukkan secara dinamis oleh pengguna, lalu program mengecek jumlah datanya: jika ganjil, program langsung mencetak angka yang berada tepat di posisi tengah

### 4. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.

#### 4.go

```go
package main

import "fmt"

const NMAX = 10000

type arrInt [NMAX]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}

func main() {
	var A arrInt
	n := 0
	var x int

	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		A[n] = x
		n++
	}

	insertionSort(&A, n)

	i := 0
	for i < n {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(A[i])
		i++
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	jarak := A[1] - A[0]
	tetap := true
	i = 2
	for i < n {
		if A[i]-A[i-1] != jarak {
			tetap = false
			break
		}
		i++
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

```

##### Output 
![Screenshot Output 4.go] (https://github.com/rizkytaputriaulia/4.go.ss/blob/main/Screenshot%202026-06-15%20232428.png)
[penjelasan] : berfungsi untuk memeriksa apakah sekumpulan bilangan bulat positif yang dimasukkan oleh pengguna membentuk pola barisan aritmetika atau memiliki selisih (jarak) yang sama antarangkanya setelah diurutkan. untuk menyusun angka-angka tersebut dari yang terkecil ke terbesar, lalu langsung mencetaknya ke layar dalam satu baris dengan pemisah spasi. Terakhir, jika data yang dimasukkan berjumlah minimal dua angka, program akan menghitung selisih antara angka pertama dan kedua sebagai acuan, lalu memeriksa sisa angka lainnya; jika semua angka terbukti memiliki selisih yang konsisten, program akan memunculkan pesan "Data berjarak [nilai jarak]", namun jika ada satu saja selisih yang berbeda atau datanya kurang dari dua, program akan memunculkan pesan Data berjarak tidak tetap.


### 5. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu

```go
package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	i := 0
	for i < *n {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
		i++
	}
}

func CetakTerfavorit(pustaka *DaftarBuku, n int) {
	idx := 0
	i := 1
	for i < n {
		if pustaka[i].rating > pustaka[idx].rating {
			idx = i
		}
		i++
	}
	fmt.Println(pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
		i++
	}
}

func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	i := 0
	for i < limit {
		fmt.Println(pustaka[i].judul)
		i++
	}
}

func CariBuku(pustaka *DaftarBuku, n int, r int) {
	lo := 0
	hi := n - 1
	ketemu := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if pustaka[mid].rating == r {
			ketemu = mid
			break
		} else if pustaka[mid].rating > r {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		b := pustaka[ketemu]
		fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int

	DaftarkanBuku(&pustaka, &nPustaka)

	CetakTerfavorit(&pustaka, nPustaka)
	UrutBuku(&pustaka, nPustaka)
	Cetak5Terbaru(&pustaka, nPustaka)

	var r int
	fmt.Scan(&r)
	CariBuku(&pustaka, nPustaka, r)
}

```

##### Output 
![Screenshot Output 5.go] (https://github.com/rizkytaputriaulia/5.go.ss/blob/main/Screenshot%202026-06-15%20232915.png)
[penjelasan] : berfungsi untuk mengelola data perpustakaan digital sederhana dengan fitur pencatatan, pencarian, dan pengurutan buku berdasarkan ratingnya. Di bagian akhir, program akan meminta pengguna memasukkan satu angka rating lagi untuk dicari di dalam perpustakaan menggunakan metode binary search lewat fungsi CariBuku; jika rating tersebut ditemukan, detail buku akan langsung dicetak, namun jika tidak, program akan memunculkan pesan "Tidak ada buku dengan rating seperti itu".