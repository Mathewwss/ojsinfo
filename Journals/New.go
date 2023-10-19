// ----------------------------- Package ---------------------------- //

package Journals

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

func New (identity string) (Journal, error) {
	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "journal_id, path"
	query = query + " " + "FROM"
	query = query + " " + "journals"
	query = query + " " + "WHERE"
	query = query + " " + "path = '" + identity + "'"
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
		return Journal{}, err

	}

	// Check connection
	err = db.Ping()

	// Check errors
	if err != nil {
		// Stop
		return Journal{}, err

	}

	// Run query
	res, err := db.Query(query)

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
