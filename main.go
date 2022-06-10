package main

import (
	"fmt"
	"github.com/melbahja/goph"
	"log"
)

func main() {

	// Start new ssh connection with private key.
	auth, err := goph.Key("/home/mohamed/.ssh/id_rsa", "") // Diretory ssh key
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", "192.1.1.3", auth)
	if err != nil {
		log.Fatal(err)
	}

	// Closing connection
	defer client.Close()

	// Exec command
	out, err := client.Run("ls /tmp/")

	if err != nil {
		log.Fatal(err)
	}

	// Get output
	fmt.Println(string(out))
}
