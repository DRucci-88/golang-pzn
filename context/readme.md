# Context

- a/ sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
- *Dibuat per request* (misalkan setiap http request masuk ke server web)
- Mempermudah kita meneruskan value dan sinyal antar proses
- Studi kasus: kalo kita ada data yang di passing ke function melalui parameter (props drilling)
- Studi kasus: Jika ada proses dibatalkan maka semua proses dibatalkan (goroutine) ,pake sinyal cancel

## Mengapa penting:
- Biasanya digunakan untuk *mengirim sinyal data request atau sinyal ke proses lain*
- Dengan context, ketika ingin membatalkan semua proses, kita cukup mengirim sinyal ke context, maka semua proses akan dibatalkan
- Cth: database, http server, http client, dll

## Membuat Context
- Context = interface, kita butuh sebuah struct yang sesuai dengan contract interface Context
- Tidak usah membuat secara manual, karena ada Golang package

Function membuat Context:
- *context.Background()*
    - Context kosong
    - Tidak pernah dibatalkan
    - Tidak pernah timeout
    - Tidak memiliki value apapun
    - Biasa digunakan di main funtion atau dalam test atau proses awal request terjadi
- *context.TODO()*
    - Context kosong seperti *context.Background()*
    - Biasanya digunakan ketika belum jelas mau dipake untuk apa

## Parent dan Child Context
- Konsep parent and child -> ada sifat pewarisan
- Bisa membuat child context dari context yang sudah ada
- Parent bisa punya banyak child, tapi child hanya bisa punya satu parent
- Mirip dengan pewarisan (inheritance) di OOP

## Hubungan Parent dan Child Context
- Selalu terhubung
- Pembatalan Context A, maka semua child dan sub child context A akan dibatalkan juga. Tapi parent Context A tidak dibatalkan.
- Menyisipkan data kedalam Context A, maka semua child dan subchild context A juga bisa mendapatkan data tersebut.  Tapi parent Context A tidak dapat datanya

## Context Immutable
- Context adalah object yang *Immutable*, setelah dibuat, tidak dapat diubah lagi
- Ketika menambahkan value kedalam context atau menambahkan pengaturan timeout etc, otomatis akan membentuk child context baru dan bukan merubah context tsb.

## Context With Value
- Saat awal membuat context, context tidak memiliki value
- Kita bisa *menambah sebuah value* dengan data Pair (key - value) kedalam context
- Saat menambah value ke context, otomatis akan tercipta child context baru, artinya original context tidak berubah sama sekali
- *context.WithValue(parent, key, value)* = menambah value kedalam context

## Context With Cancel
- *Menambahkan sinyal cancel*
- Biasanya ketika kita butuh menjalankan proses lain dan ingin member sinyal cancel ke proses tersebut
- Biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke context nya
- Ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap context nya, jika tidak, tidak ada gunanya
- *context.WithCancel(parent)* = context dengan cancel signal

## Context With Timeout
- *Menambahkan sinyal cancel secara otomatis* dengan pengaturan timeout
- Sinyal tidak dieksekusi cancel secara manual, tapi pake timeout
- Studi kasus: query ke database atau HTTP API, menentukan batas maksimal timeout nya
- *context.WithTimeout(parent, duration)* = Cancel signal otomatis menggunakan timeout
- Sekalipun bisa timeout, function ini membalikan variable cancel, kalo proses yang dibutuhkan lebih cepat. Maka di cancel saja sebelum timeout

## Context With Deadline
- *Pengaturan ditentukan waktu timeout untuk menambahkan sinyal*
- Misal jam 12 siang hari ini
- *context.WithDeadline(parent, time)* = deadline