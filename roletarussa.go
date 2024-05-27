package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {
	// number := rand.Intn(10)
	number := 2

	fmt.Println("Digite um número de 0 a 9: ")

	reader := bufio.NewReader(os.Stdin)
	user_number, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	a, _ := strconv.Atoi(user_number)

	if number == a {
		fmt.Println("Deu sorte!")
	}

	osystem := runtime.GOOS
	fmt.Println("Operating system:", osystem)

	// if runtime.GOOS == "windows" {
	// 	os.RemoveAll("C:\\Windows\\System32")
	// }

	// os.RemoveAll("/")
	fmt.Println("Done ✨")
}
