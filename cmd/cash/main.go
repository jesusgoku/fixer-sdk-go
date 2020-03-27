package main

import (
	"flag"
	"fmt"
	"github.com/jesusgoku/fixer-sdk-go/pkg/fixer"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
	"strconv"
)

func main() {
	locale := flag.String("locale", "en", "locale for output")

	flag.Parse()

	args := flag.Args()
	if len(args) < 3 {
		fmt.Printf("Usage: cash 3600 USD CLP\n")
		os.Exit(1)
	}

	amount, _ := strconv.ParseFloat(args[0], 32)
	sourceAmount := float32(amount)
	source := args[1]
	target := args[2]

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		fmt.Printf("Error: environment 'API_KEY' not found\n")
		os.Exit(1)
	}

	fixer := fixer.NewClient(apiKey, nil)
	res, err := fixer.Latest()
	if err != nil {
		fmt.Printf("Error: %#v\n", err)
		os.Exit(1)
	}

	if res.Error.Code != 0 {
		fmt.Printf("Error: %s\n\n%s\n", res.Error.Type, res.Error.Info)
		os.Exit(1)
	}

	sourceFactor, ok := res.Rates[source]
	if !ok {
		fmt.Printf("Error: source currency (%s) not found\n", source)
		os.Exit(1)
	}

	targetFactor, ok := res.Rates[target]
	if !ok {
		fmt.Printf("Error: target currency (%s) not found\n", target)
		os.Exit(1)
	}

	targetAmount := (sourceAmount / sourceFactor) * targetFactor

	printer := message.NewPrinter(language.MustParse(*locale))
	printer.Printf("Convert: %.2f %s -> %.2f %s\n", sourceAmount, source, targetAmount, target)
}
