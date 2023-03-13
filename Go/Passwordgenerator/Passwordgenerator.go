package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

// DB Login Details
const (
	dbname = "Password_Generator"
	dbuser = "pg"
	dbpass = "E!j2zZ9!bVg4h3yd"
)

// Function that will add a new password (or not)
func DB(NewPwd string) string {
	// Make a connection to the DB
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname)

	// If an error occured, show it
	if err != nil {
		log.Fatal(err)
	}

	// Peform a GET from the DB to check if the new password already exists
	result, err := db.Query("SELECT Password FROM `passwords` WHERE Password = '" + NewPwd + "'; ")

	// If an error occured, show it
	if err != nil {
		log.Fatal(err)
	}

	// If the GET query returned a row (which means the password exists in DB), stop the function and return a MSG
	if result.Next() {
		fmt.Println("Password already exists in the Database")
		return ""
	}

	// Peform a INSERT to the DB, will add the new password to the DB
	insert, err := db.Query("INSERT INTO `passwords` (`id`, `Password`) VALUES (NULL, '" + NewPwd + "');")

	// If an error occured, show it
	if err != nil {
		log.Fatal(err)
	}

	// Close the connection to the row
	defer insert.Close()

	// Print out a msg that confirms that a new password has been added
	fmt.Println("Password has been added to the Database")

	// Return NULL
	return ""
}

// variable for the base string that will be used to generate a password
var BaseString string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Function to randomize the password string
func Randomize(length int) string {
	char := make([]byte, length)
	for i := range char {
		char[i] = BaseString[rand.Intn(len(BaseString))]
	}
	return string(char)
}

// Main
func main() {
	// Msg that informs the user of what there password is
	var msg string = "Your password is: "

	// Flag that defines the length of the password
	Length := flag.Int("L", 16, "How long should the password be? (Default value: 16)")

	// Flag that defines what should be included in the password
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

	// Print out the password
	NewPassword := Randomize(*Length)

	// Print a msg with the new password
	fmt.Println(msg, NewPassword)

	// Run the DB function with the new password
	DB(NewPassword)
}
