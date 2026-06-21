package goroutine_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Le Rucco"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Le Rucco Give Me Response"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Le Rucco Channel"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	fmt.Println("Capacity Channel ", cap(channel))
	fmt.Println("Jumlah data ", len(channel))

	channel <- "Le"
	channel <- "Rucco"
	channel <- "Blazzing"

	fmt.Println("Capacity Channel ", cap(channel))
	fmt.Println("Jumlah data ", len(channel))

	fmt.Println("Channel: ", <-channel) // 1
	fmt.Println("Capacity Channel ", cap(channel))
	fmt.Println("Jumlah data ", len(channel))
	fmt.Println("Channel: ", <-channel) // 2
	fmt.Println("Capacity Channel ", cap(channel))
	fmt.Println("Jumlah data ", len(channel))
	fmt.Println("Channel: ", <-channel) // 3
	fmt.Println("Capacity Channel ", cap(channel))
	fmt.Println("Jumlah data ", len(channel))

	fmt.Println("Selesai")
}

func TestBufferedChannelGoroutine(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Le"
		channel <- "Rucco"
		channel <- "Blazzing"
	}()

	go func() {
		fmt.Println("Channel: ", <-channel)
		fmt.Println("Channel: ", <-channel)
		fmt.Println("Channel: ", <-channel)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel1 := make(chan string)
	fmt.Println("Capacity Channel 1 ", cap(channel1))
	fmt.Println("Jumlah data 1 ", len(channel1))

	channel2 := make(chan string)
	fmt.Println("Capacity Channel 2 ", cap(channel2))
	fmt.Println("Jumlah data 2 ", len(channel2))

	fmt.Println("channel 1 IN")
	go func() {
		for i := 0; i < 100; i++ {
			channel1 <- "Index " + strconv.Itoa(i)
		}
		close(channel1)
	}()

	fmt.Println("channel 2 IN")
	go func() {
		for i := 0; i < 100; i++ {
			channel2 <- "Index " + strconv.Itoa(i)
		}
		close(channel2)
	}()

	fmt.Println("channel 1 OUT")
	go func() {
		for data := range channel1 {
			fmt.Println("Channel 1 :", data)
		}
	}()

	fmt.Println("channel 2 OUT")
	go func() {
		for data := range channel2 {
			fmt.Println("Channel [2] :", data)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// /// Cuma ambil satu saja yang paling cepat di ambil
	// select {
	// case data := <-channel1:
	// 	fmt.Println("Channel 1 ", data)
	// case data := <-channel2:
	// 	fmt.Println("Channel 2 ", data)
	// }
	// select {
	// case data := <-channel1:
	// 	fmt.Println("Channel 1 ", data)
	// case data := <-channel2:
	// 	fmt.Println("Channel 2 ", data)
	// }

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel 2 ", data)
			counter++

		}
		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Channel 2 ", data)
			counter++
		default:
			fmt.Println("Meunggu data")
		}
		if counter == 2 {
			break
		}
	}

	fmt.Println("Selesai")
}
