
// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

type Submission struct {
	ID int
	Locale string
	Start string
	// language: title
	Section map[string]string
	// language: title
	Titles map[string]string
	// language: journal name
	JournalNames map[string]string
	// language: [keywords]
	Keywords map[string][]string
}

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// ------------------------------------------------------------------ //
