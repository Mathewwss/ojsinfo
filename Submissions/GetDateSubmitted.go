// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *Submission) GetDateSubmitted () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "CASE"
	query = query + " " + "WHEN date_submitted IS NULL THEN \"\""
	query = query + " " + "ELSE date_submitted"
	query = query + " " + "END"
	query = query + " " + "FROM"
	query = query + " " + "submissions"
	query = query + " " + "WHERE"
	query = query + " " + "submission_id = '" + fmt.Sprint(s.ID) + "'"
	query = query + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&s.Start)

		// Check errors
		if err != nil {
			// Stop
			return err

		}
	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
