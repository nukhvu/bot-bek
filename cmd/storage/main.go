package main

import (
	"fmt"

	"github.com/nukhvu/bot-bek/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("work", st)
}
