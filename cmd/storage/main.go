package main

import (
	"fmt"
	"log"

	"github.com/nukhvu/bot-bek/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("noha", []byte("noha-byte"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)

	fmt.Println(st.Files)

	fmt.Println("working:", restoredFile)
}
