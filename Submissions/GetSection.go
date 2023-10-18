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

func (s *Submission) GetSection () error {
	// Sql query
	query := "SELECT"
	query = query + " " + "t1.locale, t1.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "section_settings AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "publications AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.section_id = t2.section_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.setting_name = 'title'"
	query = query + " " + "AND t2.submission_id = '" + fmt.Sprint(s.ID)
	query = query + "'"
	query = query + " " + "ORDER BY"
	query = query + " " + "t1.locale"
	query = query + ";"

	// Database conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect database
	db, err := sql.Open(driver, con)

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
	title := ""
	s.Section = map[string]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &title)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Update map
		s.Section[locale] = title

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
