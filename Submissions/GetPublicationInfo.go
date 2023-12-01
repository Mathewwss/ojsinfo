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

func (s *Submission) GetPublicationInfo () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Sql query
	query := fmt.Sprint("SELECT DISTINCT")
	query = query + " " + "volume, number, year, published"
	query = query + " " + "FROM"
	query = query + " " + "issues"
	query = query + " " + "WHERE"
	query = query + " " + "issue_id = ("
	query = query + " " + "SELECT DISTINCT"
	query = query + " " + "t2.setting_value"
	query = query + " " + "FROM"
	query = query + " " + "publications AS t1"
	query = query + " " + "INNER JOIN"
	query = query + " " + "publication_settings AS t2"
	query = query + " " + "ON"
	query = query + " " + "t1.publication_id = t2.publication_id"
	query = query + " " + "WHERE"
	query = query + " " + "t2.setting_name = 'issueId'"
	query = query + " " + "AND t1.submission_id = '"
	query = query + fmt.Sprint(s.ID) + "'"
	query = query + " " + ")"
	query = query + " " + ";"

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
