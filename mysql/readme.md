# MySQL

## Package Database
- Berisikan kumpulan standard interface untuk berkomunikasi ke database
- Kode program jenis database apapun untuk kode yang sama

```txt
Aplikasi -> Database Interface -> Database Driver -> DBMS
```

```txt
https://github.com/go-sql-driver/mysql
```

## Database Driver
- Wajib menambah Database Driver
- Karena package database golang hanya memberikan interface

## Membuat Koneksi ke Database
- Ini hal pertama yang kita lakukan
- *sql.Open(driver, dataSourceName)* = Buat koneksi database dari object *sql.DB*
- driver = "mysql", registrasi driver nya "mysql"
- dataSourceName = format penulisan berbeda tiap database, lokasi dari address, port, username, password etc
- object *sql.DB* ditutup jika tidak digunakan lagi *function close()*

## Database Pooling
- a/ kumpulan koneksi ke Database / Management koneksi
- *sql.DB* bukanlah sebuah koneksi ke Database, tapi sebuah *pool ke Database* (Database Pooling)
- *sql.DB* golang melakukan management koneksi ke database otomatis.
- Kemampuan Database pooling golang, kita bisa menentukan jumlah minimal dan maksimal koneksi. Sehingga tidak membanjiri koneksi ke database

Pengaturen Database Pooling

| Method | Keterangan |
|-------|-------|
| (DB) SexMaxIdleConns(number) | Jumlah koneksi minimal | 
| (DB) SetMaxOpenConns(number) | Jumlah koneksi maksimal | 
| (DB) SetConnMaxIdleTime(duration) | Lama koneksi yang sudah digunakan akan dihapus | 
| (DB) SetConnMaxLifetime(duration) | Lama koneksi boleh digunakan |

## Eksekusi Perintah SQL
- *(DB) ExecContext(context,sql,params)* = Mengirim perintah SQL
- *Mengirim perintah SQL wajib mengirimkan Context*
- Jadi dengan Context ini, kita bisa membatalkan pengiriman perintah SQL nya

## Query SQL
- *(DB) QueryContext(context, sql, params)* = Query ke Database

## Rows
- Hasil dari *QueryContext* = structs *sql.Rows*
- Digunakan untuk iterasi terhadap hasil dari query
- *(Rows) Next() (boolean)* = iterasi hasil query, mirip seperti cursor query. Jika false = sudah tidak ada data lagi
- *(Rows) Scan(params...)* = Membaca tiap data
- *(Rows) Close()* = Wajib di tutup

## Tipe Data Column

| Tipe Data MySQL | Tipe Data Golang |
|-------|-------|
| VARCHAR, CHAR | string | 
| INT, BIGINT | int32, int64 | 
| FLOAT, DOUBLE | float32, float64 | 
| BOOLEAN | bool | 
| DATE, DATETIME, TIME, TIMESTAMP | time.Time | 

## Error Tipe Data Date
```txt
=== RUN TestQuerySql --- FAIL: TestQuerySql (0.01s) panic: Error 1193 (HY000): Unknown system variable 'parseDate' [recovered, repanicked]
```

- Default driver MySQL untuk Golang akan query tipe data DATE, DATETIME, TIMESTAMP menjadi []byte / []uint8.
- Dimana ini bisa dikonversi menjadi String lalu di parsing
- Tambah parameter *parseTime=true*

```go
db, err := sql.Open(
    "mysql",
    "root:root123@tcp(localhost:3306)/golang_pzn?parseTime=true",
)
```

## Nullable Type
- Golang tidak paham dengan tipe data NULL di database
- Tidak berubah menjadi nil, tapi bakal error ketika di *(Rows) Scan(params...)*

## Error Data Null
```txt
=== RUN   TestQuerySql
--- FAIL: TestQuerySql (0.02s)
panic: sql: Scan error on column index 2, name "email": converting NULL to string is unsupported [recovered, repanicked]
```

- Konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang
- Kita perlu memakai tipe data dari package sql

| Tipe Data Golang | Tipe Data Nullable |
|-------|-------|
| string |  database/sql.NullString  | 
| bool | database/sql.NullBool | 
| float64 | database/sql.NullFloat64 | 
| int32 | database/sql.NullInt32 | 
| int64 | database/sql.NullInt64 | 
| time.Time | database/sql.NullTime | 


## SQL Injection
- a/ teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam lapisan database
- Mengirim input dari user dengan perintah yang salah

Solusi ?
- Jangan membuat query SQL secara manual dengan menggabungkan String raw

## SQL dengan Parameter
- Perintah SQL tergantung dari input user
- Lalu membuat script query sql dari input SQL
- Bahaya SQL Injection ketika menggabungkan string raw
- Kita bisa memanfaatkan parameter ke 3 dalam function Exec dan Query

Contoh:
- SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
- INSERT INTO user(username, password) VALUES (?,?)

## Auto Increment
- *(Result) LastInsertId()* = Mengambil data id terakhir yang dibuat secara auto increment. Mencegah 2x query untuk SELECT lagi
- setelah Insert ke dalam MySQL secara otomatis 

## Query atau Exec dengan Parameter
- Proses Query dan Exec menggunakan parameter, sebenarnya sudah implementasi *Preprare Statement*
- Studi Kasus: ingin melakukan beberapa hal sekaligus, cuma beda parameternya. Cth Insert langsung banyak
- Pembuatan *Prepare Statement* bisa dilakukan manual tanpa menggunakan Query atau Exec dengan parameter
- Saat looping Query / Exec, connection pool nya bisa beda

## Prepare Statement
- Otomatis mengenali koneksi database yang digunakan
- Ketika menggunakan *Prepare Statement* berkali-kali, maka menggunakan *koneksi yang sama dan lebih efisien* karena hanya buat sekali diawal saja
- *(DB) Prepare(context,sql)* = Membuat prepare statement
- Terdapat di dalam *database/sql.Stmt* package
- Harus di *Close()*

## Database Transaction di Golang
- Perintah SQL yang dikirim Golang, *otomatis AUTO COMMIT*
- Bisa disable Auto Commit nya
- *(DB) Begin()* = struct Tx representasi Transaction, menggantikan DB untuk transaksi
- *(Tx) Commit()* = Commit atau Rollback

## Repository Pattern
- Repository is a mechanism for encapsulating storage, retieval and search behavior, which emulates a collection of objects
- Pattern Repository = Jembatan business logic aplikasi dengan perintah SQL ke Database

```txt
Business Logic call Repository 
Repository use Entity/Model
Repository impl Repository Implementation
Repository Implementation call Database
```

## Entity / Model
- Reprensentasi data dalam bentuk Struct
- Repository -> dibungkus dalam bentuk Struct. Proses Execute SQL menggunakan struct tersebut