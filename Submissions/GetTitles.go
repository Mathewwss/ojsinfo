// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
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
	query := fmt.Sprintf(`
		SELECT
			t2.locale, t2.setting_name, t2.setting_value
		FROM
			publications AS t1
		INNER JOIN
			publication_settings AS t2
		ON
			t1.publication_id = t2.publication_id
		WHERE
			(
				t2.setting_name = 'title'
				OR t2.setting_name = 'subtitle'
			)
			AND t1.submission_id = %v
		ORDER BY
			t2.locale ASC,
			t2.setting_name DESC
		;
	`, s.ID)

	// Same line
	Regex.OneLine(&query)

	// Run query
	res, err := DbCfg.Db_conf.Con.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variables
	s.Titles = map[string][]string{}
	value := ""
	locale := ""
	name := ""

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&locale, &name, &value)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		if name == "title" {

			s.Titles[locale] = []string{"", ""}

			s.Titles[locale][0] = value

		} else {

			s.Titles[locale][1] = value

		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
