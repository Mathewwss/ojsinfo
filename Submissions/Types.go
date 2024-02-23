
// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/Users"

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
	Reviews [][]Review
	Access map[string]int64
}

type Review struct {
	Users.User
	Recommendation string
	Round int64
	Unconsidered int64
	Cancelled int64
	Declined int64
	DateAssigned string
	DateNotified string
	DateConfirmed string
	DateCompleted string
	DateAcknowledged string
	DateDue string
	DateResponseDue string
}

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// ------------------------------------------------------------------ //
