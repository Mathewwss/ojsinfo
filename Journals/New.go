// ----------------------------- Package ---------------------------- //

package Journals

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
import "errors"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func New (identity string) (Journal, error) {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return Journal{}, err

	}

	// Sql query
	query := fmt.Sprintf(`
		SELECT DISTINCT
			journal_id, path
		FROM
			journals
		WHERE
			path = '%v'
		;
	`, identity)

	// One line
	Regex.OneLine(&query)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return Journal{}, err

	}

	// Start variables
	j := Journal{}
	j.ID = -1

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&j.ID, &j.Path)

		// Check errors
		if err != nil {
			// Stop
			return Journal{}, err

		}

	}

	// Check journal id
	if j.ID == -1 {
		// Create error
		msg := "[ERROR] -> Not found journal by path '" + identity
		msg = msg + "'!"
		err = errors.New(msg)

		// Stop
		return Journal{}, err

	}

	// Show journal
	return j, nil

}

// ------------------------------------------------------------------ //
