// ----------------------------- Package ---------------------------- //

package Journals

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

func (j *Journal) GetPublishedSubmissions () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := "SELECT DISTINCT"
	query = query + " " + "submission_id"
	query = query + " " + "FROM"
	query = query + " " + "submissions"
	query = query + " " + "WHERE"
	query = query + " " + "status = '3'"
	query = query + " " + "AND context_id = '" + fmt.Sprint(j.ID) + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "date_submitted"
	query = query + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variable
	j.SubmissionsPublished = []int{}
	num := -1

	// View results
	for res.Next() {
		// Get values
		err := res.Scan(&num)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update slice
		j.SubmissionsPublished = append(j.SubmissionsPublished, num)

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
