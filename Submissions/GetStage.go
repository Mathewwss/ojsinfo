// ----------------------------- Package ---------------------------- //

package Submissions

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

// Get date start
func (s *Submission) GetStage () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Stages
	possibles := map[string]string{
		"1:1:0":"Submission",
		"1:1:1":"Incompleted",
		"1:1:2":"Incompleted",
		"1:1:3":"Incompleted",
		"1:1:4":"Incompleted",
		"1:3:0":"Review",
		"1:4:0":"Copyediting",
		"1:5:0":"Production",
		"3:1:0":"Published",
		"3:3:0":"Published",
		"3:4:0":"Published",
		"3:5:0":"Published",
		"4:1:0":"Rejected",
		"4:1:2":"Rejected",
		"4:1:3":"Rejected",
		"4:1:5":"Rejected",
		"4:3:0":"Rejected",
		"4:4:0":"Rejected",
		"4:5:0":"Rejected",
		"5:5:0":"Scheduled",
	}

	// Sql query
	query := fmt.Sprintf(`
		SELECT
			CONCAT(status, ":", stage_id, ":", submission_progress)
		FROM
			submissions
		WHERE
			submission_id = %v
		;
	`, s.ID)

	// Same line
	Regex.OneLine(&query)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variables
	status := ""

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&status)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update map
		s.Stage = possibles[status]

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
