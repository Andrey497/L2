package pkg

import (
	"log"
	"os"
)

func WriteDatInfile(nameFile string, data []byte) {
	filePtr, err := os.OpenFile("../"+nameFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
	_, err = filePtr.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
