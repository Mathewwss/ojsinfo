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

	// Database conf
	driver := DbCfg.Db_conf.Driver
	con := DbCfg.Db_conf.Settings

	// Connect database
	db, err := sql.Open(driver, con)

	// Finish Connection
	defer db.Close()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Run query
	res, err := db.Query(query)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Start variables
	pub_sts := 0
	volume := &s.PublicationVolume
	year := &s.PublicationYear
	number := &s.PublicationNumber

	// View results
	for res.Next() {
		// Get values
		err = res.Scan(volume, number, year, &pub_sts)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

	}

	// Check values
	if *volume == 0 && *year == 0 && *number == "" {
		// Update struct
		*volume = -1
		*year = -1
		*number = ""
		s.Published = false

	} else {
		// Update struct
		s.Published = true

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
