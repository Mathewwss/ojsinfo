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

func (s *Submission) GetJournalNames () error {
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
			t2.locale, t2.setting_value
		FROM
			submissions AS t1
		INNER JOIN
			journal_settings AS t2
		ON
			t1.context_id = t2.journal_id
		WHERE
			t2.setting_name = 'name'
			AND t1.submission_id = %v
		ORDER BY
			t2.locale
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
