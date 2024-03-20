// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "errors"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func New (identity int) (Submission, error) {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return Submission{}, err

	}

	// Sql query
	query := fmt.Sprintf(`
		SELECT DISTINCT
			submission_id, date_submitted, last_modified, locale
		FROM
			submissions
		WHERE
			submission_id = %v
		;
	`, identity)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return Submission{}, err

	}

	// Start variables
	s := Submission{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(
			&s.ID, &s.DateStart, &s.DateLastChange, &s.Locale,
		)

		// Check errors
		if err != nil {
			// Stop
			return Submission{}, err

		}
	}

	// Check submission ID
	if s.ID == 0 {
		// Create error
		msg := "[ERROR] -> Not found submission!"
		err = errors.New(msg)

		// Stop
		return Submission{}, err

	}

	// Finish
	return s, nil
}

// ------------------------------------------------------------------ //
