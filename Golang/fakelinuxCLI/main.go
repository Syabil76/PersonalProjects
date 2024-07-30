package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 1 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "--h":
		printHelp()
	case "touch":
		makefile()
	case "cat":
		readfile()
	case "mousepad":
		writefile()
	case "rm":
		removefile()
	case "test":
		test()
	default:
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	filepath := "help.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error!", err)
		return
	}
	fmt.Println(string(data))
}

func readfile() {
	if len(os.Args) < 3 {
		fmt.Println("missing file")
		os.Exit(0)
	}
	file := os.Args[2]
	data, _ := os.ReadFile(file)
	if !validateFile(file) {
		panic("file not found")
	}
	fmt.Println(string(data))

}

func writefile() {

	if len(os.Args) < 3 {
		fmt.Println("missing file")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}

	exepath := "C:\\Windows\\system32\\notepad.exe"
	cmd := exec.Command(exepath, file)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Start failed: %s", err)
	}
	fmt.Printf("Waiting for command to finish.\n")
	err = cmd.Wait()
	fmt.Printf("Command finished with error: %v\n", err)

}

func makefile() {
	filename := os.Args[2]

	if len(os.Args) < 3 {
		fmt.Println("missing name")
		os.Exit(0)
	}

	f, err := os.Create(filename)
	fmt.Println("WARNING!!! file permission set to 0666 (+rw for all), may be insecure")
	if err != nil {
		fmt.Println("couldn't make file")
		os.Exit(0)
	}

	exepath := "C:\\Windows\\system32\\notepad.exe"
	cmd := exec.Command(exepath, filename)
	cmd.Start()

	defer f.Close()
	fmt.Println(f.Name())
}

func removefile() {
	file := os.Args[2]

	if len(os.Args) < 3 {
		fmt.Println("ERROR: missing file name")
		os.Exit(2)
	}

	e := os.Remove(file)
	if e != nil {
		fmt.Println("couldn't delete file")
		os.Exit(0)
	}

}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func test() {
	//insert test functions here
}
