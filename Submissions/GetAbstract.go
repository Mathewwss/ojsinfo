// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
import "database/sql"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// Get date start
func (s *Submission) GetAbstract () error {
	// Sql query
	query := "SELECT"
	query = query + " " + "t1.locale, t1.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "publication_settings AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "publications AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.publication_id = t2.publication_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.setting_value <> ''"
	query = query + " " + "AND t1.setting_name = 'abstract'"
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
	abs := ""
	locale := ""
	s.Abstract = map[string]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &abs)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// One line
		abs := Regex.Html2Text(abs)

		// Update map
		s.Abstract[locale] = abs

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
