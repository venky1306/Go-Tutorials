// Example 1
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	defer fmt.Println("Good Night!")
// 	defer fmt.Println("Bye.")
// 	fmt.Println("Hello, Venkat")
// }

// Example 2
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := write("rd.txt", "This is readme file"); err != nil {
		log.Fatal("failed to write in readme.txt", err)
	}

	if err := fileCopy("rd.txt", "new-readme.txt"); err != nil {
		log.Fatal("failed to copy from readme.txt to new-readme.txt. ", err)
	}
}

func write(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}
	// file.Close()
	return file.Close()
}

func fileCopy(file1, file2 string) error {
	f1, err := os.Open(file1)
	if err != nil {
		return err
	}
	defer f1.Close()

	f2, err := os.Create(file2)
	if err != nil {
		return err
	}
	defer f2.Close()

	n, err := io.Copy(f2, f1)
	if err != nil {
		return err
	}
	fmt.Printf("Copied %d bytes from %s to %s\n", n, file1, file2)
	if err = f1.Close(); err != nil {
		return err
	}

	return f2.Close()

}
