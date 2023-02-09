package main

import (
	"fmt"
	"strings"
	"time"
)

func LPN_Check(LPN string) bool {

	// Create an array with pre-defined values
	var LPN_Array = []string{
		"ACB-34-1",
		"DCE-43-2",
		"FGH-56-3",
		"IJK-65-4",
		"LMN-78-5"}

	// For loop, for each value in LPN_Array check if they are equal to the LPN and if so return "true", if not return "false"
	for _, value := range LPN_Array {
		if value == LPN {
			return true
		}
	}
	return false
}

func DayPart(cTime string) string {
	// Welkom Message
	var Msg string = "[groet]! Welkom bij Fonteyn Vakantieparken."

	// Make an empty variable that will later be used to define the daypart
	var Greeting string = ""

	// If current time is greater than or equal to 07:00:00 and less than or equal to 12:00:00
	if "07:00:00" <= cTime && cTime <= "12:00:00" {
		Greeting = "Goedemorgen"
	}

	// If current time is greater than or equal to 12:00:00 and less than or equal to 18:00:00
	if "12:00:00" <= cTime && cTime <= "18:00:00" {
		Greeting = "Goedemiddag"
	}

	// If current time is greater than or equal to 18:00:00 and less than or equal to 23:00:00
	if "18:00:00" <= cTime && cTime <= "23:00:00" {
		Greeting = "Goedenavond"
	}

	// If current time is greater than or equal to 23:00:00 and less than or equal to 07:00:00
	if "23:00:00" <= cTime && cTime <= "07:00:00" {
		// Edit the message to the following
		Msg = "Sorry, de parkeerplaats is 's nachts gesloten"

		// Return the edited message
		return Msg
	}

	// Return the message edited with the correct Daypart greeting
	return strings.Replace(Msg, "[groet]", Greeting, 1)
}

func main() {
	// Get the current time in 24hr format
	var Current_Time string = time.Now().Format("15:04:05")

	// Make an empty variable that will later be used to define the license plate number
	var LPN string

	// Ask the user for a license plate number and save the input in the variable "LPN"
	fmt.Println("Vul hier het kenteken van uw auto in:")
	fmt.Scanln(&LPN)

	if LPN_Check(LPN) {
		//Give the current time to the DayPart function and print out what it returns
		fmt.Println(DayPart(Current_Time))
	}

	if !LPN_Check(LPN) {
		//Print out the following message
		fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
	}
}
