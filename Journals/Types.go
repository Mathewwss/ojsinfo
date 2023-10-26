// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

type Journal struct {
	ID int
	Names map[string]string
	Path string
	Groups map[int]map[string]string
	Sections map[int]map[string]string
	Reviewers []string
	Submissions []int
	SubmissionsIncomplete []int
	SubmissionsNew []int
	SubmissionsReview []int
	SubmissionsProduction []int
	SubmissionsCopyediting []int
	SubmissionsScheduled []int
	SubmissionsPublished []int
	SubmissionsDeclined []int
}

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// ------------------------------------------------------------------ //
