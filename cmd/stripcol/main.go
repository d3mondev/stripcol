package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/d3mondev/stripcol"
	"github.com/spf13/pflag"
)

func main() {
	helpFlag := pflag.BoolP("help", "h", false, "Print help message")
	pflag.Parse()

	if *helpFlag == true {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Reads from stdin and strips ANSI color codes from the input\n")
		fmt.Fprintf(os.Stderr, "Example:\n    echo -e '\\033[31mHello\\033[0m World' | %s \n", os.Args[0])
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(stripcol.StripColor(text))
	}
}
