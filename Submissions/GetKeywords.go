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
func (s *Submission) GetKeywords () error {
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
			controlled_vocab_entry_settings AS t1
		INNER JOIN
			controlled_vocab_entries AS t2
		ON
			t1.controlled_vocab_entry_id = t2.controlled_vocab_entry_id
		INNER JOIN
			controlled_vocabs AS t3
		ON
			t2.controlled_vocab_id = t3.controlled_vocab_id
		INNER JOIN
			publications AS t4
		ON
			t3.assoc_id = t4.publication_id
		WHERE
			t1.setting_name = 'submissionKeyword'
			AND t4.submission_id = %v
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
