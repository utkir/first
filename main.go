package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*fmt.Println("s")
	f, err := os.Create("logs/log.txt")
	check(err)
	defer f.Close()

	n3, err := f.WriteString("writes\n")
	log.Print("ss")
	fmt.Printf("wrote %d bytes\n", n3)*/

	filename := "logs/log.txt"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString("sakljskjsdfkl")
	if err != nil {
		panic(err)
	}

}
