// ----------------------------- Package ---------------------------- //

package Users

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

type User struct {
	UID uint32
	Email string
	Username string
	RealNames map[string]string
	PublicNames map[string]string
	Groups map[string]map[int]string
	URL string
	Phone string
	Country string
	Signature map[string]string
	Affiliation map[string]string
	Address string
	Languages []string
	Biography map[string]string
	ORCID string
}

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// ------------------------------------------------------------------ //
