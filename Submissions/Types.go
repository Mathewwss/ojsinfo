
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
	Authors map[int]map[string]string
	// language: title
	Titles map[string]string
	// language: journal name
	JournalNames map[string]string
	Published bool
	PublicationYear int
	PublicationVolume int
	PublicationNumber string
	// language: [keywords]
	Keywords map[string][]string
	// language: abstract
	Abstract map[string]string
}

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// ------------------------------------------------------------------ //
