// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "database/sql"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (j *Journal) GetCopyeditingSubmissions () error {
	// Sql query
	query := "SELECT DISTINCT"
	query = query + " " + "submission_id"
	query = query + " " + "FROM"
	query = query + " " + "submissions"
	query = query + " " + "WHERE"
	query = query + " " + "status = '1'"
	query = query + " " + "AND stage_id = '4'"
	query = query + " " + "AND submission_progress = '0'"
	query = query + " " + "AND context_id = '" + fmt.Sprint(j.ID) + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "date_submitted"
	query = query + ";"

	// Database connection settings
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect db
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Run query
	res, err := db.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variable
	j.SubmissionsCopyediting = []int{}
	num := -1

	// View results
	for res.Next() {
		// Get values
		err := res.Scan(&num)

		// CHeck errors
		if err != nil {
			// Stop
			return err

		}

		// Update slice
		j.SubmissionsCopyediting = append(j.SubmissionsCopyediting, num)

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
