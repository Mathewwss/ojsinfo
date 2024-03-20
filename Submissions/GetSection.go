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

func (s *Submission) GetSection () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := fmt.Sprintf(`
		SELECT DISTINCT
			t1.locale, t1.setting_value
		FROM
			section_settings AS t1
		INNER JOIN
			publications AS t2
		ON
			t1.section_id = t2.section_id
		WHERE
			t1.setting_name = 'title'
			AND t2.submission_id = %v
		ORDER BY
			t1.locale
		;
	`, s.ID)

	// Same Line
	Regex.OneLine(&query)

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
