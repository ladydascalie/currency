# Accounting

The accounting package is used when dealing with currency math that requires a high degree of precision.

A few business rules and assumptions are made:

- Banker's Rounding (or round half to even) is applied whenever rounding mode is available.
  - see: [https://en.wikipedia.org/wiki/Rounding#Round_half_to_even](https://en.wikipedia.org/wiki/Rounding#Round_half_to_even)
- The maximum precision allowed after the decimal dot is **2**.

Examples:

```go
package main

import (
        "fmt"

        "github.com/ladydascalie/currency/accounting"
        "github.com/ladydascalie/currency"
)

func main() {
        // Create an amount from an int64 and a currency
        amount := accounting.MakeAmount(currency.GBP, 1234)
        
        // Amounts contain the minor representation of that currency. 
        fmt.Println(amount.MinorValue)
        // output: 1234
 
        // alternatively floats are also ok!
        gbp := accounting.Float64ToAmount(currency.GBP, 32.32)
        fmt.Printf("minor value: %d, stringer: %s", gbp.MinorValue, gbp)
        // output: minor value: 3232, stringer: 32.32 GBP

        // Currencies with no minor decimals drop the invalid precision.
        jpy := accounting.Float64ToAmount(currency.JPY, 32.32)
        fmt.Printf("minor value: %d, stringer: %s", jpy.MinorValue, jpy)
        // output: minor value: 32, stringer: 32 JPY

        // Exchanging currencies is also supported.
        usd := accounting.Float64ToAmount(currency.USD, 100.0)
        eur, err := accounting.Exchange(usd, currency.EUR, 1.08968)
        if err != nil {
                // handle...
        }
        fmt.Println(eur)
        // output: 91.77 EUR
}
```
