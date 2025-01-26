# Reverse String

Program ini membalikkan urutan string menggunakan bahasa Go.

## Penjelasan Singkat
Fungsi ReverseString menerima string sebagai input, menghitung panjangnya, lalu membuat sebuah buffer berupa slice of byte dengan panjang yang sama. Selanjutnya, iterasi dilakukan untuk menyalin karakter dari string input ke buffer secara terbalik, di mana karakter terakhir dari input ditempatkan di indeks pertama buffer, dan seterusnya. Setelah seluruh karakter disalin, buffer tersebut dikonversi kembali menjadi string terbalik yang akan dikembalikan sebagai output.

## Contoh Penggunaan
```go
input := "appel"
fmt.Println(ReverseString(input)) // Output: leppal