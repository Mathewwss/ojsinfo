// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *Submission) GetSection () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

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

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

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
