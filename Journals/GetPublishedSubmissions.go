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

func (j *Journal) GetPublishedSubmissions () error {
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
