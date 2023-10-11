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
func (s *Submission) GetTitles () error {
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

	// Database
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	db, err := sql.Open(driver, con)

	if err != nil {

		return err

	}

	err = db.Ping()

	if err != nil {

		return err

	}

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

		res, err := db.Query(run)

		if err != nil {

			return err

		}

		value := ""
		locale := ""

		for res.Next() {

			err = res.Scan(&locale, &value)

			if err != nil {

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

	return nil

}

// ------------------------------------------------------------------ //