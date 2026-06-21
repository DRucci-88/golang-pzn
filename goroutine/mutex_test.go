package goroutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0

	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Nilai X : ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(6 * time.Second)
	fmt.Println("Final Balance : ", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {

	// Simulasi Transfer duit
	// User1 -> Transfer Bank -> User2
	user1.Lock()
	fmt.Println("Lock User1", user1.Name)
	user1.Change(-amount) // Kita Negatif kan

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Le",
		Balance: 1_000_000, // Duit awal
	}
	user2 := UserBalance{
		Name:    "Rucco",
		Balance: 1_000_000, // Duit awal
	}
	fmt.Println("User 1", user1.Name, " Balance ", user1.Balance)
	fmt.Println("User 2", user2.Name, " Balance ", user2.Balance)

	// User1 TF 500k ke User2
	go Transfer(&user1, &user2, 100_000)

	// Ternyata oh ternyata User2 juga TF ke User1
	go Transfer(&user2, &user1, 200_000)

	time.Sleep(5 * time.Second)

	fmt.Println("Setelah proses transfer selesai")
	fmt.Println("User 1", user1.Name, " Balance ", user1.Balance)
	fmt.Println("User 2", user2.Name, " Balance ", user2.Balance)

	/*
		=== RUN   TestDeadLock
		User 1 Le  Balance  1000000
		User 2 Rucco  Balance  1000000
		Lock User1 Rucco
		Lock User1 Le
		Setelah proses transfer selesai
		User 1 Le  Balance  900000
		User 2 Rucco  Balance  800000

		Mengapa demikian ?
		Perhatikan kenapa tidak ada logger "Lock User1 ..." ?
		NAH Ini dia Deadlock nya.

		Ketika Transfer pertama, User 1 ("Rucco") nya di Lock.
		Di saat bersamaan Transfer kedua juga running dan Lock User 1 ("Le") nya.

		Transfer pertama sudah Change Amount User 1 ("Rucco"),
		terus ingin Lock User 2 ("Le") nya.
		Tapi tapi tapiiii
		Si ("Le") ini sedang di Lock pada proses Transfer kedua.

		Hal ini berlaku juga untuk hal yang sama bagi PoV nya Transfer Kedua
		Rangukaman:
		Transfer pertama User 2 ("Le") = Transfer kedua User 1 ("Le")
		Transfer pertama User 1 ("Rucco") = Transfer kedua User 2 ("Rucco")

		NNAAHHH ini lah yang dinamakan "DEADLOCK"
		Karena proses Lock dan Unlock nya ini tidak kelar-kelar, sudah keburu timeout deh.
		Itulah mengapa Balance masing-masing setelah selesai, hanya mengurangi amount yang ingin di transfer
	*/
}
