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

func (j *Journal) GetNewSubmissions () error {
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
			status = 1
			AND stage_id = 1
			AND submission_progress = 0
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
	j.SubmissionsNew = []int{}
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
		j.SubmissionsNew = append(j.SubmissionsNew, num)

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
