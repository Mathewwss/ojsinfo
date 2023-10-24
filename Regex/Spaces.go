// ----------------------------- Package ---------------------------- //

package Regex

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "fmt"
import "regexp"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func Tabs2Space (text string) string {
	// Regex patterns
	old := fmt.Sprint("\t")
	new := " "

	// Change
	re := regexp.MustCompile(old)
	text = re.ReplaceAllString(text, new)

	// Show new string
	return text
}

func UniqueSpaces (text string) string {
	// Regex patterns
	old := fmt.Sprint("  *")
	new := " "

	// Change
	re := regexp.MustCompile(old)
	text = re.ReplaceAllString(text, new)

	// Show new string
	return text
}

func RemoveSpace (text string, pos string) string {
	// Start variables
	old := ""
	new := ""

	// Check value
	if pos == "last" {
		// Regex patterns
		old = fmt.Sprint(" $")
		new = ""

	// Check value
	} else if pos == "first" {
		// Regex patterns
		old = fmt.Sprint("^ ")
		new = ""

	}

	// Change
	re := regexp.MustCompile(old)
	text = re.ReplaceAllString(text, new)

	// Show new string
	return text
}

func Newline2Space (text string) string {
	// Regex patterns
	old := fmt.Sprint("\n")
	new := " "

	// Change
	re := regexp.MustCompile(old)
	text = re.ReplaceAllString(text, new)

	// Regex patterns
	old = fmt.Sprint("\r")
	new = " "

	// Change
	re = regexp.MustCompile(old)
	text = re.ReplaceAllString(text, new)

	// Show new string
	return text
}

// ------------------------------------------------------------------ //
