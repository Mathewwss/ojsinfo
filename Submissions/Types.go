
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
	Stage string
	// language: title
	Section map[string]string
	// Order: langugage: [name, email]
	Authors map[int]map[string][]string
	// language: title
	Titles map[string]string
	// language: journal name
	JournalNames map[string]string
	Published bool
	PublicationYear int64
	PublicationVolume int64
	PublicationNumber string
	// language: [keywords]
	Keywords map[string][]string
	// language: abstract
	Abstract map[string]string
	// Round: [reviewers]
	// [[round, email, decision]]
	Reviewers [][]string
	Access map[string]int64
}

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// ------------------------------------------------------------------ //
