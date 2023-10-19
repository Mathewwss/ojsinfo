// ----------------------------- Package ---------------------------- //

package Submissions

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

func (s *Submission) GetJournalNames () error {
	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "t2.locale, t2.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "submissions AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "journal_settings AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.context_id = t2.journal_id"
	query = query + " " + "WHERE"
	query = query + " " + "t2.setting_name = 'name'"
	query = query + " " + "AND t1.submission_id = '" + fmt.Sprint(s.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t2.locale"
	query = query + ";"

	// Database conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect database
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

	// Start variables
	locale := ""
	name := ""
	s.JournalNames = map[string]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &name)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update map
		s.JournalNames[locale] = name

	}

	// Finish
	return nil

}

// ------------------------------------------------------------------ //
