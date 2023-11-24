// ----------------------------- Package ---------------------------- //

package Submissions

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/ojsinfo/DbCfg"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

// Get date start
func (s *Submission) GetAccess () error {
	// Check connection
	err := DbCfg.Db_conf.CheckCon()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Check publication
	err = s.GetPublicationInfo()

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Check publication status
	if s.Published == false {
		// Stop
		return nil

	}

	// Access type -> [assoc_type, file_type]
	access_type := [][]string{
		[]string{"Abstract", "1048585", "-1"},
		[]string{"HTML", "515", "1"},
		[]string{"PDF", "515", "2"},
		[]string{"Others", "515", "3"},
	}

	// Base query
	base := "SELECT"
	base = base + " " + "CASE"
	base = base + " " + "WHEN SUM(metric) IS NULL THEN 0"
	base = base + " " + "ELSE SUM(metric)"
	base = base + " " + "END"
	base = base + " " + "FROM"
	base = base + " " + "metrics"
	base = base + " " + "WHERE"
	base = base + " " + "submission_id = '" + fmt.Sprint(s.ID)
	base = base + "'"

	// Start map
	s.Access = map[string]int{}
	s.Access["All"] = 0

	// View types
	for a := 0; a < len(access_type); a++ {
		// Default value
		s.Access[access_type[a][0]] = 0

		// Update query
		query := base
		query = query + " " + "AND assoc_type = '"
		query = query + access_type[a][1] + "'"

		// Check file type
		if access_type[a][2] != "-1" {
			// Update query
			query = query + " " + "AND file_type = '"
			query = query + access_type[a][2] + "'"

		}

		// Finish query
		query = query + ";"

		// Run query
		res, err := DbCfg.Db_conf.Con.Query(query)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// View results
		for res.Next() {
			// Get value
			num := 0
			err = res.Scan(&num)

			// Check errors
			if err != nil {
				// Stop
				return err

			}

			// Update map
			s.Access[access_type[a][0]] = num
			s.Access["All"] = s.Access["All"] + num

		}
	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
