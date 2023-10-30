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

func (s *Submission) GetTitles () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Base query
	query := "SELECT"
	query = query + " " + "t2.locale, t2.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "publications AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "publication_settings AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.publication_id = t2.publication_id"
	query = query + " " + "WHERE"
	query = query + " " + "t1.submission_id = '" + fmt.Sprint(s.ID)
	query = query + "'"

	// Titles type
	titles := []string{
		"title",
		"subtitle",
	}

	for a := 0; a < len(titles); a++ {
		// View loop
		if a == 0 {
			// Start map
			s.Titles = map[string]string{}

		}

		// Finaly query
		run := query + " " + "AND t2.setting_name = '" + titles[a]
		run = run + "'"
		run = run + " " + "ORDER BY"
		run = run + " " + "t2.locale"
		run = run + ";"

		// Run query
		res, err := DbCfg.Db_conf.Con.Query(query)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Start variables
		value := ""
		locale := ""

		// View results
		for res.Next() {
			// Get values
			err = res.Scan(&locale, &value)

			// Check errors
			if err != nil {
				// Stop
				return err

			}

			// Empty value
			if value == "" {
				// Next loop
				continue

			}

			// Check size
			if len(s.Titles[locale]) == 0 {
				// Title
				s.Titles[locale] = value

			} else {
				// Second title
				s.Titles[locale] = s.Titles[locale] + " " + value

			}
		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
