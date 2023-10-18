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

func Html2Text (html string) string {
	// Copy string
	text := html

	// Regex patterns
	patterns := [][]string{
		[]string{fmt.Sprint("<"), "\n<"},
		[]string{">", ">\n"},
		[]string{"<a.*mailto:", "Email: "},
		[]string{"<a.*href:", "URL: "},
		[]string{"<img src.*>", "IMG: "},
		[]string{"<.*>", ""},
		[]string{"</.*>", ""},
		[]string{string([]byte{92, 92, 34, 62}), ""},
		[]string{"\n", " "},
		[]string{"&nbsp;", " "},
		[]string{"&lt;", "<"},
		[]string{"&gt;", ">"},
		[]string{"&amp;", "&"},
		[]string{"  *", " "},
		[]string{"^ ", ""},
		[]string{" $", ""},
	}

	// View patterns
	for a := 0; a < len(patterns); a++ {
		// Change string
		old := patterns[a][0]
		new := patterns[a][1]
		re := regexp.MustCompile(old)
		text = re.ReplaceAllString(text, new)

	}

	// Show new string
	return text
}

// ------------------------------------------------------------------ //
