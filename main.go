package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:embed data/KnownWords.json
var knownWords embed.FS

//go:embed data/Top1000.json
var challenge embed.FS

func main() {
	var helpFlag bool
	var challengeFlag bool

	flag.BoolVar(&helpFlag, "help", false, "Display help information")
	flag.BoolVar(&helpFlag, "h", false, "Display help information")
	flag.BoolVar(&challengeFlag, "challenge", false, "Challenge mode with over 1000 most common words")
	flag.BoolVar(&challengeFlag, "c", false, "Challenge mode with over 1000 most common words")

	flag.Parse()

	if len(flag.Args()) > 1 {
		fmt.Println("Can't have more than one argument. Exiting...")
		os.Exit(1)
	}

	if helpFlag {
		fmt.Println("Usage: german [--challenge | -c] | [--help | -h] | [no arguments]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

	if !challengeFlag {
		fmt.Println("Would you like to take on the Challenge Mode (Y/N)? (Default: N)")
		scanner.Scan()

		challengeFlag = strings.ToLower(scanner.Text()) == "y"
	}

	fmt.Println("How many would you like to try out? (default: 10)")

	scanner.Scan()

	numWords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Print("Defaulting to 10\n\n")
		numWords = 10
	}

	correct := 0

	var file []byte

	if challengeFlag {
		file, err = challenge.ReadFile("data/Top1000.json")
	} else {
		file, err = knownWords.ReadFile("data/KnownWords.json")
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	words, err := UnmarshalWords(file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	PlayQuiz(words, scanner, numWords, &correct)

	fmt.Println(string(colorGreen), "\nYou Got: "+strconv.Itoa(correct)+" Correct "+"Out of "+strconv.Itoa(numWords), string(colorReset))
}
