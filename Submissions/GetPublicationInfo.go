// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "github.com/Mathewwss/ojsinfo/Regex"
import "database/sql"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *Submission) GetPublicationInfo () error {
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
			volume, number, year, published
		FROM
			issues
		WHERE
			issue_id = (
				SELECT DISTINCT
					t2.setting_value
				FROM
					publications AS t1
				INNER JOIN
					publication_settings AS t2
				ON
					t1.publication_id = t2.publication_id
				WHERE
					t2.setting_name = 'issueId'
					AND t1.submission_id = %v
			)
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
	pub_sts := 0
	volume := sql.NullInt64{}
	year := sql.NullInt64{}
	number := sql.NullString{}

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(&volume, &number, &year, &pub_sts)

		// Check errors
		if err != nil {
			// Stop
			return err

		}
	}

	// Check status
	if pub_sts == 1 {
		// Update value
		s.Published = true

	} else {
		// Update value
		s.Published = false

	}

	// Update struct
	s.PublicationYear = year.Int64
	s.PublicationVolume = volume.Int64
	s.PublicationNumber = number.String

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
