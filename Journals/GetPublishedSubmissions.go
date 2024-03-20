// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
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
	query := fmt.Sprintf(`
		SELECT DISTINCT
			submission_id
		FROM
			submissions
		WHERE
			status = 3
			AND context_id = %v
		ORDER BY
			date_submitted
		;
	`, j.ID)

	// Same line
	Regex.OneLine(&query)

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
