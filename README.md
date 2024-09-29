# currency

This package generates structs containing all the up-to-date `ISO4217` currency codes and minor units, along with a very simple validator.

Data is graciously provided by:

- [International Organization for Standardization](https://www.iso.org/iso-4217-currency-codes.html)
- [Currency Code Services – ISO 4217 Maintenance Agency](https://www.currency-iso.org)

## Usage:

```go
package main

import (
	"fmt"
	"github.com/ladydascalie/currency"
)

func main() {
	// Validate currency code.
	if !currency.Valid("ABC") {
		// Handle invalid currency code.
		fmt.Println("Invalid currency code")
	}

	// Retrieve and print currency values.
	fmt.Println(currency.USD.Code())      // Output: USD
	fmt.Println(currency.USD.MinorUnit()) // Output: 2
}
```
