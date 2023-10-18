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

// Get date start
func (s *Submission) GetKeywords () error {
	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "t1.locale, t1.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "controlled_vocab_entry_settings AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "controlled_vocab_entries AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.controlled_vocab_entry_id ="
	query = query + " " + "t2.controlled_vocab_entry_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "controlled_vocabs AS t3"
	query = query + " " + "ON"
	query = query + " " + "t2.controlled_vocab_id ="
	query = query + " " + "t3.controlled_vocab_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "publications AS t4"
	query = query + " " + "ON"
	query = query + " " + "t3.assoc_id = t4.publication_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.setting_name = 'submissionKeyword'"
	query = query + " " + "AND t4.submission_id = '6698'"
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
	s.Keywords = map[string][]string{}
	locale := ""
	keyword := ""

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &keyword)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// View size
		if len(s.Keywords[locale]) == 0 {
			// Start variables
			s.Keywords[locale] = []string{}
			s.Keywords[locale] = append(s.Keywords[locale], keyword)

		} else {
			// Append slice
			s.Keywords[locale] = append(s.Keywords[locale], keyword)

		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
