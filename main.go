// main.go

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	logLevel      = flag.String("level", "", "Log level to filter (INFO, ERROR, WARN, etc.)")
	customKeyword = flag.String("keyword", "", "Custom keyword to filter logs")
	logFilePath   = flag.String("file", "sample.log", "Path to the log file")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nCommands:")
		fmt.Println("  -level=INFO     : Filter logs by log level (e.g., INFO, ERROR, WARN)")
		fmt.Println("  -keyword=string : Filter logs by a custom keyword")
		fmt.Println("  -file=path      : Specify the path to the log file")
	}
}

func main() {
	flag.Parse()

	f, err := os.Open(*logFilePath)
	if err != nil {
		fmt.Printf("Error opening log file: %+v\n", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		// Check if the log entry contains the specified log level or custom keyword
		if (*logLevel == "" || strings.Contains(s, *logLevel)) &&
			(*customKeyword == "" || strings.Contains(s, *customKeyword)) {
			fmt.Print(s)
		}
	}
}
