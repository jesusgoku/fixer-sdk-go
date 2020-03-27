# Fixer SDK in Golang

[Fixer](https://fixer.io/) is a simple and lightweight API for current and historical foreign exchange (forex) rates.

Get a free API key from [here](https://fixer.io/product) with 1000 call/month.

## Example

```golang
package main

import (
	"fmt"
	"github.com/jesusgoku/fixer-sdk-go/pkg/fixer"
)

func main() {
	fixer := fixer.NewClient("YOUR_FIXER_API_HERE", nil)
	res, err := fixer.Latest()
	if err != nil {
		fmt.Printf("Error: %#v\n", err)
	}

	fmt.Printf("%#v\n", res)
}

```
