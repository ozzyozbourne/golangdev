package main

import (
	"fmt"
	"log"
	"os"
)

var name []byte

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("Pls provide a input parameters to the main")
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		log.Fatalln("Error in reading the files in the folder -> " + args[0])
	}

	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			log.Fatalln("Error in getting file info")
		}
		if info.Size() == 0 {
			name = append(name, file.Name()...)
		}
		fmt.Println(file.Name())
	}
	name = append(name, "Evacuate the dance floors"...)
	er := os.WriteFile("dancefloor.txt", name, 0777)
	if er != nil {
		log.Fatalln("Error in writing to the file")
	}

	fmt.Print("\033[2J")
}
