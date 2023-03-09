package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var BaseString string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Randomize(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = BaseString[rand.Intn(len(BaseString))]
	}
	return string(b)
}

func main() {
	// Msg that informs the user of what there password is
	var msg string = "Your password is: "

	//Flag that defines the length of the password
	Length := flag.Int("L", 16, "How long should the password be? (Default value: 16)")

	//Flag that defines what should be included in the password
	Numbers := flag.Bool("n", false, "Should numbers be used? (Default value: False)")
	Logograms := flag.Bool("l", false, "Should logograms be used? (Default value: False)")
	Punctuations := flag.Bool("p", false, "Should punctuation be used? (Default value: False)")
	Quotations := flag.Bool("q", false, "Should quotations be used? (Default value: False)")
	Dashes_Slashes := flag.Bool("ds", false, "Should dashes & slashes be used? (Default value: False)")
	Math_Symbols := flag.Bool("m", false, "Should math symbols be used? (Default value: False)")
	Brackets := flag.Bool("b", false, "Should brackets be used? (Default value: False)")

	// Parse the above flags
	flag.Parse()

	// Flag for Numbers, if True add all the Numbers to the basestring
	if *Numbers {
		BaseString = BaseString + "1234567890"
	}

	// Flag for Logograms, if True add all the Logograms to the basestring
	if *Logograms {
		BaseString = BaseString + "#$%&@^~" + "`"
	}

	// Flag for Punctuations, if True add all the Punctuation to the basestring
	if *Punctuations {
		BaseString = BaseString + ".,:;"
	}

	// Flag for Quotations, if True add all the Quotations to the basestring
	if *Quotations {
		BaseString = BaseString + `"'`
	}

	// Flag for Dashes & Slashes, if True add all the Dashes and Slashes to the basestring
	if *Dashes_Slashes {
		BaseString = BaseString + `\/|_-`
	}

	// Flag for Math Symbols, if True add all the Math Symbols to the basestring
	if *Math_Symbols {
		BaseString = BaseString + `<>*+!?=`
	}

	// Flag for Brackets, if True add all the Math Symbols to the basestring
	if *Brackets {
		BaseString = BaseString + "{[()]}"
	}

	//Print out the password
	fmt.Println(msg, Randomize(*Length))
}
