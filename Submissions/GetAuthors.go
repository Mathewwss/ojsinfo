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

func (s *Submission) GetAuthors () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Base query
	query := "SELECT DISTINCT"
	query = query + " " + "t1.seq, t3.locale, t1.email,"
	query = query + " " + "t3.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "authors AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "publications AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.publication_id = t2.publication_id"
	query = query + " " + "INNER JOIN"
	query = query + " " + "author_settings AS t3"
	query = query + " " + "ON"
	query = query + " " + "t1.author_id = t3.author_id"
	query = query + " " + "WHERE"
	query = query + " " + "t2.submission_id = '" + fmt.Sprint(s.ID)
	query = query + " " + "'"
	query = query + " " + "AND ("
	query = query + " " + "t3.setting_name = 'givenName'"
	query = query + " " + "OR t3.setting_name = 'familyName'"
	query = query + " " + ")"
	query = query + " " + "ORDER BY"
	query = query + " " + "t1.seq ASC,"
	query = query + " " + "t3.locale ASC,"
	query = query + " " + "t3.setting_name DESC"
	query = query + " " + ";"

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variables
	seq := -1
	last_seq := -1
	locale := ""
	last_locale := ""
	email := ""
	value := ""
	name := ""
	s.Authors = map[int]map[string][]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&seq, &locale, &email, &value)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Check last values
		if last_seq == seq && last_locale == locale {
			// Update name
			name = name + " " + value

			if len(s.Authors[seq]) == 0 {
				// Start map
				s.Authors[seq] = map[string][]string{}

			}

			// Email error
			if email == "<![CDATA[]]>" {
				// Slice update (empty email)
				s.Authors[seq][locale] = []string{name, ""}

			} else {
				// Slice update (empty email)
				s.Authors[seq][locale] = []string{name, email}

			}

		} else {
			// Update last variables
			last_seq = seq
			last_locale = locale

			// Update name
			name = value

		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
