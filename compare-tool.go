package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	sortLines := flag.Bool("s", false, "Sort lines alphabetically")
	printNewlyAddedLines := flag.Bool("na", false, "Print only the newly added lines")
	printUniqueLines := flag.Bool("r", false, "Print only the unique lines")
	showHelp := flag.Bool("h", false, "Show available arguments")

	flag.Parse()

	if *showHelp {
		printHelp()
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Please provide at least one file path.")
		return
	}

	filePaths := flag.Args()
	lines := make([]string, 0)

	for _, filePath := range filePaths {
		fileLines := readLines(filePath)
		lines = append(lines, fileLines...)
	}

	if *sortLines {
		sort.Strings(lines)
	}

	if *printNewlyAddedLines {
		lines = filterNewlyAddedLines(lines)
	}

	if *printUniqueLines {
		lines = removeDuplicates(lines)
	}

	fmt.Println("Result:")
	for _, line := range lines {
		fmt.Println(line)
	}
}

// readLines reads a text file and returns its lines as a slice of strings
func readLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

// removeDuplicates removes duplicate entries from a slice of strings
func removeDuplicates(lines []string) []string {
	seenLines := make(map[string]bool)
	uniqueLines := make([]string, 0)

	for _, line := range lines {
		if !seenLines[line] {
			seenLines[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines
}

// filterNewlyAddedLines filters out the lines that are not newly added (duplicates)
func filterNewlyAddedLines(lines []string) []string {
	lineMap := make(map[string]int)
	newlyAddedLines := make([]string, 0)

	for i, line := range lines {
		if lineMap[line] == 0 {
			lineMap[line] = i + 1
		}
	}

	for _, line := range lines {
		if lineMap[line] > 0 {
			newlyAddedLines = append(newlyAddedLines, line)
			lineMap[line] = 0
		}
	}

	return newlyAddedLines
}

// printHelp displays the available arguments and their descriptions
func printHelp() {
	fmt.Println("Available arguments:")
	flag.PrintDefaults()
}
