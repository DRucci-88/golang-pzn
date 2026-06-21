package main

import "fmt"

func main() {
	fmt.Println("Hello World a")
	fmt.Println(number())
	fmt.Println(len("LeRucco"))
	fmt.Println("LeRucco"[0])

	// var name string
	name := ""

	name = "Le Rucco"
	fmt.Println(name)
	name = "Kamvret"
	fmt.Println(name)

	const name1 string = "Le Const"
	fmt.Println(name1)
	// name1 = "cannot assign to name1"

	var (
		firstName = "First Name"
		lastName  = "Last Name"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		firstNameConst = "First Name Const"
		lastNameConst  = "Last Name Const"
	)
	fmt.Println(firstNameConst)
	fmt.Println(lastNameConst)

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai8, nilai16, nilai32, nilai64)

	var name2 = "LeRucco"
	var e uint8 = name2[0]
	var eString = string(e)
	fmt.Println(eString)

	type NoKTP string // NoKTP menjadi tipe data baru, alias dari tipe data string
	var ktpLe string = "1111111111"
	fmt.Println(ktpLe)
	fmt.Println(NoKTP(fmt.Sprint(12345678)))
	fmt.Println(NoKTP(ktpLe))

	/// Array
	var names [3]string
	names[0] = "Le"
	names[1] = "Rucco"
	// names[2] = "John"
	fmt.Println(names)

	var ages = [3]int8{100, 90}
	fmt.Println(ages)

	var ages1 = [...]int8{100, 90}
	fmt.Println(ages1)
	fmt.Println(len(ages1))

	/*
		/// Slice
		- Potongan dari data Array
		- Mirip Array, tapi ukuran Slice bisa berubah
		- Mirip kayak List lah ya kalo di programming lain

		- Tipe data Slice punya 3 data:
			- pointer = penunjuk data pertama di array pada slice
			- length = panjang dari slice, length tidak boleh lebih dari capacity
			- capacity = kapasitas slice

		/// Membuat Slice dari Array
		- array[low:high]	= Dimulai dari index low sampai index sebelum high
		- array[low:] 		= Dimulai dari index low sampai index akhir di array
		- array[:high]		= Dimulai index 0 sampai index sebelum high
		- array[:]			= Dimulai index 0 sampai index akhir di array
	*/

	// names1 := [...]string{"Le", "Rucco", "Robert", "De", "Junior"}
	// namesSlice := names1[4:len(names1)]
}

func number() int8 {
	return 10
}
