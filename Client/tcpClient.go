package main

// Library yang digunakan
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Mengambil input dari terminal
	arguments := os.Args

	/*
		Bila hanya ada 1 argumen, batalkan karena
		argumen pertama adalah argumen untuk mengeksekusi program
	*/
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port")
		return
	}

	// Menyimpan isi argumen kedua ke variabel CONNECT
	CONNECT := arguments[1]

	// Koneksi ke TCP server
	c, err := net.Dial("tcp", CONNECT)

	// Mencetak pesan error dan membatalkan program
	if err != nil {
		fmt.Println(err)
		return
	}

	// Endless loop
	for {
		// Membaca input dari client
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")

		/*
			Menyimpan input ke dalam variabel text
			_ merupakan blank identifier / dummy variable
		*/
		text, _ := reader.ReadString('\n')

		// Mengirim input dari client ke server
		fmt.Fprintf(c, text)

		// Menerima respon server
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)

		// Ketik 'STOP' untuk mengakhiri koneksi
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting. . .")
			return
		}
	}
}
