package main

import (
	"fmt"
	"strings"
	"time"
)

func DayPart(cTime string) string {
	// Welkom Message
	var Msg string = "[groet]! Welkom bij Fonteyn Vakantieparken"

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

	//Print uit the message
	fmt.Println(DayPart(Current_Time))
}
