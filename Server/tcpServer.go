package main

// Library yang digunakan
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Mengambil input dari terminal
	arguments := os.Args

	/*
		Bila hanya ada 1 argumen, batalkan karena
		argumen pertama adalah argumen untuk mengeksekusi program
	*/
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	// Menyimpan isi argumen kedua ke variabel PORT
	PORT := ":" + arguments[1]

	// Membuat server TCP
	l, err := net.Listen("tcp", PORT)

	// Mencetak pesan error dan membatalkan program
	if err != nil {
		fmt.Println(err)
		return
	}

	// Menutup koneksi bila program selesai dieksekusi
	defer l.Close()

	// Menerima koneksi client
	c, err := l.Accept()

	// Mencetak pesan error dan membatalkan program
	if err != nil {
		fmt.Println(err)
		return
	}

	// Endless loop
	for {

		// Menerima data dari client
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		// Menghentikan server bila client mengirim pesan 'STOP'
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		// Mengirim balasan ke client
		fmt.Print("Client message -> ", string(netData))

		// Mencetak waktu saat ini
		t := time.Now()
		myTime := t.Format(time.RFC1123) + "\n"
		c.Write([]byte(myTime))
	}
}
