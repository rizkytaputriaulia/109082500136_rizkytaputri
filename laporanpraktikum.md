# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Rizkyta Putri Aulia] - [109082500136]</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### soal1.go

```go
package main

import "fmt"

const MAX = 1000

func main() {
	var arr [MAX]float64
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	min := arr[0]
	max := arr[0]

	for j := 1; j < n; j++ {
		if arr[j] < min {
			min = arr[j]
		}
		if arr[j] > max {
			max = arr[j]
		}
	}
	fmt.Println(min, max)
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul10/output/output-soal1.png)
program tersebut di gunakan untuk mendata berat anak kelinci. pertama-tama membuat program array yang dimana aarray tersebut menyimpan kapasitas maksimum sebesar 1000 kapasitas tersebut tidak dapat di ubah. kemudian membuat program func main, didalam func main membuat variabel arr dan n, kemudian pengguna menginputkan 1 bilangan yang dimana bilangan tersebut disimpan dalam variabel n, dan inputan n tersebut untuk menentukan jumlah kelinci. Kemudian masuk kedalam code perulangan yang dimana perulangan tersebut datanya disimpan pada array. kemudian setelah array tersimpan program mulai mencari berat kelinci terbesar dan terkecil, pencarian tersebut menggunakan percabangan, dimana jika nilai lebih kecil dari min maka nilai akan menggantikan min dan jika nilai lebih besar dari max maka akan menggantikan max dan perulangan tersebut akan terus berjalan sampai elemen dalam array diperiksa semua. kemudian program menampilkan berat terkecil dan berat terbesar

## Unguided 

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### soal2.go

```go

package main

import "fmt"

const MAX = 1000

func main() {
	var arr [MAX]float64
	var hasil [MAX]float64
	var x, y int
	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}
	idx := 0
	wadah := 0
	for idx < x {
		total := 0.0
		hitung := 0

		for hitung < y && idx < x {
			total += arr[idx]
			idx++
			hitung++
		}
		hasil[wadah] = total
		wadah++
	}
	for j := 0; j < wadah; j++ {
		fmt.Print(hasil[j], " ")
	}
	fmt.Println()
	penjumlahan := 0.0
	for k := 0; k < wadah; k++ {
		penjumlahan += hasil[k]
	}
	rata := penjumlahan / float64(wadah)
	fmt.Println(rata)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul10/output/output-soal1.png)
program tersebut di gunakan untuk menentukan tarif ikan yang akan dijual dipasar. pertamaa tama membuat program array terlebih dahulu yang dimana kapasitas maksimalnya 1000. kemudian membuat func main didalam func main terdapat variabel arr, variabel hasil dan variabel x dan y. kemudian pengguna menginputkan 2 bilangan yaitu bilangan x menyatakan jumlah ikan dan jumlah y menyatakan kapasitas dalam 1 wadah. kemudian program masuk kedalam perulangan idx < x, dan didalam tersebut terdapat perulangan lagi yaitufor hitung < y && idx > x. fungsi perulangan ini untuk memastikan jumlah ikan dalam satu wadah dan tidak melebihi kapasitas y. dimana variabel idx ditambah ke variabel total dan idx akan terus bertambah ke ikan berikutnya dan variabel hitung akan bertambah untuk mencatat jumlah ikan dalam wadah. kemudian terdapat perulangan lagi dan setelah itu nilai total disimpan dalam array dan setelah ikan di kelompokan program akan mencetak hail pada posisi jumlah wadah.  kemudian program membuat perulangan lagi dan membuat perhitungan rata-rata  berat per wadah. dimana rumusnya rata = penjumlahan dibagi dengan jumlah wadah. 

## Unguided

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data
berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data
yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go

package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, n int, bMin, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func rerata(arr arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &bMin, &bMax)
	r := rerata(data, n)

	fmt.Printf("\nBerat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Rerata balita : %.2f kg\n", r)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul10/output/output-soal1.png)
Program ini dibuat untuk membantu petugas posyandu dalam mengolah data berat badan balita yang sudah dicatat. Data yang dimasukkan akan disimpan ke dalam sebuah array, lalu program akan membaca satu per satu untuk mencari nilai terkecil dan terbesar. Selain itu, program juga menghitung rata-rata dari seluruh data yang ada supaya bisa memberikan gambaran umum kondisi berat badan balita yang diperiksa. pengguna diminta memasukkan jumlah data lalu menginput berat balita satu per satu. Setelah semua data masuk, program langsung memproses tanpa langkah yang rumit dan menampilkan hasilnya dengan format yang rapi. Hasil akhir yang ditampilkan berupa berat minimum, maksimum, dan rata-rata sehingga memudahkan petugas dalam melihat perbandingan data yang sudah dikumpulkan.