// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "database/sql"
import "errors"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func New (identity int) (Submission, error) {
	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "submission_id"
	query = query + " " + "FROM"
	query = query + " " + "submissions"
	query = query + " " + "WHERE"
	query = query + " " + "submission_id = '" + fmt.Sprint(identity)
	query = query + "'"
	query = query + ";"

	// Database conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect databae
	db, err := sql.Open(driver, con)

	// Check errors
	if err != nil {
		// Stop
		return Submission{}, err

	}

	// Run query
	res, err := db.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return Submission{}, err

	}

	// Start variables
	s := Submission{}
	s.ID = -1

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&s.ID)

		// Check errors
		if err != nil {
			// Stop
			return Submission{}, err

		}
	}

	// Check submission ID
	if s.ID == -1 {
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
