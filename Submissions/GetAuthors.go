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

func (s *Submission) GetAuthors () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	query := fmt.Sprintf(`
		SELECT
			s1.seq, s1.locale, s1.email,
			CONCAT(s1.setting_value, " ",  s2.setting_value)
		FROM
			(
				SELECT DISTINCT
					t1.seq, t3.locale, t1.email, t3.setting_value
				FROM
					authors AS t1
				INNER JOIN
					publications AS t2
				ON
					t1.publication_id = t2.publication_id
				INNER JOIN
					author_settings AS t3
				ON
					t1.author_id = t3.author_id
				WHERE
					t3.setting_name = 'givenName'
					AND t2.submission_id = %v
					AND t3.setting_value <> ''
				ORDER BY
					t1.seq ASC, t3.locale ASC
			) AS s1
		INNER JOIN
			(
				SELECT DISTINCT
					t1.seq, t3.locale, t1.email, t3.setting_value
				FROM
					authors AS t1
				INNER JOIN
					publications AS t2
				ON
					t1.publication_id = t2.publication_id
				INNER JOIN
					author_settings AS t3
				ON
					t1.author_id = t3.author_id
				WHERE
					t3.setting_name = 'familyName'
					AND t2.submission_id = %v
					AND t3.setting_value <> ''
				ORDER BY
					t1.seq ASC, t3.locale ASC
			) AS s2
		ON
			s1.email = s2.email
		ORDER BY
			s1.seq DESC,
			s1.locale ASC
		;
	`, s.ID, s.ID)

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
	seq := uint8(0)
	locale := ""
	email := ""
	name := ""
	s.Authors = map[uint8]map[string][]string{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&seq, &locale, &email, &name)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Check size
		if len(s.Authors[seq + 1]) == 0 {
			// Start nested map
			s.Authors[seq + 1] = map[string][]string{}

		}

		// Update map
		s.Authors[seq + 1][locale] = []string{name, email}

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
