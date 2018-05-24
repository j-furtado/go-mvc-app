package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	db_host := os.Args[1]
	fmt.Println(db_host)
	log.Print(db_host)
}
