package main

import (
	"fmt"
	"log"
)

func main() {
	var data string
	fmt.Print("Введите произвольные данные: ")
	_, err := fmt.Scan(&data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели: %s\n", data)
}
