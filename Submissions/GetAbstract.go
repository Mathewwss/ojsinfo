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

// Get date start
func (s *Submission) GetAbstract () error {
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
			publication_settings AS t1
		INNER JOIN
			publications AS t2
		ON
			t1.publication_id = t2.publication_id
		WHERE
			t1.setting_value <> ''
			AND t1.setting_name = 'abstract'
			AND t2.submission_id = %v
		ORDER BY
			t1.locale
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
