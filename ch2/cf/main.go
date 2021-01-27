// cf converts it's numeric arguments from Celisus to Faherenheit
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pradeepnnv/gopl-training/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n", c, tempconv.CToF(c), f, tempconv.FToC(f))

	}
}
