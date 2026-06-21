# Goroutine

## Parallel Programming

a/ memecahkan suatu masalah dengan cara membaginya menjadi yang lebih kecil dan dijalankan bersamaan pada waktu yang bersamaan pula.

Contoh Parallel
- Menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, editor, browser)
- Beberapa kopi menyiapkan makanan di restoran dan menyiapkan menu berbeda
- Antrian di bank, dimana teller melayani nasabah

| Process | Thread |
|-------|-------|
| a/ eksekusi program | segmen dari process | 
| Mengkonsumsi memory besar | menggunakan memory kecil | 
| Terisolasi dengan process lain | Saling terhubung jika dalam proces yang sama | 
| Lama untuk dijalankan & dihentikan | Cepat untuk dijalankan dan dihentikan |

## Parallel vs Concurrency

| Process | Thread |
|-------|-------|
| Menjalankan pekerjaan secara bersamaan | Menjalankan pekerjaan secara bergantian |
| Butuh banyak Thread | Butuh sedikit Thread |

Contoh Concurrency
- Saat kita makan di cafe, bisa sekalian makan, ngobrol, minum, makan, ngobrol, minum, makan, vapor. Walaupun terasa bersamaan, tetap satu aksi yang dilakukan secara bersamaan. One at the time. Namun bisa berganti kapanpun yang kita mau.

## CPU-Bound

- Banyak algoritma dibuat yang hanya membutuhkan CPU. Algoritma jenis ini ketergantungan dengan kecepatan CPU
- Contoh Machine Learning. Oleh karena itu, sekarang banyak sekali teknologi Machine Learning banyak yang menggunakan GPU, karena punya Core yang lebih banyak.
- Algoritma jenis ini tidak ada benefitnya menggunakan Concurrency, namun bisa dibantu dengan implementasi Paraller Programming

## I/O-Bound

- a/ algoritma / aplikasi sangat terganntung dengan kecepatan input output devices yang digunakan.
- Contoh: membaca data dari file, database dll
- Mostyly of the time, kita banyak buat applikasi seperti ini
- Aplikasi jenis IO Bound, walaupun bisa terbantu dengan implementasi Parallel Programming, tapi benefitnya lebih baik jika menggunakan Concurrency Programming
- Bayangkan kita membaca data dari database, dan Thread harus menunggu 1 detik untuk mendapa balasan dari database, 1 detik itu jika menggunakan Concurrency Programming bisa digunakan untuk melakukan hal lain lagi.

## Gouroutine

Pengertian:
- a/ Thread ringan (sebenarnya dia running didalam thread) yang dikelola Go Runtime
- Ukuran sekitar 2Kb, jauh lebih kecil drpd Thread sekitar 1Mb
- Goroutine berjalan secara concurrent, sementara thread secara parallel
- Dijalankan secara concurent didalam sebuah thread

Cara kerja:
- Goroutine dijalankan oleh Go Scheduler dalam thread, dimana jumlah threadnya sebanyak GOMAXPROCS (sejumlah core CPU)
- Jadi goroutine bukan pengganti Thread, karena Goroutine berjalan diatas Thread
- Kemudahan bagi developer, kita tidak perlu management Thread secara manual. Semua sudah diatur Go Scheduler

## Channel

- a/ tempat komunikasi secara synchronous yang bisa dilakukan oleh goroutine
- Terdapat pengirim dan penerima, dan bisanya pengirim dan penerima adalah goroutin yang berbeda
- Saat melakukan pengiriman data ke Channel, goroutine akan ter-block, sampai ada yang menerima data tersebut
- Channel disebut sebagai alat komunikasi synchronous (blocking)
- Alternatif mekanisme seperti async await yang terdapat di beberapa bahasa pemrograman lain
- Tempat mengirim data antara goroutine pengirim dan goroutine penerima

### Karateristik Channel:
- Default channel hanya bisa menampung satu data. Jika ingin menambahkan data lagi, harus menunggu data yang ada di channel tsb diambil
- Hanya bisa menerima satu jenis data
- Bisa diambil dari lebih dari satu goroutine
- Harus di close jika tidak digunakan (memory leak)

### Membuat Channel:
- Direpresentasikan dengan tipe data *chan*
- Menggunakan *make()*
- Tentukan tipe data

### Komunikasi Channel
- Mengirim data: *channel <- data*
- Menerima data: *data <- channel*
- Menutup channel: *close()*

### Channel sebagai Parameter
- In real apps, seringnya kita akan mengirim channel ke function lain via parameter
- By default golang, parameter = pass by value, kita harus kirim pointer = pass by reference
- *Channel otomatis pass by reference*

### Channel In dan Out
- Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan meneriima data dari channel tersebut
- Kita bisa memberi ke function, bahwa channel tsb hanya digunakan untuk mengirim data atau menerima data saja
- Channel In = Mengirim Data *chan-<*
- Channel Out = Menerima data *<-chan*

### Buffered Channel

- Default channel hanya menerima 1 data
- Jika menambah data ke 2, maka kita akan diminta menunggu sampai data ke 1 ada yang mengambil
- Terkadang ada kasus dimana kecepatan pengirim > penerima, jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
- *Buffered Channel* = buffer yang bisa digunakan untuk menampung data antrian di Channel

### Buffer Capacity
- Penyimpanan di dalam channel
- Bebas memasukan jumlah kapasitas antrian di dalam buffer
- Jika set 5, maka bisa menerima 5 data di buffer
- Jika mengirim data ke 6, maka diminta menunggu sampai buffer ada yang kosong
- Cocok jika goroutine menerima data lembih lambat dari yang mengirim data
- Menunggu hanya ketika buffer sudah habis (capacity nya)

### Range Channel
- Channel dikirim data secara terus menerus oleh pengirim, 
- Tidak tahu kapan channel tsb berhenti menerima data atau mengirim data
- Tidak tahu jumlah channel menerima atau mengirim
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data dari channel
- Ketika channel di close(), maka perluangan berhenti
- Lebih sederhana drpd pengecekan manual secara manual

### Select Channel
- Ada kasus membuat banyak channel dan menjalankan goroutine nya
- Lalu ingin mendapatkan data dari semua channel tersebut
- *Select Channel* bisa mendapatkan semua data dari channel tsb
- *Select Channel* bisa memilih data tercepat dari banyak channel, jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random

### Default Select
- Apa yang terjadi jikta melakukan select terhadap channel yang ternyata tidak ada datanya ? *Deadlock*
- Maka kita akan menunggu sampai data ada
- Kadang kita ingin melakukan sesuatu jika semua channel tidak ada datanya ketika kita melakukan select channel
- Dalam select kita bisa menambahkan *default*, dimana akan di eksekusi jika memang semua channel yang di select tidak ada datanya.

## Race Conditions
Masalah dengan Goroutine:
- Saat menggunakan goroutine, tidak hanya berjalan secara concurrent, tapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
- Hal ini berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine secara bersamaan
- Masalah ini dinamakan *Race Condition*

## sync.Mutex (Mutual Exclusion)
- Mengatasi masalah *Race Condition*, golang punya struct bernama *sync.Mutex*
- *Mutex* digunakan untuk melakukan *locking* dan *unlocking*, dimana ketika melakukan locking terhadap mutex, maka tidak ada yang bisa melakukan locking sebelum unlocking.
- Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang diperbolehkan, setelah goroutine tersebut di unlock, baru goroutine yang lain bisa melakukan locking
- *Satu goroutine yang locking mutex pada satu waktu*
- Jika ada variable yang di sharing antar goroutine

## sync.RWMutex (Read Write Mutual Exclusion)
- Kasus ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga membaca data
- Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara membaca dan mengubah
- Di Golang telah disediakan struct RWMutex (Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis ini memiliki dua lock.
- *Lock untuk Read*
- *Lock untuk Write*

## Deadlock
- a/ keadaan dimana sebuah proses goroutine saling menungg lock sehingga tidak ada satupun gorouting yang bisa jalan
- Sering terjadi pada aplikasi parallel dan concurrent.

## sync.WaitGroup
- a/ fitur yang bisa dilakukan untuk menunggu sebuah proses selesai dilakukan.
- *Menunggu proses satu / banyak goroutine selesai*. Jangan pake time.Sleep karena gak tau selesai nya berapa lama si goroutine nya
- Misal kita ingin menjalankan beberapa proses menggunakan goroutine, tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai
- *Add(int)* = Ada proses goroutine, int = jumlah goroutine
- *Done()* = Setelah proses goroutine selesai, sama aja kayak *Add(-1)*
- *Wait()* = Menunggu semua proses selesai, sampe int == 0
- *Tidak harus goroutine, yang penting ada proses yang harus ditunggu*
- *Jangan lupa Done(), nanti counter int nya kagak pernah 0*

## sync.Once
- a/ Fitur untuk memastikan bahwa sebuah function di eksekusi hanya sekali
- Jadi berapa banyak pun goroutine yang mengakses, dipastikan bahwa goroutine pertama yang hanya bisa mengeksekusi functionnya


## sync.Pool
- a/ implementasi design pattern bernama *object pool pattern*
- Sederhananya digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari Pool dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
- Implementasi Pool di Golang ini sudah *aman dari problem race condition*
- Biasanya digunakan manage koneksi ke Database

## sync.Map
- Mirip Map biasa tapi *aman untuk menggunakan concurrent menggunakan goroutine*
- Beberapa function:
    - *Store(key, value)* = menyimpan data ke Map
    - *Load(key)* = mengambil data dari Map
    - *Delete(key)* = menghapus data dari map
    - *Range(function(key, value))* = Iterasi seluruh data dari Map

## sync.Cond
- a/ implementasi *locking berbasis kondisi*
- Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya, namun berbeda dengan Locker biasanya, di Cond terdapat *function Wait()* untuk menunggu berdasakan kondisi
- *Function Signal()* digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi.
- *Function Broadcast()* digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
- *Function sync.NewCond(Locker)* untuk membuat Cond

## sync/Atomic
- Atomic = menggunakan data primitive secara aman pada proses concurrent
- Contoh: kita menggunakan Mutex untuk melakukan Locking ketika ingin menaikan angka di counter. Hal ini bisa menggunakan Atomic

## time.Timer
- a/ *representasi satu kejadian*
- Kadang kita butuh delay jobs dan sebagainya, "Oke nanti kasih tau saja lima detik kemudian". Kurleb begitulah bahasanya
- Ketika waktu timer sudah expire, maka event akan dikirim kedalam channel
- *time.NewTimer(duration)*

## time.After()
- Kadang kita hanya butuh channel nya saja, tidak butuh data Timer nya

## time.AfterFunc()
- Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
- Kita tidak perlu pake channel nya, cukup kirim kan function yang akan dipanggil ketika Timer mengirim kejadian (event)

## time.Ticker
- a/ representasi kejadian yang berulang
- Ketika waktu ticker sudah expire, maka event akan dikirim kedalam channel
- *time.NewTicker(duration)* = Membuat Ticker
- *Ticker.Stop()* = Menghentikan Ticker

## time.Tick
- Kadang tidak butuh data Ticker, hanya butuh channel

## GOMAXPROCS
- Goroutine sebenarnya dijalankan didalam Thread
- Seberapa banyak Thread yang ada di golang ketika aplikasi kita berjalan ?
- GOMAXPROCS = function di package runtime untuk *mengubah atau mengambil jumlah thread*
- Default: jumlah thread di golang sebanyak jumlah CPU di komputer
- *function runtime.NumCpu()* = Melihat jumlah CPU kita