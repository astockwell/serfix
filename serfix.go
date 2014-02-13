package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime"
	// "strconv"
	// "strings"
)

const (
	helpFlagUsage  = "Help and usage instructions"
	forceFlagUsage = "Force overwrite of destination file if it exists"
)

var helpPtr = flag.Bool("help", false, helpFlagUsage)
var forcePtr = flag.Bool("force", false, forceFlagUsage)
var counter int = 0
var lexer = regexp.MustCompile(`s:\d+:\\?\".*?\\?\";`)
var re = regexp.MustCompile(`(s:)(\d+)(:\\?\")(.*?)(\\?\";)`)
var esc = regexp.MustCompile(`(\\"|\\'|\\\\|\\a|\\b|\\n|\\r|\\s|\\t|\\v)`)

// var escstrs = []string{`\"`, `\'`, `\\`, `\a`, `\b`, `\n`, `\r`, `\s`, `\t`, `\v`}

func init() {
	// Short flags too
	flag.BoolVar(helpPtr, "h", false, helpFlagUsage)
	flag.BoolVar(forcePtr, "f", false, forceFlagUsage)
}

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Handle flags
	flag.Parse()

	args := flag.Args()

	if *helpPtr {
		printUsage()
		return
	}

	if len(args) > 0 {
		filename := fmt.Sprintf("%s", args[0])
		// Open provided file
		infile, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		// close provided file on exit and check for its returned error
		defer infile.Close()

		newDestination := false
		outfilename := filename
		tempfilename := fmt.Sprintf("%s~", outfilename)
		if len(args) > 1 {
			newDestination = true
			outfilename = fmt.Sprintf("%s", args[1])
			tempfilename = fmt.Sprintf("%s~", outfilename)
			if !*forcePtr {
				if _, err := os.Stat(outfilename); err == nil {
					fmt.Println("Destination file already exists, aborting serfix.")
					return
				}
			}
		}

		// Open out file
		tempfile, err := os.Create(tempfilename)
		if err != nil {
			println(err)
		}
		// close out file
		defer tempfile.Close()

		r := bufio.NewReaderSize(infile, 2*1024*1024)

		line, err := r.ReadString('\n')
		for err == nil {
			tempfile.WriteString(lexer.ReplaceAllStringFunc(string(line), replace))

			line, err = r.ReadString('\n')
		}
		// if isPrefix {
		// 	fmt.Println(errors.New("buffer size too small"))
		// 	return
		// }
		if err != io.EOF {
			fmt.Println(err)
			return
		}

		// Close the in/out files
		if err := tempfile.Close(); err != nil {
			fmt.Println(err)
			return
		}
		if err := infile.Close(); err != nil {
			fmt.Println(err)
			return
		}

		if !newDestination {
			// Remove original file
			if err := os.Remove(filename); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			// If destination exists and force flag is used, remove destination file
			if _, err := os.Stat(outfilename); err == nil {
				if !*forcePtr {
					if err := os.Remove(outfilename); err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		}
		if err := os.Rename(tempfilename, outfilename); err != nil {
			fmt.Println(err)
			return
		}

	} else {
		r := bufio.NewReaderSize(os.Stdin, 2*1024*1024)

		line, isPrefix, err := r.ReadLine()
		for err == nil && !isPrefix {
			fmt.Println(lexer.ReplaceAllStringFunc(string(line), replace))

			line, isPrefix, err = r.ReadLine()
		}
		if isPrefix {
			fmt.Println(errors.New("buffer size too small"))
			return
		}
		if err != io.EOF {
			fmt.Println(err)
			return
		}
	}
}

func replace(matches string) string {
	parts := re.FindStringSubmatch(matches)

	str_len := len(parts[4]) - len(esc.FindAllString(parts[4], -1))
	// esc_len := 0
	// for _, escstr := range escstrs {
	// 	esc_len = esc_len + strings.Count(parts[4], escstr)
	// }

	return fmt.Sprintf("%s%d%s%s%s", parts[1], str_len, parts[3], parts[4], parts[5])
}

func printUsage() {
	fmt.Println("Usage: serfix [flags] filename [outfilename]")
	fmt.Println("Alt. Usage: cat filename | serfix")
	fmt.Println("")
	fmt.Println("\t -f, --force \t\t\t Force overwrite of destination file if it exists.")
	fmt.Println("\t -h, --help  \t\t\t Print serfix help.")
	fmt.Println("")
}
